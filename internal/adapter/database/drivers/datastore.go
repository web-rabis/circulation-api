package drivers

import (
	"context"

	"github.com/web-rabis/circulation-api/internal/domain/model"
	"github.com/web-rabis/circulation-api/internal/domain/model/order"
)

type DataStore interface {
	Base
}

type Base interface {
	Name() string
	Ping(ctx context.Context) error
	Close(ctx context.Context) error
	Connect(ctx context.Context) error
	Order() OrderRepository
}

type OrderRepository interface {
	List(ctx context.Context, filter *model.OrderFilter, paging *model.Paging) ([]*order.Order, error)
	Count(ctx context.Context, filter *model.OrderFilter) (int64, error)
	ById(ctx context.Context, id int64) (*order.Order, error)
}
