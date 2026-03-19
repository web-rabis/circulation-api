package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/web-rabis/order-client/client"
	"github.com/web-rabis/order-client/model"
	"github.com/web-rabis/order-client/protobuf"
)

type OrderService struct {
	client protobuf.OrderServiceClient
}

var _ client.OrderService = &OrderService{}

func NewOrderServiceClient(client protobuf.OrderServiceClient) client.OrderService {
	return &OrderService{
		client: client,
	}
}

func (c *OrderService) List(ctx context.Context, paging *model.Paging, filters *model.OrderFilters) (int64, []*model.Order, error) {
	request := &protobuf.OrderListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	response, err := c.client.List(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		orders := make([]*model.Order, len(response.GetOrders()))
		for i, o := range response.GetOrders() {
			orders[i] = model.NewOrderFormProto(o)
		}
		return response.Count, orders, nil
	default:
		return 0, nil, err
	}
}
func (c *OrderService) ById(ctx context.Context, id int64) (*model.Order, error) {
	request := &protobuf.ByIdRequest{
		Id: id,
	}
	response, err := c.client.ById(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return model.NewOrderFormProto(response), nil
	default:
		return nil, err
	}
}
