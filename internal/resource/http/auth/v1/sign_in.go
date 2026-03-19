package v1

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/web-rabis/circulation-api/internal/resource/http/auth/v1/dto"
	"github.com/web-rabis/httperrors"
)

// @Summary Dictionary Sex
// @Description Dictionary Sex
// @Accept json
// @Produce json
// @Tags Dictionary
// @Success 200 {object} reader.Dictionary
// @Failure 400 {object} httperrors.Response
// @Failure 401 {object} httperrors.Response
// @Failure 403 {object} httperrors.Response
// @Failure 404 {object} httperrors.Response
// @Failure 500 {object} httperrors.Response
// @Router /admin/api/v1/auth/sign-in [get]
func (res *Resource) signIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var request dto.SignInRequest
	var err error

	if err = json.NewDecoder(r.Body).Decode(&request); err != nil {
		_ = render.Render(w, r, httperrors.BadRequest(err))
		return
	}

	user, err := res.userMan.SignIn(ctx, request.Username, request.Password)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	token, err := res.authMan.NewAccessToken(user.Id)
	if err != nil {
		_ = render.Render(w, r, httperrors.Internal(err))
		return
	}
	render.JSON(w, r, dto.NewSignInResponse(token, user))
}
