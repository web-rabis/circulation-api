package dictionary

import (
	"context"

	"github.com/web-rabis/order-client/client"
	"github.com/web-rabis/order-client/model"
)

type IManager interface {
	ReasonRejectionList(ctx context.Context) (int64, []*model.ReasonRejection, error)
	DepartmentList(ctx context.Context) (int64, []*model.Department, error)
}
type Manager struct {
	dictionaryCli client.DictionaryService
}

func NewManager(dictionaryCli client.DictionaryService) *Manager {
	return &Manager{
		dictionaryCli: dictionaryCli,
	}
}
func (m *Manager) ReasonRejectionList(ctx context.Context) (int64, []*model.ReasonRejection, error) {
	return m.dictionaryCli.ReasonRejectionList(ctx, nil, nil)
}
func (m *Manager) DepartmentList(ctx context.Context) (int64, []*model.Department, error) {
	return m.dictionaryCli.DepartmentList(ctx, nil, nil)
}
