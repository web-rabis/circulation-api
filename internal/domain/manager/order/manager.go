package order

import (
	"context"
	"errors"
	"fmt"

	"github.com/web-rabis/circulation-api/internal/domain/model"
	ebookClient "github.com/web-rabis/ebook-client/client"
	ebookFilters "github.com/web-rabis/ebook-client/model"
	ebookModel "github.com/web-rabis/ebook-client/model/ebook"
	orderClient "github.com/web-rabis/order-client/client"
	orderModel "github.com/web-rabis/order-client/model"
	readerClient "github.com/web-rabis/reader-client/client"
	readerModel "github.com/web-rabis/reader-client/model"
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
	Issue(ctx context.Context, id int64, userId, invId int64, barcode string) error
	IssueOrders(ctx context.Context, id []int64, userId int64) error
	Redirect(ctx context.Context, ids []int64, departmentId, userId int64) error
	Create(ctx context.Context, ticketNumber, ebookId, departmentId, invId, userId int64) (int64, error)
	Search(ctx context.Context, query string, paging *orderModel.Paging) (int64, []*model.Order, error)
	Audit(ctx context.Context, orderId int64) ([]*model.OrderAudit, error)
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
	// order-client отдаёт только идентификаторы книги и экземпляра, карточка
	// подтягивается здесь из сервиса ebook. Кэш на выборку, чтобы не дёргать
	// ebook по разу на каждый заказ с одной и той же книгой.
	var (
		briefs = make(map[int64]*ebookModel.Ebook)
		invs   = make(map[int64]map[int64]*ebookModel.Inv)
	)
	var orders_ = make([]*model.Order, len(orders))
	for i, order := range orders {
		var reader *readerModel.Reader
		if order.ReaderTicketNumber != 0 {
			reader, err = m.readerCl.ReaderById(ctx, order.ReaderTicketNumber)
			if err != nil {
				reader = &readerModel.Reader{TicketNumber: order.ReaderTicketNumber}
			}
		}
		var e *ebookModel.Ebook
		var inv *ebookModel.Inv
		if order.EbookId != 0 {
			var ok bool
			if e, ok = briefs[order.EbookId]; !ok {
				e, _ = m.ebookCl.EbookById(ctx, order.EbookId, false)
				briefs[order.EbookId] = e
			}
			if order.EbookInvId != 0 {
				byId, ok := invs[order.EbookId]
				if !ok {
					byId = make(map[int64]*ebookModel.Inv)
					if _, list, err := m.ebookCl.InvList(ctx, &ebookFilters.InvFilters{EbookId: order.EbookId}, nil); err == nil {
						for _, v := range list {
							byId[v.Id] = v
						}
					}
					invs[order.EbookId] = byId
				}
				inv = byId[order.EbookInvId]
			}
		}
		orders_[i] = model.NewOrder(order, reader, e, inv)
	}
	return count, orders_, nil
}
func (m *Manager) ById(ctx context.Context, id int64) (*orderModel.Order, error) {
	return m.orderCl.ById(ctx, id)
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
func (m *Manager) Issue(ctx context.Context, id int64, userId, invId int64, barcode string) error {
	user, err := m.getUserById(ctx, userId)
	if err != nil {
		return err
	}
	if barcode != "" {
		if _, err := m.ebookCl.InvBarcodeSave(ctx, invId, barcode); err != nil {
			return err
		}
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
func (m *Manager) Create(ctx context.Context, ticketNumber, ebookId, departmentId, invId, userId int64) (int64, error) {
	// Проверяем, не выдана ли уже эта книга данному читателю (активный заказ)
	activeStates := []string{
		orderModel.OrderStateOrdered,
		orderModel.OrderStateInStorage,
		orderModel.OrderStateInReadingHall,
		orderModel.OrderStateInHands,
		orderModel.OrderStatePostponed,
		orderModel.OrderStateInAuxiliaryFund,
	}
	count, _, err := m.orderCl.List(ctx, &orderModel.Paging{Limit: 1}, &orderModel.OrderFilters{
		TicketNumber: ticketNumber,
		EbookId:      ebookId,
		States:       activeStates,
	})
	if err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, errors.New("книга уже выдана данному читателю")
	}

	user, err := m.getUserById(ctx, userId)
	if err != nil {
		return 0, err
	}
	return m.orderCl.CreateOnDemanIssue(ctx, ticketNumber, ebookId, departmentId, invId, user)
}

// Search ищет заказы по строке query: если query — число, ищет по номеру читательского билета,
// иначе ищет по номеру заказа/названию (передаётся как query в фильтр).
func (m *Manager) Search(ctx context.Context, query string, paging *orderModel.Paging) (int64, []*model.Order, error) {
	filters := &orderModel.OrderFilters{}
	var ticketNumber int64
	if _, err := fmt.Sscanf(query, "%d", &ticketNumber); err == nil {
		filters.TicketNumber = ticketNumber
	} else {
		filters.Query = query
	}
	return m.List(ctx, filters, paging)
}

func (m *Manager) Audit(ctx context.Context, orderId int64) ([]*model.OrderAudit, error) {

	list, err := m.orderCl.Audit(ctx, orderId)
	if err != nil {
		return nil, err
	}
	var audits = make([]*model.OrderAudit, len(list))
	for i, audit := range list {
		audits[i] = model.NewOrderAudit(audit)
		if audit.UserId != 0 {
			audits[i].User, _ = m.userCl.ById(ctx, audit.UserId)
		}
	}
	return audits, nil
}

func (m *Manager) getUserById(ctx context.Context, id int64) (*orderModel.User, error) {
	user, err := m.userCl.ById(ctx, id)
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
