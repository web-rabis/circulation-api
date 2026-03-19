package dto

import (
	orderModel "github.com/web-rabis/order-client/model"
)

type OrdersResponse struct {
	Result []*orderModel.Order `json:"result"`
	Count  int64               `json:"count"`
}
