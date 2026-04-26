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

func (c *UserService) UserById(ctx context.Context, id int64) (*model.User, error) {
	request := &protobuf.ByIdRequest{
		Id: id,
	}
	response, err := c.client.UserById(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return model.NewUserFromProto(response), nil
	default:
		return nil, err
	}
}
