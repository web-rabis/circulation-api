package ebook_form

import "github.com/web-rabis/ebook-client/protobuf"

// DictionaryRef - ссылка на значение справочника.
// Code и Name заполняются только на чтении, на записи игнорируются.
type DictionaryRef struct {
	Id   int64  `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewDictionaryRefFromProto(r *protobuf.DictionaryRef) *DictionaryRef {
	if r == nil {
		return nil
	}
	return &DictionaryRef{Id: r.Id, Code: r.Code, Name: r.Name}
}

func (r *DictionaryRef) ToProto() *protobuf.DictionaryRef {
	if r == nil {
		return nil
	}
	return &protobuf.DictionaryRef{Id: r.Id, Code: r.Code, Name: r.Name}
}
