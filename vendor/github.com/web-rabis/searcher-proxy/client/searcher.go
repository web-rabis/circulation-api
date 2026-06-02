package client

import (
	"context"

	"github.com/web-rabis/searcher-proxy/model"
	"github.com/web-rabis/searcher-proxy/model/dictionary"
	"github.com/web-rabis/searcher-proxy/model/ebook"
	"github.com/web-rabis/searcher-proxy/model/periodical"
)

type Base interface {
	Connect() error
	Close() error
	EbookSearchSvc() EbookSearchService
	PeriodicalSearchSvc() PeriodicalSearchService
	DictionarySvc() DictionaryService
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name DictionaryService
type DictionaryService interface {
	CatalogList(ctx context.Context) (int64, []*dictionary.Catalog, error)
	TypeDescriptionList(ctx context.Context) (int64, []*dictionary.TypeDescription, error)
	BibliographicLevelList(ctx context.Context) (int64, []*dictionary.BibliographicLevel, error)
	LanguageList(ctx context.Context) (int64, []*dictionary.Language, error)
	SearchFieldList(ctx context.Context) (int64, []*dictionary.SearchField, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name EbookSearchService
type EbookSearchService interface {
	Search(ctx context.Context, filters *ebook.EbookSearchFilters, paging *model.Paging) (int64, []*ebook.Ebook, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name PeriodicalSearchService
type PeriodicalSearchService interface {
	Search(ctx context.Context, filters *periodical.PeriodicalSearchFilters, paging *model.Paging) (int64, []*periodical.Periodical, error)
	PeriodicalNumbers(ctx context.Context, id int64) (int64, []*periodical.PeriodicalNumber, error)
}
