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

func (s *EbookService) EbookList(ctx context.Context, filters *model.EbookFilters, paging *model.Paging) (int64, []*ebook.Ebook, error) {
	request := &protobuf.EbookListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	e, err := s.client.EbookList(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return e.Count, ebook.NewEbookListFromProto(e.Result), nil
	default:
		return 0, nil, err
	}
}

func (s *EbookService) EbookById(ctx context.Context, id int64, withCard bool) (*ebook.Ebook, error) {
	e, err := s.client.EbookById(ctx, &protobuf.EbookByIdRequest{Id: id, WithCard: withCard})
	switch status.Code(err) {
	case codes.OK:
		return ebook.NewEbookFromProto(e), nil
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

func (s *EbookService) EbookDelete(ctx context.Context, id, userId int64) (*ebook.Ebook, error) {
	e, err := s.client.EbookDelete(ctx, &protobuf.EbookDeleteRequest{Id: id, UserId: userId})
	switch status.Code(err) {
	case codes.OK:
		return ebook.NewEbookFromProto(e), nil
	default:
		return nil, err
	}
}
