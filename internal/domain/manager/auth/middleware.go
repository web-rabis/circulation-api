package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"
	"github.com/lestrrat-go/jwx/jwt"

	"github.com/web-rabis/httperrors"
)

var ErrInvalidToken = errors.New("token is incorrect or expired")

type key int

const (
	UserIDKey key = iota + 1
)

type UserAccessCtx struct {
	jwtKey []byte
}

func NewUserAccessCtx(jwtKey []byte) *UserAccessCtx {
	return &UserAccessCtx{
		jwtKey: jwtKey,
	}
}

func (ua UserAccessCtx) ChiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			_ = render.Render(w, r, httperrors.Unauthorized(err))
			return
		}

		if err := jwt.Validate(token); err != nil {
			_ = render.Render(w, r, httperrors.Unauthorized(ErrInvalidToken))
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIDKey, claims["id"])

		next.ServeHTTP(w, r.WithContext(ctx))

	})

}
