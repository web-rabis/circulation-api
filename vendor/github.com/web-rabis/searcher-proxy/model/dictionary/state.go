package dictionary

import "github.com/web-rabis/searcher-proxy/protobuf"

type State struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewStateFromProto(s *protobuf.State) *State {
	if s == nil {
		return nil
	}
	return &State{
		Id:   s.Id,
		Code: s.Code,
		Name: s.Name,
	}
}
