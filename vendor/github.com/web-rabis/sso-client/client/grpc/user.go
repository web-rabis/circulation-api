package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/web-rabis/sso-client/client"
	"github.com/web-rabis/sso-client/model"
	"github.com/web-rabis/sso-client/protobuf"
)

type UserService struct {
	client protobuf.UserServiceClient
}

var _ client.UserService = &UserService{}

func NewUserServiceClient(client protobuf.UserServiceClient) client.UserService {
	return &UserService{
		client: client,
	}
}

func (c *UserService) ById(ctx context.Context, id int64) (*model.User, error) {
	request := &protobuf.ByIdRequest{
		Id: id,
	}
	response, err := c.client.ById(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return model.NewUserFromProto(response), nil
	default:
		return nil, err
	}
}
func (c *UserService) List(ctx context.Context, filters *model.UserFilters, paging *model.Paging) (int64, []*model.User, error) {
	request := &protobuf.UserListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	response, err := c.client.List(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.NewUsersFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}
