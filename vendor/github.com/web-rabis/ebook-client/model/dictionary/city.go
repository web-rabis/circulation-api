package dictionary

import "github.com/web-rabis/ebook-client/protobuf"

type City struct {
	Id   int64  `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewCityFromProto(s *protobuf.DictionaryCity) *City {
	if s == nil {
		return nil
	}
	return &City{
		Id:   s.Id,
		Code: s.Code,
		Name: s.Name,
	}
}
