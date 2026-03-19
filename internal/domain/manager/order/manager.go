package order

import (
	"context"

	orderClient "github.com/web-rabis/order-client/client"
	orderModel "github.com/web-rabis/order-client/model"
)

type IManager interface {
	List(ctx context.Context, filters *orderModel.OrderFilters, paging *orderModel.Paging) (int64, []*orderModel.Order, error)
	ById(ctx context.Context, id int64) (*orderModel.Order, error)
}
type Manager struct {
	orderCl orderClient.OrderService
}

func NewOrderManager(orderCl orderClient.OrderService) *Manager {
	return &Manager{
		orderCl: orderCl,
	}
}
func (m *Manager) List(ctx context.Context, filters *orderModel.OrderFilters, paging *orderModel.Paging) (int64, []*orderModel.Order, error) {
	return m.orderCl.List(ctx, paging, filters)
}
func (m *Manager) ById(ctx context.Context, id int64) (*orderModel.Order, error) {
	return m.ById(ctx, id)
}
