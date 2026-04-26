package model

import (
	"time"

	"github.com/web-rabis/ebook-client/model"
	orderModel "github.com/web-rabis/order-client/model"
)

type Order struct {
	Id                int64                       `json:"id"`
	CreatedAt         time.Time                   `json:"createdAt"`
	UpdatedAt         time.Time                   `json:"updatedAt"`
	Type              string                      `json:"type"`
	Reader            *orderModel.Reader          `json:"reader"`
	Ebook             *model.Ebook                `json:"ebook"`
	InvNumber         *orderModel.EbookInv        `json:"invNumber"`
	Periodical        *orderModel.Periodical      `json:"periodical"`
	State             *orderModel.State           `json:"state"`
	Department        *orderModel.Department      `json:"department"`
	StorageDepartment *orderModel.Department      `json:"storageDepartment"`
	IsAuxiliaryFund   bool                        `json:"isAuxiliaryFund"`
	ReasonRejection   *orderModel.ReasonRejection `json:"reasonRejection"`
}

func NewOrder(o *orderModel.Order, e *model.Ebook) *Order {
	if o == nil {
		return nil
	}
	if e == nil && o.Type == "ebook" && o.Ebook != nil {
		e = &model.Ebook{
			Id:     o.Ebook.Id,
			Author: o.Ebook.Author,
			Title:  o.Ebook.Title,
		}
	}
	return &Order{
		Id:                o.Id,
		CreatedAt:         o.CreatedAt,
		UpdatedAt:         o.UpdatedAt,
		Type:              o.Type,
		Reader:            o.Reader,
		Ebook:             e,
		InvNumber:         o.InvNumber,
		Periodical:        o.Periodical,
		State:             o.State,
		Department:        o.Department,
		StorageDepartment: o.StorageDepartment,
		IsAuxiliaryFund:   o.IsAuxiliaryFund,
		ReasonRejection:   o.ReasonRejection,
	}

}
