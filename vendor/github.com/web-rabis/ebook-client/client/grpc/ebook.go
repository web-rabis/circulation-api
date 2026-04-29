package grpc

import (
	"context"

	"github.com/web-rabis/ebook-client/client"
	"github.com/web-rabis/ebook-client/model"
	"github.com/web-rabis/ebook-client/model/ebook"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/web-rabis/ebook-client/protobuf"
)

type EbookService struct {
	client protobuf.EbookSvcClient
}

var _ client.EbookService = &EbookService{}

func NewEbookServiceClient(client protobuf.EbookSvcClient) client.EbookService {
	return &EbookService{
		client: client,
	}
}
func (s *EbookService) EbookBriefById(ctx context.Context, id int64) (*ebook.EbookBrief, error) {
	e, err := s.client.EbookBriefById(ctx, &protobuf.EntityByIdRequest{Id: id})
	switch status.Code(err) {
	case codes.OK:
		return ebook.NewEbookBriefProto(e), nil
	default:
		return nil, err
	}
}
func (s *EbookService) EbookCardById(ctx context.Context, id int64) (*ebook.EbookCard, error) {
	response, err := s.client.EbookCardById(ctx, &protobuf.EntityByIdRequest{Id: id})
	switch status.Code(err) {
	case codes.OK:
		return ebook.NewEbookCardFromProto(response), nil
	}
	return nil, err
}
func (s *EbookService) InvList(ctx context.Context, filters *model.InvFilters, paging *model.Paging) (int64, []*ebook.Inv, error) {
	request := &protobuf.InvListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	response, err := s.client.InvList(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return response.Count, ebook.NewInvListFromProto(response.Result), nil
	}
	return 0, nil, err
}
