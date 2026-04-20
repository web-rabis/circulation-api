package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	
	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/domain/manager/dictionary"
)

type DictionaryResource struct {
	path    string
	authMan *auth.Manager
	dictMan dictionary.IManager
}

func NewDictionaryResource(path string, authMan *auth.Manager, dictMan dictionary.IManager) *DictionaryResource {
	return &DictionaryResource{
		path:    path,
		authMan: authMan,
		dictMan: dictMan,
	}
}

func (res *DictionaryResource) Path() string {
	return res.path
}

func (res *DictionaryResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(res.authMan.JWTAuth()))
		r.Use(auth.NewUserAccessCtx(res.authMan.JWTKey()).ChiMiddleware)
		r.Get("/reason-rejection", res.reasonRejection)
		r.Get("/department", res.department)
	})

	return r
}
