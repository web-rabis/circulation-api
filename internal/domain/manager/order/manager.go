package order

import (
	"context"

	orderClient "github.com/web-rabis/order-client/client"
	orderModel "github.com/web-rabis/order-client/model"
	readerClient "github.com/web-rabis/reader-client/client"
)

type IManager interface {
	List(ctx context.Context, filters *orderModel.OrderFilters, paging *orderModel.Paging) (int64, []*orderModel.Order, error)
	ById(ctx context.Context, id int64) (*orderModel.Order, error)
	StateCounts(ctx context.Context, filters *orderModel.StateCountFilters) ([]*orderModel.StateCount, error)
	Reject(ctx context.Context, ids []int64, rejectId, userId int64) error
	Redirect(ctx context.Context, ids []int64, departmentId, userId int64) error
}
type Manager struct {
	orderCl  orderClient.OrderService
	readerCl readerClient.ReaderService
}

func NewOrderManager(orderCl orderClient.OrderService, readerCl readerClient.ReaderService) *Manager {
	return &Manager{
		orderCl:  orderCl,
		readerCl: readerCl,
	}
}
func (m *Manager) List(ctx context.Context, filters *orderModel.OrderFilters, paging *orderModel.Paging) (int64, []*orderModel.Order, error) {
	count, orders, err := m.orderCl.List(ctx, paging, filters)
	if err != nil {
		return 0, nil, err
	}
	for _, order := range orders {
		reader, err := m.readerCl.ReaderById(ctx, order.Reader.TicketNumber)
		if err != nil {
			continue
		}
		order.Reader.Firstname = reader.Firstname
		order.Reader.Lastname = reader.Lastname
		order.Reader.Middlename = reader.Middlename
		order.Reader.Barcode = reader.Barcode
		order.Reader.IsEmployee = reader.IsEmployee
		order.Reader.Department = reader.Department
	}
	return count, orders, nil
}
func (m *Manager) ById(ctx context.Context, id int64) (*orderModel.Order, error) {
	return m.ById(ctx, id)
}
func (m *Manager) StateCounts(ctx context.Context, filters *orderModel.StateCountFilters) ([]*orderModel.StateCount, error) {
	return m.orderCl.StateCounts(ctx, filters)
}
func (m *Manager) Reject(ctx context.Context, ids []int64, rejectId, userId int64) error {
	return m.orderCl.Reject(ctx, ids, rejectId, userId)
}
func (m *Manager) Redirect(ctx context.Context, ids []int64, departmentId, userId int64) error {
	return m.orderCl.Redirect(ctx, ids, departmentId, userId)
}
