package dictionary

import "github.com/web-rabis/ebook-client/protobuf"

type Department struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Type string
}

func NewDepartmentFromProto(d *protobuf.Department) *Department {
	if d == nil {
		return nil
	}
	return &Department{
		Id:   d.Id,
		Code: d.Code,
		Name: d.Name,
		Type: d.Type,
	}
}
