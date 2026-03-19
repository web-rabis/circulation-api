package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/domain/manager/user"
)

type Resource struct {
	path    string
	authMan *auth.Manager
	userMan *user.Manager
}

func NewAuthResource(basePath string, authMan *auth.Manager, userMan *user.Manager) *Resource {
	return &Resource{
		path:    basePath,
		authMan: authMan,
		userMan: userMan,
	}
}

func (res *Resource) Path() string {
	return res.path
}

func (res *Resource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/sign-in", res.signIn)
	})

	return r
}
