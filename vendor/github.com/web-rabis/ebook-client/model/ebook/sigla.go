package ebook

import (
	"github.com/web-rabis/ebook-client/model/dictionary"
	"github.com/web-rabis/ebook-client/protobuf"
)

type Sigla struct {
	Id      int64             `json:"id"`
	EbookId int64             `json:"ebookId"`
	Sigla   *dictionary.Sigla `json:"sigla"`
}

func NewSiglaFromProto(i *protobuf.Sigla) *Sigla {
	if i == nil {
		return nil
	}
	return &Sigla{
		Id:      i.Id,
		EbookId: i.EbookId,
		Sigla:   dictionary.NewSiglaFromProto(i.Sigla),
	}
}
func NewSiglaListFromProto(s []*protobuf.Sigla) []*Sigla {
	if s == nil {
		return nil
	}
	var result []*Sigla
	for _, v := range s {
		result = append(result, NewSiglaFromProto(v))
	}
	return result
}
