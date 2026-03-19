package drivers

import (
	"context"

	"github.com/web-rabis/db/eorder"
	"github.com/web-rabis/db/model"
)

type DataStore interface {
	Base
}

type Base interface {
	Name() string
	Ping(ctx context.Context) error
	Close(ctx context.Context) error
	Connect(ctx context.Context) error

	EOrder() EOrder
}
type EOrder interface {
	List(ctx context.Context, filter *eorder.EOrderFilter, paging *model.Paging) ([]*eorder.EOrder, error)
	Count(ctx context.Context, filter *eorder.EOrderFilter) (int64, error)
}
