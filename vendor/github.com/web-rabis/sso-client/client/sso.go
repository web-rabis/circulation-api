package client

import (
	"context"

	"github.com/web-rabis/sso-client/model"
)

type Base interface {
	Connect() error
	Close() error
	User() UserService
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name UserService
type UserService interface {
	UserById(ctx context.Context, id int64) (*model.User, error)
}
