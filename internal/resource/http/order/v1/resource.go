package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	ssoClient "github.com/web-rabis/sso-client/client"

	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/domain/manager/order"
)

type OrderResource struct {
	path     string
	authMan  *auth.Manager
	orderMan order.IManager
	userSvc  ssoClient.UserService
}

func NewOrderResource(path string, authMan *auth.Manager, orderMan order.IManager, userSvc ssoClient.UserService) *OrderResource {
	return &OrderResource{
		path:     path,
		authMan:  authMan,
		orderMan: orderMan,
		userSvc:  userSvc,
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
		r.Post("/return", res.returnOrder)
		r.Post("/return-to-storage", res.returnToStorageOrder)
		r.Post("/postponed", res.postponedOrder)
		r.Post("/issue", res.issueOrders)
		r.Post("/{id}/issue", res.issueOrder)
		r.Post("/archive", res.archiveOrder)
		r.Post("/send-to-pf", res.sendToPfOrder)
		r.Post("/cancel-reject", res.cancelRejectOrder)
		r.Post("/reject", res.rejectOrder)
		r.Post("/redirect", res.redirectOrder)
		r.Get("/state-counts", res.stateCounts)
	})

	// SSE — токен передаётся через ?token=..., отдельная группа с собственным verifier
	r.Group(func(r chi.Router) {
		r.Use(sseVerifier(res.authMan.JWTAuth()))
		r.Use(auth.NewUserAccessCtx(res.authMan.JWTKey()).ChiMiddleware)
		r.Get("/sse-state-counts", res.sseStateCounts)
	})

	return r
}
