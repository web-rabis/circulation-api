package model

import (
	"github.com/web-rabis/order-client/protobuf"
)

type OrderFilters struct {
	EbookId      int64    `json:"merchantId"`
	PeriodicalId int64    `json:"terminalId"`
	TicketNumber int64    `json:"ticketNumber"`
	States       []string `json:"states"`
}

func (f *OrderFilters) ToProto() *protobuf.OrderFilters {
	if f == nil {
		return nil
	}
	return &protobuf.OrderFilters{
		EbookId:      f.EbookId,
		PeriodicalId: f.PeriodicalId,
		TicketNumber: f.TicketNumber,
		States:       f.States,
	}
}
