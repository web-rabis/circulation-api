package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"

	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/resource/http/order/v1/dto"
	"github.com/web-rabis/httperrors"
)

func (res *OrderResource) sendToPfOrder(w http.ResponseWriter, r *http.Request) {

	var request dto.SendToPfOrderRequest
	var err error

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	userId, err := auth.UserIdFromContext(r.Context())
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	err = res.orderMan.SendToPf(r.Context(), request.Ids, userId)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	render.HTML(w, r, "OK")
}
