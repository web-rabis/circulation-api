package client

import (
	"context"

	"github.com/web-rabis/order-client/model"
)

type Base interface {
	Connect() error
	Close() error
	Order() OrderService
	Dictionary() DictionaryService
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name OrderService
type OrderService interface {
	List(ctx context.Context, paging *model.Paging, filter *model.OrderFilters) (int64, []*model.Order, error)
	ById(ctx context.Context, id int64) (*model.Order, error)
	Reject(ctx context.Context, ids []int64, reasonRejectId int64, userId int64) error
	Redirect(ctx context.Context, ids []int64, departmentId int64, userId int64) error
	StateCounts(ctx context.Context, filters *model.StateCountFilters) ([]*model.StateCount, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name DictionaryService
type DictionaryService interface {
	ReasonRejectionList(ctx context.Context, paging *model.Paging, filter *model.ReasonRejectionFilters) (int64, []*model.ReasonRejection, error)
	ReasonRejectionById(ctx context.Context, id int64) (*model.ReasonRejection, error)
	DepartmentList(ctx context.Context, paging *model.Paging, filter *model.DepartmentFilters) (int64, []*model.Department, error)
	DepartmentById(ctx context.Context, id int64) (*model.Department, error)
}
