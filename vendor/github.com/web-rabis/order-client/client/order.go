package client

import (
	"context"

	"github.com/web-rabis/order-client/model"
)

type Base interface {
	Connect() error
	Close() error
	Order() OrderService
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name OrderService
type OrderService interface {
	List(ctx context.Context, paging *model.Paging, filter *model.OrderFilters) (int64, []*model.Order, error)
	ById(ctx context.Context, id int64) (*model.Order, error)
}
