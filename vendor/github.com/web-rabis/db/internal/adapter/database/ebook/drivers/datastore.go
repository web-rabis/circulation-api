package drivers

import (
	"context"

	"github.com/web-rabis/db/ebook"
	"github.com/web-rabis/db/model"
)

type DataStore interface {
	Base
}

type Base interface {
	Name() string
	Ping(ctx context.Context) error
	Close(ctx context.Context) error
	Connect(ctx context.Context) error

	Ebook() Ebook
	EbookInv() EbookInv
}
type Ebook interface {
	List(ctx context.Context, paging *model.Paging) ([]*ebook.Ebook, error)
	ById(ctx context.Context, id int64) (*ebook.Ebook, error)
	Create(ctx context.Context, ebook *ebook.Ebook) error
	Update(ctx context.Context, ebook *ebook.Ebook) error
	Delete(ctx context.Context, id int64) error
}
type EbookInv interface {
	List(ctx context.Context, paging *model.Paging) ([]*ebook.EbookInv, error)
	ById(ctx context.Context, id int64) (*ebook.EbookInv, error)
	Create(ctx context.Context, ebook *ebook.EbookInv) error
	Update(ctx context.Context, ebook *ebook.EbookInv) error
	Delete(ctx context.Context, id int64) error
}
