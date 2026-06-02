package search

import (
	"context"

	orderClient "github.com/web-rabis/order-client/client"
	orderModel "github.com/web-rabis/order-client/model"
	searcherClient "github.com/web-rabis/searcher-proxy/client"
	"github.com/web-rabis/searcher-proxy/model"
	ebookModel "github.com/web-rabis/searcher-proxy/model/ebook"
	periodicalModel "github.com/web-rabis/searcher-proxy/model/periodical"
)

type IManager interface {
	EbookSimpleSearch(ctx context.Context, query string, paging *model.Paging) (int64, []*ebookModel.Ebook, error)
	PeriodicalSimpleSearch(ctx context.Context, query string, paging *model.Paging) (int64, []*periodicalModel.Periodical, error)
}
type Manager struct {
	orderCl      orderClient.OrderService
	searchClient searcherClient.Base
}

func NewManager(orderCl orderClient.OrderService, searchClient searcherClient.Base) *Manager {
	return &Manager{
		orderCl:      orderCl,
		searchClient: searchClient,
	}
}
func (m *Manager) EbookSimpleSearch(ctx context.Context, query string, paging *model.Paging) (int64, []*ebookModel.Ebook, error) {
	filters := &ebookModel.EbookSearchFilters{
		Query:          query,
		Type:           "BEGIN",
		IsSimpleSearch: true,
	}
	count, ebooks, err := m.searchClient.EbookSearchSvc().Search(ctx, filters, paging)
	if err != nil {
		return 0, nil, err
	}
	for _, ebook := range ebooks {
		ebook.Inv, _ = m.filterInv(ctx, ebook.Id, ebook.Inv)
	}
	return count, ebooks, nil
}
func (m *Manager) PeriodicalSimpleSearch(ctx context.Context, query string, paging *model.Paging) (int64, []*periodicalModel.Periodical, error) {
	filters := &periodicalModel.PeriodicalSearchFilters{
		Query:          query,
		Type:           "BEGIN",
		IsSimpleSearch: true,
	}
	return m.searchClient.PeriodicalSearchSvc().Search(ctx, filters, paging)
}
func (m *Manager) filterInv(ctx context.Context, ebookId int64, inv []*ebookModel.Inv) ([]*ebookModel.Inv, error) {
	states := []string{
		orderModel.OrderStateInHands,
		orderModel.OrderStateOrdered,
		orderModel.OrderStateInStorage,
		orderModel.OrderStateInReadingHall,
		orderModel.OrderStatePostponed,
		orderModel.OrderStateInAuxiliaryFund,
	}
	filters := &orderModel.OrderFilters{
		EbookId: ebookId,
		States:  states,
	}
	count, orders, err := m.orderCl.List(ctx, nil, filters)
	if err == nil && count > 0 {
		var inv_ []*ebookModel.Inv
		for _, cinv := range inv {
			var invFounded bool
			for _, order := range orders {
				if order.InvNumber != nil && order.InvNumber.Id == cinv.Id {
					invFounded = true
					break
				}
			}
			if !invFounded {
				inv_ = append(inv_, cinv)
			}
		}
		return inv_, nil
	}
	return inv, nil
}
