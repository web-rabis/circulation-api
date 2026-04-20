package client

import (
	"context"

	"github.com/web-rabis/reader-client/model"
)

type Base interface {
	Connect() error
	Close() error
	ReaderSvc() ReaderService
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name ReaderService
type ReaderService interface {
	ReaderById(ctx context.Context, ticketNumber int64) (*model.Reader, error)
}
