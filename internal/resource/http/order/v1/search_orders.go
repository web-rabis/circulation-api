package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/web-rabis/circulation-api/internal/resource/http/order/v1/dto"
	"github.com/web-rabis/httperrors"
	orderModel "github.com/web-rabis/order-client/model"
)

func (res *OrderResource) searchOrders(w http.ResponseWriter, r *http.Request) {
	var request dto.SearchOrdersRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	if request.Query == "" {
		_ = render.Render(w, r, httperrors.BadRequest(errors.New("query is required")))
		return
	}

	paging, err := orderModel.PagingParseFromHttp(r)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	count, orders, err := res.orderMan.Search(r.Context(), request.Query, paging)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	render.JSON(w, r, dto.SearchOrdersResponse{
		Result: orders,
		Count:  count,
	})
}
