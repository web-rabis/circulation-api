package dictionary

import (
	"github.com/web-rabis/searcher-proxy/protobuf"
)

type BibliographicLevel struct {
	Id         int64  `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	TypeEbooks string `json:"typeEbooks"`
}

func NewBibliographicLevelFromProto(b *protobuf.BibliographicLevel) *BibliographicLevel {
	if b == nil {
		return nil
	}
	return &BibliographicLevel{
		Id:         b.Id,
		Code:       b.Code,
		Name:       b.Name,
		TypeEbooks: b.TypeEbooks,
	}
}
