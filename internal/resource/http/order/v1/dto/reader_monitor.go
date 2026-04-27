package dto

import (
	"strconv"

	"github.com/web-rabis/circulation-api/internal/domain/model"
	orderModel "github.com/web-rabis/order-client/model"
)

// ReaderDTO — читатель в мониторинге.
type ReaderDTO struct {
	TicketNumber string `json:"ticketNumber"`
	Lastname     string `json:"lastname"`
	Firstname    string `json:"firstname"`
	Middlename   string `json:"middlename"`
}

// ReaderMonitorItem — строка мониторинга: читатель + счётчики заказов.
type ReaderMonitorItem struct {
	Reader    ReaderDTO `json:"reader"`
	Total     int       `json:"total"`
	Completed int       `json:"completed"`
}

// monitorStates — статусы которые нас интересуют
// (без отложенных, обработанных, отказанных, возвращённых)
var monitorStates = []string{
	orderModel.OrderStateOrdered,
	orderModel.OrderStateInStorage,
	orderModel.OrderStateInReadingHall,
	orderModel.OrderStateInAuxiliaryFund,
}

// completedStates — статусы "готово к выдаче"
var completedStates = map[string]bool{
	orderModel.OrderStateInReadingHall:   true,
	orderModel.OrderStateInAuxiliaryFund: true,
}

// MonitorStates возвращает список статусов для фильтрации.
func MonitorStates() []string {
	return monitorStates
}

// NewReaderMonitor группирует заказы по читателям и строит список мониторинга.
func NewReaderMonitor(orders []*model.Order) []ReaderMonitorItem {
	type readerAgg struct {
		reader    *orderModel.Reader
		total     int
		completed int
	}

	index := make(map[int64]*readerAgg)
	keys := make([]int64, 0)

	for _, o := range orders {
		if o.Reader == nil {
			continue
		}
		tn := o.Reader.TicketNumber
		agg, exists := index[tn]
		if !exists {
			agg = &readerAgg{reader: o.Reader}
			index[tn] = agg
			keys = append(keys, tn)
		}
		agg.total++
		if o.State != nil && completedStates[o.State.Code] {
			agg.completed++
		}
	}

	result := make([]ReaderMonitorItem, 0, len(keys))
	for _, tn := range keys {
		agg := index[tn]
		r := agg.reader
		result = append(result, ReaderMonitorItem{
			Reader: ReaderDTO{
				TicketNumber: strconv.FormatInt(r.TicketNumber, 10),
				Lastname:     r.Lastname,
				Firstname:    r.Firstname,
				Middlename:   r.Middlename,
			},
			Total:     agg.total,
			Completed: agg.completed,
		})
	}
	return result
}
