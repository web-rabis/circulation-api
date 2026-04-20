package dto

import "github.com/web-rabis/order-client/model"

type RejectionReason struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewRejectionReasonResponse(result []*model.ReasonRejection) []RejectionReason {
	var rr = make([]RejectionReason, len(result))
	for i, r := range result {
		rr[i] = RejectionReason{
			Id:   r.Id,
			Code: r.Code,
			Name: r.Name,
		}
	}
	return rr
}
