package reader

import (
	"context"
	
	"github.com/web-rabis/db/internal/adapter/database/user/drivers"
	"github.com/web-rabis/db/model"
	"github.com/web-rabis/db/user"
)

type IManager interface {
	List(ctx context.Context, paging *model.Paging) ([]*user.User, error)
	ById(ctx context.Context, id int) (*user.User, error)
	ByUsername(ctx context.Context, username string) (*user.User, error)
}
type Manager struct {
	ds drivers.DataStore
}

func NewManager(ds drivers.DataStore) *Manager {
	return &Manager{
		ds: ds,
	}
}

func (m *Manager) List(ctx context.Context, paging *model.Paging) ([]*user.User, error) {
	return m.ds.User().List(ctx, paging)
}
func (m *Manager) ById(ctx context.Context, id int) (*user.User, error) {
	return m.ds.User().ById(ctx, id)
}
func (m *Manager) ByUsername(ctx context.Context, username string) (*user.User, error) {
	return m.ds.User().ByUsername(ctx, username)
}
