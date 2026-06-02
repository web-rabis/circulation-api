package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/render"

	"github.com/web-rabis/circulation-api/internal/resource/http/search/v1/dto"
	"github.com/web-rabis/httperrors"
	"github.com/web-rabis/searcher-proxy/model"
)

func (res *SearchResource) ebookSimpleSearch(w http.ResponseWriter, r *http.Request) {

	var request dto.EbookSimpleSearchRequest
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
	count, result, err := res.searchMan.EbookSimpleSearch(r.Context(), request.Query, paging)
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	render.JSON(w, r, dto.EbookSimpleSearchResponse{
		Count:  count,
		Result: result,
	})
}
func pagingParse(r *http.Request) (*model.Paging, error) {
	paging := model.Paging{}

	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		limit, err := strconv.ParseInt(limitParam, 10, 64)
		if err != nil {
			return nil, err
		}
		paging.Limit = limit
	}

	if pageParam := r.URL.Query().Get("offset"); pageParam != "" {
		offset, err := strconv.ParseInt(pageParam, 10, 64)
		if err != nil {
			return nil, err
		}
		paging.Offset = offset
	}

	if orderParam := r.URL.Query().Get("order"); orderParam != "" {
		paging.SortVal = 1
		order, err := strconv.Atoi(orderParam)
		if err == nil {
			paging.SortVal = int32(order)
		}
	}

	if sortKeyFromQuery := r.URL.Query().Get("orderBy"); sortKeyFromQuery != "" {
		paging.SortKey = sortKeyFromQuery
	}

	return &paging, nil
}
