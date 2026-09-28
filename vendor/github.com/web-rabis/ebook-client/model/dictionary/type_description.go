package dictionary

import "github.com/web-rabis/ebook-client/protobuf"

type TypeDescription struct {
	Id     int64  `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	NameKz string `json:"nameKz"`
}
type TypeDescriptionFilters struct{}

func NewTypeDescriptionFromProto(t *protobuf.TypeDescription) *TypeDescription {
	if t == nil {
		return nil
	}
	return &TypeDescription{
		Id:     t.Id,
		Code:   t.Code,
		Name:   t.Name,
		NameKz: t.NameKz,
	}
}
func (t *TypeDescription) ToProto() *protobuf.TypeDescription {
	if t == nil {
		return nil
	}
	return &protobuf.TypeDescription{
		Id:     t.Id,
		Code:   t.Code,
		Name:   t.Name,
		NameKz: t.NameKz,
	}
}
func (f *TypeDescriptionFilters) ToProto() *protobuf.TypeDescriptionFilters {
	if f == nil {
		return nil
	}
	return &protobuf.TypeDescriptionFilters{}
}
