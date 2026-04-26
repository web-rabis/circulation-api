package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	errors2 "github.com/pkg/errors"

	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/resource/http/order/v1/dto"
	"github.com/web-rabis/httperrors"
)

func (res *OrderResource) issueOrder(w http.ResponseWriter, r *http.Request) {

	var request dto.IssueOrderRequest
	var err error
	var orderId int64
	var id string

	if id = chi.URLParam(r, "id"); id == "" {
		_ = render.Render(w, r, httperrors.BadRequest(errors2.Errorf("id is empty")))
		return
	}
	orderId, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	userId, err := auth.UserIdFromContext(r.Context())
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	err = res.orderMan.Issue(r.Context(), orderId, userId, request.InventoryId)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	render.HTML(w, r, "OK")
}
