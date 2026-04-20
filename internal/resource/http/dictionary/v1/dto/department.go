package dto

import "github.com/web-rabis/order-client/model"

type Department struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func NewDepartmentResponse(result []*model.Department) []Department {
	var rr = make([]Department, len(result))
	for i, r := range result {
		rr[i] = Department{
			Id:   r.Id,
			Code: r.Code,
			Name: r.Name,
			Type: r.Type,
		}
	}
	return rr
}
