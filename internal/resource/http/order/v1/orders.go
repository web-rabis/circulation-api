package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/web-rabis/circulation-api/internal/domain/model/order"
	"github.com/web-rabis/circulation-api/internal/resource/http/order/v1/dto"
	"github.com/web-rabis/httperrors"
	orderModel "github.com/web-rabis/order-client/model"
)

func (res *OrderResource) orders(w http.ResponseWriter, r *http.Request) {

	paging, err := orderModel.PagingParseFromHttp(r)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	filters, err := parseFilters(r)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	count, orders, err := res.orderMan.List(r.Context(), filters, paging)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	response := dto.OrdersResponse{
		Result: orders,
		Count:  count,
	}

	render.JSON(w, r, response)
}
func parseFilters(r *http.Request) (*orderModel.OrderFilters, error) {
	var ticketNumber_ int64
	ticketNumber := r.URL.Query().Get("ticketNumber")
	if ticketNumber != "" {
		ticketNumber_, _ = strconv.ParseInt(ticketNumber, 10, 64)
	}

	return &orderModel.OrderFilters{
		TicketNumber: ticketNumber_,
		States: []string{
			order.OrderStateInHands,
			order.OrderStateOrdered,
			order.OrderStateInStorage,
			order.OrderStateInReadingHall,
			order.OrderStatePostponed,
			order.OrderStateRejected,
		},
	}, nil
}
