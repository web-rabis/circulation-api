package v1

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/web-rabis/circulation-api/internal/resource/http/dictionary/v1/dto"
	"github.com/web-rabis/httperrors"
)

func (res *DictionaryResource) department(w http.ResponseWriter, r *http.Request) {
	_, list, err := res.dictMan.DepartmentList(r.Context())
	if err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}
	render.JSON(w, r, dto.NewDepartmentResponse(list))
}
