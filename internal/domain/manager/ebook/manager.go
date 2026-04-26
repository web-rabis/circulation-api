package ebook

import (
	"context"

	ebookClient "github.com/web-rabis/ebook-client/client"
	ebookModel "github.com/web-rabis/ebook-client/model"
)

type IManager interface {
	EbookById(ctx context.Context, id int64) (*ebookModel.Ebook, error)
	EbookInventory(ctx context.Context, id int64) ([]*ebookModel.EbookInv, error)
}
type Manager struct {
	ebookCl ebookClient.EbookService
}

func NewManager(ebookCl ebookClient.EbookService) *Manager {
	return &Manager{
		ebookCl: ebookCl,
	}
}
func (m *Manager) EbookById(ctx context.Context, id int64) (*ebookModel.Ebook, error) {
	return m.ebookCl.EbookById(ctx, id)
}
func (m *Manager) EbookInventory(ctx context.Context, id int64) ([]*ebookModel.EbookInv, error) {
	return m.ebookCl.EbookInventory(ctx, id)
}
