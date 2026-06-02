package grpc

import (
	"context"

	"github.com/web-rabis/searcher-proxy/client"
	"github.com/web-rabis/searcher-proxy/model"
	"github.com/web-rabis/searcher-proxy/model/ebook"
	"github.com/web-rabis/searcher-proxy/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EbookSearchService struct {
	client protobuf.EbookSearchSvcClient
}

var _ client.EbookSearchService = &EbookSearchService{}

func NewEbookSearchServiceClient(client protobuf.EbookSearchSvcClient) client.EbookSearchService {
	return &EbookSearchService{
		client: client,
	}
}

func (s *EbookSearchService) Search(ctx context.Context, filters *ebook.EbookSearchFilters, paging *model.Paging) (int64, []*ebook.Ebook, error) {
	req := &protobuf.EbookSearchRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	res, err := s.client.Search(ctx, req)
	switch status.Code(err) {
	case codes.OK:
		return res.Count, model.NewListFromProto(res.Result, ebook.NewEbookFromProto), nil
	default:
		return 0, nil, err
	}
}
