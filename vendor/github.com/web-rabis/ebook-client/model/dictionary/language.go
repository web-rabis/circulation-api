package dictionary

import "github.com/web-rabis/ebook-client/protobuf"

type Language struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
type LanguageFilters struct {
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
func (f *LanguageFilters) ToProto() *protobuf.DictionaryLanguageFilters {
	if f == nil {
		return nil
	}
	return &protobuf.DictionaryLanguageFilters{}
}
