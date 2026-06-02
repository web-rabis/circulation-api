package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"

	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/domain/manager/search"
)

type SearchResource struct {
	path      string
	authMan   *auth.Manager
	searchMan search.IManager
}

func NewSearchResource(path string, authMan *auth.Manager, searchMan search.IManager) *SearchResource {
	return &SearchResource{
		path:      path,
		authMan:   authMan,
		searchMan: searchMan,
	}
}
func (res *SearchResource) Path() string {
	return res.path
}

func (res *SearchResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(res.authMan.JWTAuth()))
		r.Use(auth.NewUserAccessCtx(res.authMan.JWTKey()).ChiMiddleware)
		r.Post("/ebook/simple", res.ebookSimpleSearch)
		r.Post("/periodical/simple", res.periodicalSimpleSearch)
	})

	return r
}
