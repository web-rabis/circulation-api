package model

import "github.com/web-rabis/ebook-client/protobuf"

type EbookInv struct {
	Id         int64       `json:"id"`
	EbookId    int64       `json:"ebookId"`
	InvNumber  string      `json:"invNumber"`
	Barcode    string      `json:"barcode"`
	Department *Department `json:"department"`
	State      *State      `json:"state"`
}

func NewEbookInvFromProto(i *protobuf.EbookInv) *EbookInv {
	if i == nil {
		return nil
	}
	return &EbookInv{
		Id:         i.Id,
		EbookId:    i.EbookId,
		InvNumber:  i.InvNumber,
		Barcode:    i.Barcode,
		Department: NewDepartmentFromProto(i.Department),
		State:      NewStateFromProto(i.State),
	}
}
func NewEbookInvsFromProto(s []*protobuf.EbookInv) []*EbookInv {
	if s == nil {
		return nil
	}
	var result []*EbookInv
	for _, v := range s {
		result = append(result, NewEbookInvFromProto(v))
	}
	return result
}
