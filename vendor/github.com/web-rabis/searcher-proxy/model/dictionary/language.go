package dictionary

import "github.com/web-rabis/searcher-proxy/protobuf"

type Language struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewLanguageFromProto(s *protobuf.DictionaryLanguage) *Language {
	if s == nil {
		return nil
	}
	return &Language{
		Id:   s.Id,
		Code: s.Code,
		Name: s.Name,
	}
}
