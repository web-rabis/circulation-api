package dictionary

import "github.com/web-rabis/ebook-client/protobuf"

type BibliographicLevel struct {
	Id         int64  `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	TypeEbooks string `json:"typeEbooks"`
}
type BibliographicLevelFilters struct {
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
func (f *BibliographicLevelFilters) ToProto() *protobuf.BibliographicLevelFilters {
	if f == nil {
		return nil
	}
	return &protobuf.BibliographicLevelFilters{}
}
