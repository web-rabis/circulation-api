package grpc

import (
	"context"

	"github.com/web-rabis/ebook-client/client"
	"github.com/web-rabis/ebook-client/model"
	"github.com/web-rabis/ebook-client/model/dictionary"
	"github.com/web-rabis/ebook-client/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DictionaryService struct {
	client protobuf.DictionarySvcClient
}

var _ client.DictionaryService = &DictionaryService{}

func NewDictionaryServiceClient(client protobuf.DictionarySvcClient) client.DictionaryService {
	return &DictionaryService{
		client: client,
	}
}
func (s *DictionaryService) DictionaryList(ctx context.Context, filters *dictionary.DictionaryFilters, paging *model.Paging) (int64, []*dictionary.Dictionary, error) {
	req := &protobuf.DictionaryListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	e, err := s.client.DictionaryList(ctx, req)
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewDictionaryFromProto), nil
	default:
		return 0, nil, err
	}
}
func (s *DictionaryService) SearchDictionary(ctx context.Context, filters *dictionary.SearchDictionaryFilters, paging *model.Paging) (int64, []*dictionary.DictionaryBase, error) {
	req := &protobuf.SearchDictionaryRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	e, err := s.client.SearchDictionary(ctx, req)
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewDictionaryBaseFromProto), nil
	default:
		return 0, nil, err
	}
}

func (s *DictionaryService) CatalogList(ctx context.Context, filters *dictionary.CatalogFilters, paging *model.Paging) (int64, []*dictionary.Catalog, error) {
	req := &protobuf.CatalogListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	e, err := s.client.CatalogList(ctx, req)
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewCatalogFromProto), nil
	default:
		return 0, nil, err
	}
}
func (s *DictionaryService) TypeDescriptionList(ctx context.Context, filters *dictionary.TypeDescriptionFilters, paging *model.Paging) (int64, []*dictionary.TypeDescription, error) {
	req := &protobuf.TypeDescriptionListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	e, err := s.client.TypeDescriptionList(ctx, req)
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewTypeDescriptionFromProto), nil
	default:
		return 0, nil, err
	}
}
func (s *DictionaryService) BibliographicLevelList(ctx context.Context, filters *dictionary.BibliographicLevelFilters, paging *model.Paging) (int64, []*dictionary.BibliographicLevel, error) {
	req := &protobuf.BibliographicLevelListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	e, err := s.client.BibliographicLevelList(ctx, req)
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewBibliographicLevelFromProto), nil
	default:
		return 0, nil, err
	}
}
func (s *DictionaryService) LanguageList(ctx context.Context, filters *dictionary.LanguageFilters, paging *model.Paging) (int64, []*dictionary.Language, error) {
	req := &protobuf.DictionaryLanguageListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	e, err := s.client.LanguageList(ctx, req)
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewLanguageFromProto), nil
	default:
		return 0, nil, err
	}
}
func (s *DictionaryService) BlockList(ctx context.Context, filters *dictionary.BlockFilters, paging *model.Paging) (int64, []*dictionary.Block, error) {
	req := &protobuf.BlockListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	e, err := s.client.BlockList(ctx, req)
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewBlockFromProto), nil
	default:
		return 0, nil, err
	}
}
func (s *DictionaryService) BlockFieldList(ctx context.Context, filters *dictionary.BlockFieldFilters, paging *model.Paging) (int64, []*dictionary.BlockField, error) {
	req := &protobuf.BlockFieldListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	e, err := s.client.BlockFieldList(ctx, req)
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewBlockFieldFromProto), nil
	default:
		return 0, nil, err
	}
}
