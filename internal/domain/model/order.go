package model

import (
	"time"

	"github.com/web-rabis/ebook-client/model/ebook"
	orderModel "github.com/web-rabis/order-client/model"
	readerModel "github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/sso-client/model"
)

type Order struct {
	Id                int64                       `json:"id"`
	CreatedAt         time.Time                   `json:"createdAt"`
	UpdatedAt         time.Time                   `json:"updatedAt"`
	Type              string                      `json:"type"`
	Reader            *readerModel.Reader         `json:"reader"`
	Ebook             *ebook.Ebook                `json:"ebook"`
	InvNumber         *ebook.Inv                  `json:"invNumber"`
	Periodical        *orderModel.Periodical      `json:"periodical"`
	State             *orderModel.State           `json:"state"`
	Department        *orderModel.Department      `json:"department"`
	StorageDepartment *orderModel.Department      `json:"storageDepartment"`
	IsAuxiliaryFund   bool                        `json:"isAuxiliaryFund"`
	ReasonRejection   *orderModel.ReasonRejection `json:"reasonRejection"`
}
type OrderAudit struct {
	Id         int64                  `json:"id"`
	OrderId    int64                  `json:"orderId"`
	AuditDate  time.Time              `json:"auditDate"`
	State      *orderModel.State      `json:"state"`
	User       *model.User            `json:"user"`
	Department *orderModel.Department `json:"department"`
}

// NewOrder собирает ответ шлюза: заказ из order-client (только идентификаторы
// EbookId/EbookInvId/ReaderTicketNumber и денормализованные InvNumber/Barcode)
// обогащается карточкой книги e и экземпляром inv из сервиса ebook, карточкой
// читателя reader из reader-сервиса. Если ebook недоступен, в ответ уходит
// минимальная заглушка по данным самого заказа, чтобы список заказов
// оставался работоспособным.
func NewOrder(o *orderModel.Order, reader *readerModel.Reader, e *ebook.Ebook, inv *ebook.Inv) *Order {
	if o == nil {
		return nil
	}
	if e == nil && o.EbookId != 0 {
		e = &ebook.Ebook{Id: o.EbookId}
	}
	if inv == nil && (o.EbookInvId != 0 || o.InvNumber != "" || o.Barcode != "") {
		inv = &ebook.Inv{
			Id:        o.EbookInvId,
			EbookId:   o.EbookId,
			InvNumber: o.InvNumber,
			Barcode:   o.Barcode,
		}
	}
	return &Order{
		Id:                o.Id,
		CreatedAt:         o.CreatedAt,
		UpdatedAt:         o.UpdatedAt,
		Type:              o.Type,
		Reader:            reader,
		Ebook:             e,
		InvNumber:         inv,
		Periodical:        o.Periodical,
		State:             o.State,
		Department:        o.Department,
		StorageDepartment: o.StorageDepartment,
		IsAuxiliaryFund:   o.IsAuxiliaryFund,
		ReasonRejection:   o.ReasonRejection,
	}

}
func NewOrderAudit(oa *orderModel.OrderAudit) *OrderAudit {
	if oa == nil {
		return nil
	}
	var u *model.User
	if oa.UserId != 0 {
		u = &model.User{Id: oa.UserId}
	}
	return &OrderAudit{
		Id:         oa.Id,
		OrderId:    oa.OrderId,
		AuditDate:  oa.AuditDate,
		State:      oa.State,
		User:       u,
		Department: oa.Department,
	}
}
