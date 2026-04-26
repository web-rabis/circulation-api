package v1

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"

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
	stateCounts, err := res.orderMan.StateCounts(r.Context(), filters)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	render.JSON(w, r, dto.NewStateCounts(stateCounts))
}
func parseStateCountFilters(r *http.Request) (*orderModel.StateCountFilters, error) {
	return &orderModel.StateCountFilters{
		States: getStatuses(r),
		Period: getPeriod(r),
	}, nil
}
