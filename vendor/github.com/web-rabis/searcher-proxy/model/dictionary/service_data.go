package dictionary

import "github.com/web-rabis/searcher-proxy/protobuf"

type ServiceData struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewServiceDataFromProto(s *protobuf.DictionaryServiceData) *ServiceData {
	if s == nil {
		return nil
	}
	return &ServiceData{
		Id:   s.Id,
		Code: s.Code,
		Name: s.Name,
	}
}
