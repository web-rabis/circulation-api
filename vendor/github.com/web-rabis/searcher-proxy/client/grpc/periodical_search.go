package grpc

import (
	"context"

	"github.com/web-rabis/searcher-proxy/client"
	"github.com/web-rabis/searcher-proxy/model"
	"github.com/web-rabis/searcher-proxy/model/periodical"
	"github.com/web-rabis/searcher-proxy/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PeriodicalSearchService struct {
	client protobuf.PeriodicalSearchSvcClient
}

var _ client.PeriodicalSearchService = &PeriodicalSearchService{}

func NewPeriodicalSearchServiceClient(client protobuf.PeriodicalSearchSvcClient) client.PeriodicalSearchService {
	return &PeriodicalSearchService{
		client: client,
	}
}

func (s *PeriodicalSearchService) Search(ctx context.Context, filters *periodical.PeriodicalSearchFilters, paging *model.Paging) (int64, []*periodical.Periodical, error) {
	req := &protobuf.PeriodicalSearchRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	res, err := s.client.Search(ctx, req)
	switch status.Code(err) {
	case codes.OK:
		return res.Count, model.NewListFromProto(res.Result, periodical.NewPeriodicalFromProto), nil
	default:
		return 0, nil, err
	}
}

func (s *PeriodicalSearchService) PeriodicalNumbers(ctx context.Context, periodicalId int64) (int64, []*periodical.PeriodicalNumber, error) {
	req := &protobuf.PeriodicalNumbersRequest{
		PeriodicalId: periodicalId,
	}
	res, err := s.client.PeriodicalNumbers(ctx, req)
	switch status.Code(err) {
	case codes.OK:
		return res.Count, model.NewListFromProto(res.Result, periodical.NewPeriodicalNumberFromProto), nil
	default:
		return 0, nil, err
	}
}
