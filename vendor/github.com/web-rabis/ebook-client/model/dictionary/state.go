package dictionary

import "github.com/web-rabis/ebook-client/protobuf"

type StateFilters struct {
	CodeLike string
}
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
func (s *State) ToProto() *protobuf.State {
	if s == nil {
		return nil
	}
	return &protobuf.State{
		Id:   s.Id,
		Code: s.Code,
		Name: s.Name,
	}
}
func (f *StateFilters) ToProto() *protobuf.StateFilters {
	if f == nil {
		return nil
	}
	return &protobuf.StateFilters{
		CodeLike: f.CodeLike,
	}
}
