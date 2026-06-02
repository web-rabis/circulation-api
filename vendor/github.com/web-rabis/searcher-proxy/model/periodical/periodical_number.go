package periodical

import (
	"time"

	"github.com/web-rabis/searcher-proxy/model/dictionary"
	"github.com/web-rabis/searcher-proxy/protobuf"
)

type PeriodicalNumber struct {
	Id           int64                            `json:"id,omitempty"`
	PeriodicalId int64                            `json:"periodicalId,omitempty"`
	Number       string                           `json:"number,omitempty"`
	NumberDate   *time.Time                       `json:"numberDate,omitempty"`
	ReceiptDate  *time.Time                       `json:"receiptDate,omitempty"`
	Count        int64                            `json:"count"`
	Price        float64                          `json:"price"`
	Path         string                           `json:"path"`
	Type         *dictionary.PeriodicalSourceType `json:"type,omitempty"`
	Department   *dictionary.Department           `json:"department,omitempty"`
}

func NewPeriodicalNumberFromProto(s *protobuf.PeriodicalNumber) *PeriodicalNumber {
	if s == nil {
		return nil

	}
	var numberDate *time.Time
	var receiptDate *time.Time
	if s.NumberDate != nil {
		t := s.NumberDate.AsTime()
		numberDate = &t
	}
	if s.ReceiptDate != nil {
		t := s.ReceiptDate.AsTime()
		receiptDate = &t
	}
	return &PeriodicalNumber{
		Id:           s.Id,
		PeriodicalId: s.PeriodicalId,
		Number:       s.Number,
		NumberDate:   numberDate,
		ReceiptDate:  receiptDate,
		Count:        s.Count,
		Price:        s.Price,
		Path:         s.Path,
		Type:         dictionary.NewPeriodicalSourceTypeFromProto(s.Type),
		Department:   dictionary.NewDepartmentFromProto(s.Department),
	}
}
