package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/resource/http/order/v1/dto"
	"github.com/web-rabis/httperrors"
)

func (res *OrderResource) createOrder(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if request.TicketNumber == 0 {
		_ = render.Render(w, r, httperrors.BadRequest(errors.New("ticketNumber is required")))
		return
	}
	if request.EbookId == 0 {
		_ = render.Render(w, r, httperrors.BadRequest(errors.New("ebookId is required")))
		return
	}
	if request.DepartmentId == 0 {
		_ = render.Render(w, r, httperrors.BadRequest(errors.New("departmentId is required")))
		return
	}

	userId, err := auth.UserIdFromContext(r.Context())
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	_, err = res.orderMan.Create(r.Context(), request.TicketNumber, request.EbookId, request.DepartmentId, request.InventoryId, userId)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	render.HTML(w, r, "OK")
}
