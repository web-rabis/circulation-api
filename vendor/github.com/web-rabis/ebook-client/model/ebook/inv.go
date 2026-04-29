package ebook

import (
	"github.com/web-rabis/ebook-client/model/dictionary"
	"github.com/web-rabis/ebook-client/protobuf"
)

type Inv struct {
	Id         int64                  `json:"id"`
	EbookId    int64                  `json:"ebookId"`
	InvNumber  string                 `json:"invNumber"`
	Barcode    string                 `json:"barcode"`
	Department *dictionary.Department `json:"department"`
	State      *dictionary.State      `json:"state"`
}

func NewInvFromProto(i *protobuf.Inv) *Inv {
	if i == nil {
		return nil
	}
	return &Inv{
		Id:         i.Id,
		EbookId:    i.EbookId,
		InvNumber:  i.InvNumber,
		Barcode:    i.Barcode,
		Department: dictionary.NewDepartmentFromProto(i.Department),
		State:      dictionary.NewStateFromProto(i.State),
	}
}
func NewInvListFromProto(s []*protobuf.Inv) []*Inv {
	if s == nil {
		return nil
	}
	var result []*Inv
	for _, v := range s {
		result = append(result, NewInvFromProto(v))
	}
	return result
}
