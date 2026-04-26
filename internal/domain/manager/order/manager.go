package order

import (
	"context"

	"github.com/web-rabis/circulation-api/internal/domain/model"
	ebookClient "github.com/web-rabis/ebook-client/client"
	model2 "github.com/web-rabis/ebook-client/model"
	orderClient "github.com/web-rabis/order-client/client"
	orderModel "github.com/web-rabis/order-client/model"
	readerClient "github.com/web-rabis/reader-client/client"
	ssoClient "github.com/web-rabis/sso-client/client"
)

type IManager interface {
	List(ctx context.Context, filters *orderModel.OrderFilters, paging *orderModel.Paging) (int64, []*model.Order, error)
	ById(ctx context.Context, id int64) (*orderModel.Order, error)
	StateCounts(ctx context.Context, filters *orderModel.StateCountFilters) ([]*orderModel.StateCount, error)
	Reject(ctx context.Context, ids []int64, rejectId, userId int64) error
	CancelReject(ctx context.Context, ids []int64, userId int64) error
	SendToPf(ctx context.Context, ids []int64, userId int64) error
	Archive(ctx context.Context, ids []int64, userId int64) error
	Postponed(ctx context.Context, ids []int64, userId int64) error
	ReturnToStorage(ctx context.Context, ids []int64, userId int64) error
	Return(ctx context.Context, ids []int64, userId int64) error
	Issue(ctx context.Context, id int64, userId, invId int64) error
	IssueOrders(ctx context.Context, id []int64, userId int64) error
	Redirect(ctx context.Context, ids []int64, departmentId, userId int64) error
}
type Manager struct {
	orderCl  orderClient.OrderService
	readerCl readerClient.ReaderService
	userCl   ssoClient.UserService
	ebookCl  ebookClient.EbookService
}

func NewOrderManager(orderCl orderClient.OrderService, readerCl readerClient.ReaderService, userCl ssoClient.UserService, ebookCl ebookClient.EbookService) *Manager {
	return &Manager{
		orderCl:  orderCl,
		readerCl: readerCl,
		userCl:   userCl,
		ebookCl:  ebookCl,
	}
}
func (m *Manager) List(ctx context.Context, filters *orderModel.OrderFilters, paging *orderModel.Paging) (int64, []*model.Order, error) {
	count, orders, err := m.orderCl.List(ctx, paging, filters)
	if err != nil {
		return 0, nil, err
	}
	var orders_ = make([]*model.Order, len(orders))
	for i, order := range orders {
		reader, err := m.readerCl.ReaderById(ctx, order.Reader.TicketNumber)
		if err == nil {
			order.Reader.Firstname = reader.Firstname
			order.Reader.Lastname = reader.Lastname
			order.Reader.Middlename = reader.Middlename
			order.Reader.Barcode = reader.Barcode
			order.Reader.IsEmployee = reader.IsEmployee
			order.Reader.Department = reader.Department

		}
		var e *model2.Ebook
		if order.Ebook != nil {
			e, _ = m.ebookCl.EbookById(ctx, order.Ebook.Id)
		}
		order_ := model.NewOrder(order, e)
		orders_[i] = order_
	}
	return count, orders_, nil
}
func (m *Manager) ById(ctx context.Context, id int64) (*orderModel.Order, error) {
	return m.ById(ctx, id)
}
func (m *Manager) StateCounts(ctx context.Context, filters *orderModel.StateCountFilters) ([]*orderModel.StateCount, error) {
	return m.orderCl.StateCounts(ctx, filters)
}
func (m *Manager) Reject(ctx context.Context, ids []int64, rejectId, userId int64) error {
	user, err := m.getUserById(ctx, userId)
	if err != nil {
		return err
	}
	return m.orderCl.Reject(ctx, ids, rejectId, user)
}
func (m *Manager) CancelReject(ctx context.Context, ids []int64, userId int64) error {
	user, err := m.getUserById(ctx, userId)
	if err != nil {
		return err
	}
	return m.orderCl.CancelReject(ctx, ids, user)
}
func (m *Manager) Redirect(ctx context.Context, ids []int64, departmentId, userId int64) error {
	user, err := m.getUserById(ctx, userId)
	if err != nil {
		return err
	}
	return m.orderCl.Redirect(ctx, ids, departmentId, user)
}
func (m *Manager) SendToPf(ctx context.Context, ids []int64, userId int64) error {
	user, err := m.getUserById(ctx, userId)
	if err != nil {
		return err
	}
	return m.orderCl.SendToPf(ctx, ids, user)
}
func (m *Manager) Archive(ctx context.Context, ids []int64, userId int64) error {
	user, err := m.getUserById(ctx, userId)
	if err != nil {
		return err
	}
	return m.orderCl.Archive(ctx, ids, user)
}
func (m *Manager) Postponed(ctx context.Context, ids []int64, userId int64) error {
	user, err := m.getUserById(ctx, userId)
	if err != nil {
		return err
	}
	return m.orderCl.Postponed(ctx, ids, user)
}
func (m *Manager) ReturnToStorage(ctx context.Context, ids []int64, userId int64) error {
	user, err := m.getUserById(ctx, userId)
	if err != nil {
		return err
	}
	return m.orderCl.ReturnToStorage(ctx, ids, user)
}
func (m *Manager) Return(ctx context.Context, ids []int64, userId int64) error {
	user, err := m.getUserById(ctx, userId)
	if err != nil {
		return err
	}
	return m.orderCl.Return(ctx, ids, user)
}
func (m *Manager) Issue(ctx context.Context, id int64, userId, invId int64) error {
	user, err := m.getUserById(ctx, userId)
	if err != nil {
		return err
	}
	req := []orderModel.IssueOrder{{
		Id:         id,
		EbookInvId: invId,
	}}
	return m.orderCl.Issue(ctx, req, user)
}
func (m *Manager) IssueOrders(ctx context.Context, ids []int64, userId int64) error {
	user, err := m.getUserById(ctx, userId)
	if err != nil {
		return err
	}
	var req = make([]orderModel.IssueOrder, len(ids))
	for i, id := range ids {
		req[i] = orderModel.IssueOrder{Id: id}
	}
	return m.orderCl.Issue(ctx, req, user)
}
func (m *Manager) getUserById(ctx context.Context, id int64) (*orderModel.User, error) {
	user, err := m.userCl.UserById(ctx, id)
	if err != nil {
		return nil, err
	}
	u := &orderModel.User{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		State:    user.State,
	}
	if user.Department != nil {
		u.Department = &orderModel.Department{
			Id:   user.Department.Id,
			Code: user.Department.Code,
			Name: user.Department.Name,
			Type: user.Department.Type,
		}
	}
	return u, nil
}
