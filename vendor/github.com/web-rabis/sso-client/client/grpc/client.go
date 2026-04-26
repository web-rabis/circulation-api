package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/web-rabis/sso-client/client"
	"github.com/web-rabis/sso-client/model"
	"github.com/web-rabis/sso-client/protobuf"
)

type BaseClient struct {
	address  string
	conn     *grpc.ClientConn
	dialOpts []grpc.DialOption

	user client.UserService
}

var _ client.Base = &BaseClient{}

func NewClient(config *model.ConnectionConfig) (client.Base, error) {
	var grpcOpts []grpc.DialOption
	if config.Insecure == true {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	return &BaseClient{
		address:  config.Address,
		dialOpts: grpcOpts,
	}, nil
}

func (c *BaseClient) Connect() (err error) {
	c.conn, err = grpc.NewClient(c.address, c.dialOpts...)
	if err != nil {
		return err
	}

	return err
}

func (c *BaseClient) Close() error {
	return c.conn.Close()
}

func (c *BaseClient) User() client.UserService {
	if c.user == nil {
		c.user = NewUserServiceClient(protobuf.NewUserServiceClient(c.conn))
	}

	return c.user
}
