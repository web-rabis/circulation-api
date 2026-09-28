package grpc

import (
	"context"

	"github.com/web-rabis/ebook-client/client"
	"github.com/web-rabis/ebook-client/model/ebook_form"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/web-rabis/ebook-client/protobuf"
)

type EbookFormService struct {
	client protobuf.EbookFormSvcClient
}

var _ client.EbookFormService = &EbookFormService{}

func NewEbookFormServiceClient(client protobuf.EbookFormSvcClient) client.EbookFormService {
	return &EbookFormService{
		client: client,
	}
}

func (s *EbookFormService) EbookFormById(ctx context.Context, id int64) (*ebook_form.EbookForm, error) {
	e, err := s.client.EbookFormById(ctx, &protobuf.EntityByIdRequest{Id: id})
	switch status.Code(err) {
	case codes.OK:
		return ebook_form.NewEbookFormFromProto(e), nil
	default:
		return nil, err
	}
}

func (s *EbookFormService) EbookFormSave(ctx context.Context, form *ebook_form.EbookForm, userId int64) (*ebook_form.EbookForm, error) {
	e, err := s.client.EbookFormSave(ctx, &protobuf.EbookFormSaveRequest{Form: form.ToProto(), UserId: userId})
	switch status.Code(err) {
	case codes.OK:
		return ebook_form.NewEbookFormFromProto(e), nil
	default:
		return nil, err
	}
}
