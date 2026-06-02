package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"

	"github.com/web-rabis/circulation-api/internal/resource/http/search/v1/dto"
	"github.com/web-rabis/httperrors"
)

func (res *SearchResource) periodicalSimpleSearch(w http.ResponseWriter, r *http.Request) {

	var request dto.PeriodicalSimpleSearchRequest
	var err error

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	paging, err := pagingParse(r)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	count, result, err := res.searchMan.PeriodicalSimpleSearch(r.Context(), request.Query, paging)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	render.JSON(w, r, dto.PeriodicalSimpleSearchResponse{
		Count:  count,
		Result: result,
	})
}
