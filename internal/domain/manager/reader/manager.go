package reader

import (
	"context"

	readerClient "github.com/web-rabis/reader-client/client"
	"github.com/web-rabis/reader-client/model"
)

type IManager interface {
	ReaderByTicketNumber(ctx context.Context, ticketNumber int64) (*model.Reader, error)
}

type Manager struct {
	readerCli readerClient.ReaderService
}

func NewReaderManager(readerCli readerClient.ReaderService) *Manager {
	return &Manager{
		readerCli: readerCli,
	}
}

func (m *Manager) ReaderByTicketNumber(ctx context.Context, ticketNumber int64) (*model.Reader, error) {
	return m.readerCli.ReaderById(ctx, ticketNumber)
}
