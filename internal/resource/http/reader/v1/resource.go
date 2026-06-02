package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"

	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/domain/manager/reader"
)

type ReaderResource struct {
	path      string
	authMan   *auth.Manager
	readerMan reader.IManager
}

func NewReaderResource(path string, authMan *auth.Manager, readerMan reader.IManager) *ReaderResource {
	return &ReaderResource{
		path:      path,
		authMan:   authMan,
		readerMan: readerMan,
	}
}
func (res *ReaderResource) Path() string {
	return res.path
}

func (res *ReaderResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(res.authMan.JWTAuth()))
		r.Use(auth.NewUserAccessCtx(res.authMan.JWTKey()).ChiMiddleware)
		r.Get("/{ticketNumber}", res.readerByTicketNumber)
	})

	return r
}
