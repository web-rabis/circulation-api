package dictionary

import "github.com/web-rabis/searcher-proxy/protobuf"

type Sigla struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewSiglaFromProto(s *protobuf.DictionarySigla) *Sigla {
	if s == nil {
		return nil
	}
	return &Sigla{
		Id:   s.Id,
		Code: s.Code,
		Name: s.Name,
	}
}
