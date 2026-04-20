package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"

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
	var states []string
	ticketNumber := r.URL.Query().Get("ticketNumber")
	if ticketNumber != "" {
		ticketNumber_, _ = strconv.ParseInt(ticketNumber, 10, 64)
	}
	status := r.URL.Query().Get("status")
	switch status {
	case orderModel.OrderStateInHands, orderModel.OrderStateOrdered, orderModel.OrderStateInStorage, orderModel.OrderStateInReadingHall, orderModel.OrderStatePostponed, orderModel.OrderStateRejected, orderModel.OrderStateInAuxiliaryFund, orderModel.OrderStateReaderReturned, orderModel.OrderStateReturnToStorage:
		states = append(states, status)
	default:
		states = []string{
			orderModel.OrderStateInHands,
			orderModel.OrderStateOrdered,
			orderModel.OrderStateInStorage,
			orderModel.OrderStateInReadingHall,
			orderModel.OrderStatePostponed,
			orderModel.OrderStateRejected,
			orderModel.OrderStateInAuxiliaryFund,
			orderModel.OrderStateReaderReturned,
			orderModel.OrderStateReturnToStorage,
		}
	}
	departmentId, _ := strconv.ParseInt(r.URL.Query().Get("departmentId"), 10, 64)
	period := r.URL.Query().Get("period")
	switch period {
	case orderModel.OrderPeriodToday:
	case orderModel.OrderPeriodWeek:
	case orderModel.OrderPeriodMonth:
	case orderModel.OrderPeriodQuarter:
	case orderModel.OrderPeriodYear:
	default:
		period = ""
	}
	var isAuxiliaryFund *bool
	if r.URL.Query().Get("isAuxiliaryFund") != "" {
		af, err := strconv.ParseBool(r.URL.Query().Get("isAuxiliaryFund"))
		if err == nil {
			isAuxiliaryFund = &af
		}
	}
	return &orderModel.OrderFilters{
		TicketNumber:    ticketNumber_,
		States:          states,
		DepartmentId:    departmentId,
		Period:          period,
		IsAuxiliaryFund: isAuxiliaryFund,
	}, nil
}
