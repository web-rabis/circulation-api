package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/web-rabis/circulation-api/internal/resource/http/order/v1/dto"
	orderModel "github.com/web-rabis/order-client/model"
)

func (res *OrderResource) sseReaderMonitor(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	// Снимаем WriteTimeout для SSE-соединения
	rc := http.NewResponseController(w)
	_ = rc.SetWriteDeadline(time.Time{})

	// SSE-заголовки
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	w.WriteHeader(http.StatusOK)

	ctx := r.Context()

	fetchMonitor := func() ([]dto.ReaderMonitorItem, error) {
		departmentId := getDepartmentId(r)
		filters := &orderModel.OrderFilters{
			States:       dto.MonitorStates(),
			DepartmentId: departmentId,
			Period:       orderModel.OrderPeriodToday,
		}
		_, orders, err := res.orderMan.List(ctx, filters, &orderModel.Paging{Limit: 0})
		if err != nil {
			return nil, err
		}
		return dto.NewReaderMonitor(orders), nil
	}

	sendMonitor := func(items []dto.ReaderMonitorItem) {
		data, err := json.Marshal(items)
		if err != nil {
			return
		}
		fmt.Fprintf(w, "data: %s\n\n", data)
		flusher.Flush()
	}

	readerMonitorEqual := func(a, b []dto.ReaderMonitorItem) bool {
		if len(a) != len(b) {
			return false
		}
		index := make(map[string][2]int, len(a))
		for _, item := range a {
			index[item.Reader.TicketNumber] = [2]int{item.Total, item.Completed}
		}
		for _, item := range b {
			prev, ok := index[item.Reader.TicketNumber]
			if !ok || prev[0] != item.Total || prev[1] != item.Completed {
				return false
			}
		}
		return true
	}

	// Первый снимок при подключении
	lastItems, err := fetchMonitor()
	if err != nil {
		fmt.Fprintf(w, "event: error\ndata: %s\n\n", err.Error())
		flusher.Flush()
	} else {
		sendMonitor(lastItems)
	}

	pollTicker := time.NewTicker(ssePollingInterval)
	pingTicker := time.NewTicker(ssePingInterval)
	defer pollTicker.Stop()
	defer pingTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-pingTicker.C:
			fmt.Fprintf(w, ": ping\n\n")
			flusher.Flush()

		case <-pollTicker.C:
			current, err := fetchMonitor()
			if err != nil {
				fmt.Fprintf(w, "event: error\ndata: %s\n\n", err.Error())
				flusher.Flush()
				continue
			}
			if !readerMonitorEqual(lastItems, current) {
				sendMonitor(current)
				lastItems = current
			}
		}
	}
}
