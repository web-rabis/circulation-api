package dto

import (
	"github.com/web-rabis/circulation-api/internal/domain/model"
)

type OrdersResponse struct {
	Result []*model.Order `json:"result"`
	Count  int64          `json:"count"`
}
