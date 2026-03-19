package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/domain/manager/order"
)

type OrderResource struct {
	path     string
	authMan  *auth.Manager
	orderMan order.IManager
}

func NewOrderResource(path string, authMan *auth.Manager, orderMan order.IManager) *OrderResource {
	return &OrderResource{
		path:     path,
		authMan:  authMan,
		orderMan: orderMan,
	}
}

func (res *OrderResource) Path() string {
	return res.path
}

func (res *OrderResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(res.authMan.JWTAuth()))
		r.Use(auth.NewUserAccessCtx(res.authMan.JWTKey()).ChiMiddleware)
		r.Get("/", res.orders)
	})

	return r
}
