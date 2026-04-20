package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/web-rabis/order-client/client"
	"github.com/web-rabis/order-client/model"
	"github.com/web-rabis/order-client/protobuf"
)

type DictionaryService struct {
	client protobuf.DictionaryServiceClient
}

var _ client.DictionaryService = &DictionaryService{}

func NewDictionaryServiceClient(client protobuf.DictionaryServiceClient) client.DictionaryService {
	return &DictionaryService{
		client: client,
	}
}

func (c *DictionaryService) ReasonRejectionList(ctx context.Context, paging *model.Paging, filters *model.ReasonRejectionFilters) (int64, []*model.ReasonRejection, error) {
	request := &protobuf.ReasonRejectionListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	response, err := c.client.ReasonRejectionList(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		reasonRejections := make([]*model.ReasonRejection, len(response.GetReasonRejections()))
		for i, rr := range response.GetReasonRejections() {
			reasonRejections[i] = model.NewReasonRejectionProto(rr)
		}
		return response.Count, reasonRejections, nil
	default:
		return 0, nil, err
	}
}
func (c *DictionaryService) ReasonRejectionById(ctx context.Context, id int64) (*model.ReasonRejection, error) {
	request := &protobuf.ByIdRequest{
		Id: id,
	}
	response, err := c.client.ReasonRejectionById(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return model.NewReasonRejectionProto(response), nil
	default:
		return nil, err
	}
}

func (c *DictionaryService) DepartmentList(ctx context.Context, paging *model.Paging, filters *model.DepartmentFilters) (int64, []*model.Department, error) {
	request := &protobuf.DepartmentListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	response, err := c.client.DepartmentList(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		reasonRejections := make([]*model.Department, len(response.GetDepartments()))
		for i, rr := range response.GetDepartments() {
			reasonRejections[i] = model.NewDepartmentFromProto(rr)
		}
		return response.Count, reasonRejections, nil
	default:
		return 0, nil, err
	}
}
func (c *DictionaryService) DepartmentById(ctx context.Context, id int64) (*model.Department, error) {
	request := &protobuf.ByIdRequest{
		Id: id,
	}
	response, err := c.client.DepartmentById(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return model.NewDepartmentFromProto(response), nil
	default:
		return nil, err
	}
}
