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
		return &model.Ebook{
			Id:                 id,
			BibliographicLevel: nil,
			TypeDescription:    nil,
			Catalog:            nil,
			Author:             e.Author,
			Title:              e.Title,
			Placement:          nil,
			Format:             nil,
		}, nil
	default:
		return nil, err
	}
}
func (s *EbookService) EbookInventory(ctx context.Context, id int64) ([]*model.EbookInv, error) {
	response, err := s.client.EbookInventory(ctx, &protobuf.EbookInventoryRequest{Id: id})
	switch status.Code(err) {
	case codes.OK:
		inventories := make([]*model.EbookInv, len(response.Inventories))
		for i, item := range response.Inventories {
			inventories[i] = model.NewEbookFromProto(item)
		}
		return inventories, nil
	default:
		return nil, err
	}
}
