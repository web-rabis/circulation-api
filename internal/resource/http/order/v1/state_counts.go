package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/render"

	"github.com/web-rabis/circulation-api/internal/resource/http/order/v1/dto"
	"github.com/web-rabis/httperrors"
	orderModel "github.com/web-rabis/order-client/model"
)

func (res *OrderResource) stateCounts(w http.ResponseWriter, r *http.Request) {

	filters, err := parseStateCountFilters(r)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	stateCounts, err := res.orderMan.StateCounts(r.Context(), filters)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	render.JSON(w, r, dto.NewStateCounts(stateCounts))
}
func parseStateCountFilters(r *http.Request) (*orderModel.StateCountFilters, error) {
	var states []string
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
	return &orderModel.StateCountFilters{
		States:       states,
		DepartmentId: departmentId,
		Period:       period,
	}, nil
}
