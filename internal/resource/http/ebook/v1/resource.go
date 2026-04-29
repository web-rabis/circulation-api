package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"

	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/domain/manager/ebook"
)

type EbookResource struct {
	path     string
	authMan  *auth.Manager
	ebookMan ebook.IManager
}

func NewEbookResource(path string, authMan *auth.Manager, ebookMan ebook.IManager) *EbookResource {
	return &EbookResource{
		path:     path,
		authMan:  authMan,
		ebookMan: ebookMan,
	}
}

func (res *EbookResource) Path() string {
	return res.path
}

func (res *EbookResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(res.authMan.JWTAuth()))
		r.Use(auth.NewUserAccessCtx(res.authMan.JWTKey()).ChiMiddleware)
		r.Get("/{id}/card", res.ebookCardById)
	})

	return r
}
