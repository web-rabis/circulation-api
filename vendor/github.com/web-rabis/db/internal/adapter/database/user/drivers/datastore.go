package drivers

import (
	"context"
	"github.com/web-rabis/db/model"
	"github.com/web-rabis/db/user"
)

type DataStore interface {
	Base
}

type Base interface {
	Name() string
	Ping(ctx context.Context) error
	Close(ctx context.Context) error
	Connect(ctx context.Context) error

	User() User
}
type User interface {
	List(ctx context.Context, paging *model.Paging) ([]*user.User, error)
	ByUsername(ctx context.Context, username string) (*user.User, error)
	ById(ctx context.Context, id int) (*user.User, error)
}
