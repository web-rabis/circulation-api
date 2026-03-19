package model

import (
	"time"

	"github.com/web-rabis/order-client/protobuf"
)

const (
	OrderStateOrdered         = "Order.Ordered"
	OrderStateInStorage       = "Order.InStorage"
	OrderStateInReadingHall   = "Order.InReadingHall"
	OrderStateInHands         = "Order.InHands"
	OrderStatePostponed       = "Order.Postponed"
	OrderStateReturnToStorage = "Order.ReturnToStorage"
	OrderStateProcessed       = "Order.Processed"
	OrderStateRejected        = "Order.Rejected"
	OrderStateReaderReturned  = "Order.ReaderReturned"
	OrderStateDeleted         = "Order.Deleted"
)

type Order struct {
	Id         int64       `json:"id"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	Type       string      `json:"type"`
	Reader     *Reader     `json:"reader"`
	Ebook      *Ebook      `json:"ebook"`
	InvNumber  *EbookInv   `json:"invNumber"`
	Periodical *Periodical `json:"periodical"`
	State      *State      `json:"state"`
	Department *Department `json:"department"`
}

func NewOrderFormProto(o *protobuf.Order) *Order {
	if o == nil {
		return nil
	}
	return &Order{
		Id:         o.Id,
		CreatedAt:  o.CreatedAt.AsTime(),
		UpdatedAt:  o.UpdatedAt.AsTime(),
		Type:       o.Type,
		Reader:     NewReaderFromProto(o.Reader),
		Ebook:      NewEbookFromProto(o.Ebook),
		InvNumber:  NewEbookInvFromProto(o.InvNumber),
		Periodical: NewPeriodicalFromProto(o.Periodical),
		State:      NewStateFromProto(o.State),
		Department: NewDepartmentFromProto(o.Department),
	}
}
