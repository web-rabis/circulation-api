package dto

type RejectOrderRequest struct {
	Ids            []int64 `json:"ids"`
	ReasonRejectId int64   `json:"reasonRejectId"`
}
