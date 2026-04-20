package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/web-rabis/reader-client/client"
	"github.com/web-rabis/reader-client/model"
	"github.com/web-rabis/reader-client/protobuf"
)

type ReaderService struct {
	client protobuf.ReaderSvcClient
}

var _ client.ReaderService = &ReaderService{}

func NewReaderServiceClient(client protobuf.ReaderSvcClient) client.ReaderService {
	return &ReaderService{
		client: client,
	}
}

func (s *ReaderService) ReaderById(ctx context.Context, id int64) (*model.Reader, error) {
	response, err := s.client.ReaderById(ctx, &protobuf.EntityByIdRequest{Id: id})
	switch status.Code(err) {
	case codes.OK:
		return model.NewReaderFromProto(response), nil
	default:
		return nil, err
	}
}
