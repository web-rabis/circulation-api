package client

import (
	"context"

	"github.com/web-rabis/ebook-client/model"
	"github.com/web-rabis/ebook-client/model/ebook"
)

type Base interface {
	Connect() error
	Close() error
	EbookSvc() EbookService
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name EbookService
type EbookService interface {
	EbookBriefById(ctx context.Context, id int64) (*ebook.EbookBrief, error)
	EbookCardById(ctx context.Context, id int64) (*ebook.EbookCard, error)
	InvList(ctx context.Context, filters *model.InvFilters, paging *model.Paging) (int64, []*ebook.Inv, error)
}
