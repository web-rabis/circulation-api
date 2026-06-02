package dictionary

import "github.com/web-rabis/searcher-proxy/protobuf"

type TypeDescription struct {
	Id     int64  `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	NameKz string `json:"nameKz"`
}

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
