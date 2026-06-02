package client

import (
	"context"

	"github.com/web-rabis/ebook-client/model"
	"github.com/web-rabis/ebook-client/model/dictionary"
	"github.com/web-rabis/ebook-client/model/ebook"
)

type Base interface {
	Connect() error
	Close() error
	EbookSvc() EbookService
	DictionarySvc() DictionaryService
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name EbookService
type EbookService interface {
	EbookById(ctx context.Context, id int64, withCard bool) (*ebook.Ebook, error)
	EbookBriefById(ctx context.Context, id int64) (*ebook.EbookBrief, error)
	EbookCardById(ctx context.Context, id int64) (*ebook.EbookCard, error)
	InvList(ctx context.Context, filters *model.InvFilters, paging *model.Paging) (int64, []*ebook.Inv, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name DictionaryService
type DictionaryService interface {
	DictionaryList(ctx context.Context, filters *dictionary.DictionaryFilters, paging *model.Paging) (int64, []*dictionary.Dictionary, error)
	SearchDictionary(ctx context.Context, filters *dictionary.SearchDictionaryFilters, paging *model.Paging) (int64, []*dictionary.DictionaryBase, error)
	CatalogList(ctx context.Context, filters *dictionary.CatalogFilters, paging *model.Paging) (int64, []*dictionary.Catalog, error)
	TypeDescriptionList(ctx context.Context, filters *dictionary.TypeDescriptionFilters, paging *model.Paging) (int64, []*dictionary.TypeDescription, error)
	BibliographicLevelList(ctx context.Context, filters *dictionary.BibliographicLevelFilters, paging *model.Paging) (int64, []*dictionary.BibliographicLevel, error)
	LanguageList(ctx context.Context, filters *dictionary.LanguageFilters, paging *model.Paging) (int64, []*dictionary.Language, error)
	BlockList(ctx context.Context, filters *dictionary.BlockFilters, paging *model.Paging) (int64, []*dictionary.Block, error)
	BlockFieldList(ctx context.Context, filters *dictionary.BlockFieldFilters, paging *model.Paging) (int64, []*dictionary.BlockField, error)
}
