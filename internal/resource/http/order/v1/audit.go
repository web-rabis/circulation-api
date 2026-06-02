package v1

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/web-rabis/circulation-api/internal/resource/http/order/v1/dto"
	"github.com/web-rabis/httperrors"
)

func (res *OrderResource) auditOrder(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		_ = render.Render(w, r, httperrors.BadRequest(errors.New("id is required")))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	records, err := res.orderMan.Audit(r.Context(), id)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	render.JSON(w, r, dto.AuditResponse{Result: records})
}
