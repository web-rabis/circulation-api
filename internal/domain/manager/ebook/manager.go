package ebook

import (
	"context"

	ebookClient "github.com/web-rabis/ebook-client/client"
	ebookModel "github.com/web-rabis/ebook-client/model/ebook"
	orderClient "github.com/web-rabis/order-client/client"
	orderModel "github.com/web-rabis/order-client/model"
)

type IManager interface {
	EbookBriefById(ctx context.Context, id int64) (*ebookModel.EbookBrief, error)
	EbookCardById(ctx context.Context, id int64) (*ebookModel.EbookCard, error)
}
type Manager struct {
	ebookCl ebookClient.EbookService
	orderCl orderClient.OrderService
}

func NewManager(ebookCl ebookClient.EbookService, orderCl orderClient.OrderService) *Manager {
	return &Manager{
		ebookCl: ebookCl,
	}
}
func (m *Manager) EbookBriefById(ctx context.Context, id int64) (*ebookModel.EbookBrief, error) {
	return m.ebookCl.EbookBriefById(ctx, id)
}
func (m *Manager) EbookCardById(ctx context.Context, id int64) (*ebookModel.EbookCard, error) {
	card, err := m.ebookCl.EbookCardById(ctx, id)
	if err != nil {
		return nil, err
	}
	states := []string{
		orderModel.OrderStateInHands,
		orderModel.OrderStateOrdered,
		orderModel.OrderStateInStorage,
		orderModel.OrderStateInReadingHall,
		orderModel.OrderStatePostponed,
		orderModel.OrderStateInAuxiliaryFund,
	}
	filters := &orderModel.OrderFilters{
		EbookId: id,
		States:  states,
	}
	count, orders, err := m.orderCl.List(ctx, nil, filters)
	if err == nil && count > 0 {
		var inv []*ebookModel.Inv
		for _, cinv := range card.Inv {
			var invFounded bool
			for _, order := range orders {
				if order.InvNumber.Id == cinv.Id {
					invFounded = true
					break
				}
			}
			if !invFounded {
				inv = append(inv, cinv)
			}
		}
		card.Inv = inv
	}

	return card, nil
}

//func (m *Manager) getFreeCopies(ctx context.Context, e *ebook.Ebook) (int64, error) {
//	if e.BibliographicLevel.Code == "СК" && e.ParentId != nil {
//		eParent, err := m.ebookRepo.EbookById(ctx, *e.ParentId)
//		if err != nil {
//			return 0, err
//		}
//		return m.getFreeBookCopies(ctx, eParent)
//	}
//	return m.getFreeBookCopies(ctx, e)
//}
//func (m *Manager) getFreeBookCopies(ctx context.Context, e *ebook.Ebook) (int64, error) {
//	orderFilter := &model.OrderFilter{
//		EbookId: &e.Id,
//		InStateCodes: []string{
//			order.OrderStateInHands,
//			order.OrderStatePostponed,
//			order.OrderStateInReadingHall,
//			order.OrderStateReturnToStorage,
//			order.OrderStateReaderReturned,
//		},
//	}
//	orders, err := m.orderRepo.List(ctx, orderFilter, nil)
//	if err != nil {
//		return 0, err
//	}
//
//	var invState2Count int64 = 0
//	for _, inv := range e.Inv {
//		inv.Availability = &dictionary.State{
//			Id:   0,
//			Code: "AVAILABLE",
//			Name: "Свободен",
//		}
//		if inv.State != nil && inv.State.Code == ebook.EbookInvState2Code {
//			invState2Count++
//			inv.Availability = &dictionary.State{
//				Id:   0,
//				Code: "NOT_AVAILABLE",
//				Name: "Занят",
//			}
//		}
//		for _, order := range orders {
//			if inv.Id == order.InvNumber.Id {
//				inv.Availability = &dictionary.State{
//					Id:   0,
//					Code: "NOT_AVAILABLE",
//					Name: "Занят",
//				}
//				break
//			}
//		}
//	}
//	return int64(len(e.Inv)) - int64(len(orders)) - invState2Count, nil
//}
