package http

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/web-rabis/circulation-api/internal/config"
	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/domain/manager/order"
	"github.com/web-rabis/circulation-api/internal/domain/manager/user"
	"github.com/web-rabis/circulation-api/internal/resource/http"
	v1 "github.com/web-rabis/circulation-api/internal/resource/http/auth/v1"
	v2 "github.com/web-rabis/circulation-api/internal/resource/http/order/v1"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	cherver "github.com/web-rabis/servers/http"
)

const (
	maxAge        = 300
	compressLevel = 5
)

func Run(serversCtx context.Context,
	opts *config.APIServer,
	authMan *auth.Manager,
	userMan *user.Manager,
	orderMan *order.Manager,
	version string) error {
	resources := []cherver.Resource{
		http.NewVersionResource("/version", version),
		http.NewFilesResource("/files", opts.ServerConfig.FilesDir),
		//swaggerV1.NewSwaggerResource("/swagger", opts.ServerConfig.BasePath, "/files"),
		v1.NewAuthResource("/api/v1/auth", authMan, userMan),
		v2.NewOrderResource("/api/v1/orders", authMan, orderMan),
	}
	httpSrv := cherver.New(
		cherver.WithListenAddress(opts.ServerConfig.ListenAddr),
		cherver.WithCert(opts.ServerConfig.CertFile, opts.ServerConfig.KeyFile),
		cherver.WithResources(resources...),
		cherver.WithMiddlewares(middlewares(opts)...))

	return httpSrv.Run(serversCtx)

}

func middlewaresWithoutLogs(opts *config.APIServer) chi.Middlewares {
	return chi.Middlewares{
		middleware.NoCache,   // no-cache
		middleware.Recoverer, // управляемо обрабатывает паники и выдает stack trace при их возникновении
		middleware.RealIP,    // устанавливает RemoteAddr для каждого запроса с заголовками X-Forwarded-For или X-Real-IP
		middleware.NewCompressor(compressLevel).Handler,

		cors.Handler(cors.Options{
			AllowedOrigins:   allowedOrigins(opts.IsTesting),
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           maxAge, // Maximum value not ignored by any of major browsers
		})}
}

func middlewares(opts *config.APIServer) chi.Middlewares {
	return append(middlewaresWithoutLogs(opts), middleware.Logger)
}
func allowedOrigins(testing bool) []string {
	if testing {
		return []string{"*"}
	}

	return []string{}
}
