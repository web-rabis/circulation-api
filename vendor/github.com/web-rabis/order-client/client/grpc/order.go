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

func (c *OrderService) Redirect(ctx context.Context, ids []int64, departmentId int64, user *model.User) error {
	request := &protobuf.RedirectRequest{
		Ids:          ids,
		DepartmentId: departmentId,
		User:         user.ToProto(),
	}
	_, err := c.client.Redirect(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
func (c *OrderService) CancelReject(ctx context.Context, ids []int64, user *model.User) error {
	request := &protobuf.CancelRejectRequest{
		Ids:  ids,
		User: user.ToProto(),
	}
	_, err := c.client.CancelReject(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
func (c *OrderService) Reject(ctx context.Context, ids []int64, reasonRejectId int64, user *model.User) error {
	request := &protobuf.RejectRequest{
		Ids:               ids,
		ReasonRejectionId: reasonRejectId,
		User:              user.ToProto(),
	}
	_, err := c.client.Reject(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
func (c *OrderService) Postponed(ctx context.Context, ids []int64, user *model.User) error {
	request := &protobuf.PostponedRequest{
		Ids:  ids,
		User: user.ToProto(),
	}
	_, err := c.client.Postponed(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
func (c *OrderService) Return(ctx context.Context, ids []int64, user *model.User) error {
	request := &protobuf.ReturnRequest{
		Ids:  ids,
		User: user.ToProto(),
	}
	_, err := c.client.Return(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
func (c *OrderService) Issue(ctx context.Context, ids []model.IssueOrder, user *model.User) error {
	request := &protobuf.IssueRequest{
		Ids:  make([]*protobuf.IssueOrder, len(ids)),
		User: user.ToProto(),
	}
	for i, id := range ids {
		request.Ids[i] = &protobuf.IssueOrder{
			Id:         id.Id,
			EbookInvId: id.EbookInvId,
		}
	}
	_, err := c.client.Issue(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
func (c *OrderService) Archive(ctx context.Context, ids []int64, user *model.User) error {
	request := &protobuf.ArchiveRequest{
		Ids:  ids,
		User: user.ToProto(),
	}
	_, err := c.client.Archive(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
func (c *OrderService) SendToPf(ctx context.Context, ids []int64, user *model.User) error {
	request := &protobuf.SendToPfRequest{
		Ids:  ids,
		User: user.ToProto(),
	}
	_, err := c.client.SendToPf(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
func (c *OrderService) ReturnToStorage(ctx context.Context, ids []int64, user *model.User) error {
	request := &protobuf.ReturnToStorageRequest{
		Ids:  ids,
		User: user.ToProto(),
	}
	_, err := c.client.ReturnToStorage(ctx, request)
	if err != nil {
		return err
	}
	return nil
}

func (c *OrderService) StateCounts(ctx context.Context, filters *model.StateCountFilters) ([]*model.StateCount, error) {
	request := &protobuf.StateCountsRequest{
		Filters: filters.ToProto(),
	}
	response, err := c.client.StateCounts(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		stateCounts := make([]*model.StateCount, len(response.GetStateCounts()))
		for i, s := range response.GetStateCounts() {
			stateCounts[i] = model.NewStateCountProto(s)
		}
		return stateCounts, nil
	default:
		return nil, err
	}
}
