package model

import "github.com/web-rabis/order-client/protobuf"

type Department struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type State struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ReasonRejection struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewDepartmentFromProto(d *protobuf.Department) *Department {
	if d == nil {
		return nil
	}
	return &Department{
		Id:   d.Id,
		Code: d.Code,
		Name: d.Name,
		Type: d.Type,
	}
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
func NewReasonRejectionProto(s *protobuf.ReasonRejection) *ReasonRejection {
	if s == nil {
		return nil
	}
	return &ReasonRejection{
		Id:   s.Id,
		Code: s.Code,
		Name: s.Name,
	}
}
