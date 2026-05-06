package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"
	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"

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
	userId, err := auth.UserIdFromContext(r.Context())
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	user, err := res.userSvc.UserById(r.Context(), userId)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
	}
	filters.DepartmentId = user.Department.Id
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
	query := r.URL.Query().Get("query")
	ticketNumber, err := strconv.ParseInt(query, 10, 64)
	if err == nil {
		query = ""
	}
	return &orderModel.OrderFilters{
		TicketNumber:    ticketNumber,
		States:          getStatuses(r),
		Period:          getPeriod(r),
		IsAuxiliaryFund: getIsAuxiliaryFund(r),
		Query:           query,
	}, nil
}
func getStatuses(r *http.Request) []string {
	var states []string
	status := r.URL.Query().Get("status")
	switch status {
	case orderModel.OrderStateInHands, orderModel.OrderStateOrdered, orderModel.OrderStateInStorage, orderModel.OrderStateInReadingHall, orderModel.OrderStatePostponed, orderModel.OrderStateRejected, orderModel.OrderStateInAuxiliaryFund, orderModel.OrderStateReaderReturned, orderModel.OrderStateReturnToStorage, orderModel.OrderStateProcessed:
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
			orderModel.OrderStateProcessed,
		}
	}
	return states
}
func getPeriod(r *http.Request) string {
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
	return period
}
func getDepartmentId(r *http.Request) int64 {
	id, err := strconv.ParseInt(r.URL.Query().Get("departmentId"), 10, 64)
	if err != nil {
		return 0
	}
	return id
}
func getIsAuxiliaryFund(r *http.Request) *bool {
	var isAuxiliaryFund *bool
	if r.URL.Query().Get("isAuxiliaryFund") != "" {
		af, err := strconv.ParseBool(r.URL.Query().Get("isAuxiliaryFund"))
		if err == nil {
			isAuxiliaryFund = &af
		}
	}
	return isAuxiliaryFund
}
