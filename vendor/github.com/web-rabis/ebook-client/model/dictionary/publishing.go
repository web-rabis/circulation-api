package dictionary

import "github.com/web-rabis/ebook-client/protobuf"

type Publishing struct {
	Id   int64  `json:"id,omitempty"`
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewPublishingFromProto(s *protobuf.DictionaryPublishing) *Publishing {
	if s == nil {
		return nil
	}
	return &Publishing{
		Id:   s.Id,
		Code: s.Code,
		Name: s.Name,
	}
}
