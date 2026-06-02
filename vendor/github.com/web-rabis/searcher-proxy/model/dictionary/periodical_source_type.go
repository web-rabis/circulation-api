package dictionary

import "github.com/web-rabis/searcher-proxy/protobuf"

type PeriodicalSourceType struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewPeriodicalSourceTypeFromProto(s *protobuf.DictionaryPeriodicalSourceType) *PeriodicalSourceType {
	if s == nil {
		return nil
	}
	return &PeriodicalSourceType{
		Id:   s.Id,
		Code: s.Code,
		Name: s.Name,
	}
}
