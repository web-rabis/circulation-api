package dictionary

import "github.com/web-rabis/searcher-proxy/protobuf"

type SearchField struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func NewSearchFieldFromProto(s *protobuf.SearchField) *SearchField {
	if s == nil {
		return nil
	}
	return &SearchField{
		Id:   s.Id,
		Name: s.Name,
	}
}
