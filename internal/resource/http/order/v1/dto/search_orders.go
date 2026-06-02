package dto

import "github.com/web-rabis/circulation-api/internal/domain/model"

type SearchOrdersRequest struct {
	Query string `json:"query"`
}

type SearchOrdersResponse struct {
	Result []*model.Order `json:"result"`
	Count  int64          `json:"count"`
}
