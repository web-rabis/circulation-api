package grpc

import (
	"context"

	"github.com/web-rabis/searcher-proxy/client"
	"github.com/web-rabis/searcher-proxy/model"
	"github.com/web-rabis/searcher-proxy/model/dictionary"
	"github.com/web-rabis/searcher-proxy/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *DictionaryService) CatalogList(ctx context.Context) (int64, []*dictionary.Catalog, error) {
	e, err := s.client.CatalogList(ctx, &emptypb.Empty{})
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewCatalogFromProto), nil
	default:
		return 0, nil, err
	}
}
func (s *DictionaryService) TypeDescriptionList(ctx context.Context) (int64, []*dictionary.TypeDescription, error) {
	e, err := s.client.TypeDescriptionList(ctx, &emptypb.Empty{})
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewTypeDescriptionFromProto), nil
	default:
		return 0, nil, err
	}
}
func (s *DictionaryService) BibliographicLevelList(ctx context.Context) (int64, []*dictionary.BibliographicLevel, error) {
	e, err := s.client.BibliographicLevelList(ctx, &emptypb.Empty{})
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewBibliographicLevelFromProto), nil
	default:
		return 0, nil, err
	}
}
func (s *DictionaryService) LanguageList(ctx context.Context) (int64, []*dictionary.Language, error) {
	e, err := s.client.LanguageList(ctx, &emptypb.Empty{})
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewLanguageFromProto), nil
	default:
		return 0, nil, err
	}
}
func (s *DictionaryService) SearchFieldList(ctx context.Context) (int64, []*dictionary.SearchField, error) {
	e, err := s.client.SearchFieldList(ctx, &emptypb.Empty{})
	switch status.Code(err) {
	case codes.OK:
		return e.Count, model.NewListFromProto(e.Result, dictionary.NewSearchFieldFromProto), nil
	default:
		return 0, nil, err
	}
}
