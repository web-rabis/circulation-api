package model

import (
	"github.com/web-rabis/order-client/protobuf"
)

type OrderFilters struct {
	EbookId         int64    `json:"merchantId"`
	PeriodicalId    int64    `json:"terminalId"`
	TicketNumber    int64    `json:"ticketNumber"`
	States          []string `json:"states"`
	DepartmentId    int64    `json:"departmentId"`
	Period          string   `json:"period"`
	IsAuxiliaryFund *bool    `json:"isAuxiliaryFund"`
}
type StateCountFilters struct {
	States       []string `json:"states"`
	DepartmentId int64    `json:"departmentId"`
	Period       string   `json:"period"`
}
type ReasonRejectionFilters struct{}
type DepartmentFilters struct{}

func (f *OrderFilters) ToProto() *protobuf.OrderFilters {
	if f == nil {
		return nil
	}
	of := &protobuf.OrderFilters{
		EbookId:         f.EbookId,
		PeriodicalId:    f.PeriodicalId,
		TicketNumber:    f.TicketNumber,
		States:          f.States,
		DepartmentId:    f.DepartmentId,
		Period:          f.Period,
		IsAuxiliaryFund: -1,
	}
	if f.IsAuxiliaryFund != nil {
		if *f.IsAuxiliaryFund {
			of.IsAuxiliaryFund = 1
		} else {
			of.IsAuxiliaryFund = 0
		}
	}
	return of
}
func (f *StateCountFilters) ToProto() *protobuf.StateCountFilters {
	if f == nil {
		return nil
	}
	return &protobuf.StateCountFilters{
		States:       f.States,
		DepartmentId: f.DepartmentId,
		Period:       f.Period,
	}
}
func (f *ReasonRejectionFilters) ToProto() *protobuf.ReasonRejectionFilters {
	if f == nil {
		return nil
	}
	return &protobuf.ReasonRejectionFilters{}
}
func (f *DepartmentFilters) ToProto() *protobuf.DepartmentFilters {
	if f == nil {
		return nil
	}
	return &protobuf.DepartmentFilters{}
}
