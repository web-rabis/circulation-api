package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/web-rabis/circulation-api/internal/domain/manager/auth"
	"github.com/web-rabis/circulation-api/internal/resource/http/order/v1/dto"
	orderModel "github.com/web-rabis/order-client/model"
)

const (
	ssePollingInterval = 25 * time.Second
	ssePingInterval    = 25 * time.Second // меньше WriteTimeout сервера (30s)
)

func (res *OrderResource) sseStateCounts(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	// Снимаем WriteTimeout для этого соединения — SSE живёт дольше 30s
	rc := http.NewResponseController(w)
	_ = rc.SetWriteDeadline(time.Time{})

	// Аутентификация через ?token=...
	tokenString := r.URL.Query().Get("token")
	if tokenString == "" {
		http.Error(w, "token is required", http.StatusUnauthorized)
		return
	}
	token, err := res.authMan.JWTAuth().Decode(tokenString)
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	if err := jwt.Validate(token); err != nil {
		http.Error(w, auth.ErrInvalidToken.Error(), http.StatusUnauthorized)
		return
	}
	claims, err := token.AsMap(r.Context())
	if err != nil {
		http.Error(w, "invalid token claims", http.StatusUnauthorized)
		return
	}
	rawID, ok := claims["id"]
	if !ok {
		http.Error(w, "user id not found in token", http.StatusUnauthorized)
		return
	}
	userId := int64(rawID.(float64))

	user, err := res.userSvc.UserById(r.Context(), userId)
	if err != nil {
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		return
	}

	// SSE-заголовки — до первой записи в body
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")
	w.WriteHeader(http.StatusOK)

	ctx := r.Context()

	fetchCounts := func() ([]dto.StateCount, error) {
		filters := &orderModel.StateCountFilters{
			States:       getStatuses(r),
			Period:       getPeriod(r),
			DepartmentId: user.Department.Id,
		}
		stateCounts, err := res.orderMan.StateCounts(ctx, filters)
		if err != nil {
			return nil, err
		}
		return dto.NewStateCounts(stateCounts), nil
	}

	sendCounts := func(counts []dto.StateCount) {
		data, err := json.Marshal(counts)
		if err != nil {
			return
		}
		fmt.Fprintf(w, "data: %s\n\n", data)
		flusher.Flush()
	}

	// Первый снимок сразу при подключении
	lastCounts, err := fetchCounts()
	if err != nil {
		fmt.Fprintf(w, "event: error\ndata: %s\n\n", err.Error())
		flusher.Flush()
	} else {
		sendCounts(lastCounts)
	}

	pollTicker := time.NewTicker(ssePollingInterval)
	pingTicker := time.NewTicker(ssePingInterval)
	defer pollTicker.Stop()
	defer pingTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return

		case <-pingTicker.C:
			fmt.Fprintf(w, ": ping\n\n")
			flusher.Flush()

		case <-pollTicker.C:
			current, err := fetchCounts()
			if err != nil {
				fmt.Fprintf(w, "event: error\ndata: %s\n\n", err.Error())
				flusher.Flush()
				continue
			}
			// Отправляем только если данные изменились
			if !stateCountsEqual(lastCounts, current) {
				sendCounts(current)
				lastCounts = current
			}
		}
	}
}

// stateCountsEqual сравнивает два среза счётчиков по содержимому.
func stateCountsEqual(a, b []dto.StateCount) bool {
	if len(a) != len(b) {
		return false
	}
	index := make(map[string]int64, len(a))
	for _, sc := range a {
		index[sc.StateCode] = sc.Total
	}
	for _, sc := range b {
		if index[sc.StateCode] != sc.Total {
			return false
		}
	}
	return true
}

func tokenFromQuery(r *http.Request) string {
	return r.URL.Query().Get("token")
}

func sseVerifier(ja *jwtauth.JWTAuth) func(http.Handler) http.Handler {
	return jwtauth.Verify(ja, tokenFromQuery)
}
