package client

import (
	"context"

	"github.com/web-rabis/ebook-client/model"
)

type Base interface {
	Connect() error
	Close() error
	EbookSvc() EbookService
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name EbookService
type EbookService interface {
	EbookById(ctx context.Context, id int64) (*model.Ebook, error)
	EbookInventory(ctx context.Context, id int64) ([]*model.EbookInv, error)
}
