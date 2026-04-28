package grpc

import (
	"context"

	"github.com/web-rabis/ebook-client/client"
	"github.com/web-rabis/ebook-client/model"
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
func (s *EbookService) EbookById(ctx context.Context, id int64) (*model.Ebook, error) {
	e, err := s.client.EbookById(ctx, &protobuf.EntityByIdRequest{Id: id})
	switch status.Code(err) {
	case codes.OK:
		return model.NewEbookFromProto(e), nil
	default:
		return nil, err
	}
}
func (s *EbookService) EbookInventory(ctx context.Context, id int64) ([]*model.EbookInv, error) {
	response, err := s.client.EbookInventory(ctx, &protobuf.EbookInventoryRequest{Id: id})
	switch status.Code(err) {
	case codes.OK:
		return model.NewEbookInvsFromProto(response.Inventories), nil
	default:
		return nil, err
	}
}
