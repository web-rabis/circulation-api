package model

import (
	"time"

	"github.com/web-rabis/order-client/protobuf"
)

type OrderAudit struct {
	Id         int64       `json:"id"`
	OrderId    int64       `json:"orderId"`
	AuditDate  time.Time   `json:"auditDate"`
	State      *State      `json:"state"`
	UserId     int64       `json:"user"`
	Department *Department `json:"department"`
}

func NewAuditRecordFromProto(a *protobuf.OrderAudit) *OrderAudit {
	if a == nil {
		return nil
	}
	return &OrderAudit{
		Id:         a.Id,
		OrderId:    a.OrderId,
		AuditDate:  a.AuditDate.AsTime(),
		State:      NewStateFromProto(a.State),
		UserId:     a.UserId,
		Department: NewDepartmentFromProto(a.Department),
	}
}
