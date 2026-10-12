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

func (s *ReaderService) ReaderByTicketNumber(ctx context.Context, ticketNumber int64) (*model.Reader, error) {
	response, err := s.client.ReaderByTicketNumber(ctx, &protobuf.ByTicketNumberRequest{TicketNumber: ticketNumber})
	switch status.Code(err) {
	case codes.OK:
		return model.ReaderFromProto(response), nil
	default:
		return nil, err
	}
}
func (s *ReaderService) ReaderList(ctx context.Context, filters *model.ReaderFilter, paging *model.Paging) (int64, []*model.Reader, error) {
	request := &protobuf.ReaderListRequest{
		Filters: filters.ToProto(),
		Paging:  paging.ToProto(),
	}
	response, err := s.client.ReaderList(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return response.Count, model.ReaderListFromProto(response.Result), nil
	default:
		return 0, nil, err
	}
}
func (s *ReaderService) ReaderRegistration(ctx context.Context, readerRegistrationRequest *model.ReaderRegistrationRequest) (int64, string, error) {
	request := readerRegistrationRequest.ToProto()
	response, err := s.client.ReaderRegistration(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return response.TicketNumber, response.Status, nil
	default:
		return 0, "", err
	}
}
func (s *ReaderService) ReaderCreate(ctx context.Context, reader *model.Reader, photo []byte) (int64, error) {
	request := &protobuf.ReaderCreateRequest{
		Reader: reader.ToProto(),
		Photo:  photo,
	}
	response, err := s.client.ReaderCreate(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return response.TicketNumber, nil
	default:
		return 0, err
	}
}
func (s *ReaderService) ReaderUpdate(ctx context.Context, reader *model.Reader, photo []byte) error {
	request := &protobuf.ReaderUpdateRequest{
		Reader: reader.ToProto(),
		Photo:  photo,
	}
	_, err := s.client.ReaderUpdate(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return nil
	default:
		return err
	}
}
func (s *ReaderService) ReaderDelete(ctx context.Context, ticketNumber int64) error {
	request := &protobuf.ReaderDeleteRequest{
		TicketNumber: ticketNumber,
	}
	_, err := s.client.ReaderDelete(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return nil
	default:
		return err
	}
}
func (s *ReaderService) ReaderApprove(ctx context.Context, ticketNumber int64) error {
	request := &protobuf.ReaderApproveRequest{
		TicketNumber: ticketNumber,
	}
	_, err := s.client.ReaderApprove(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return nil
	default:
		return err
	}
}
func (s *ReaderService) ReaderReject(ctx context.Context, ticketNumber int64) error {
	request := &protobuf.ReaderRejectRequest{
		TicketNumber: ticketNumber,
	}
	_, err := s.client.ReaderReject(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return nil
	default:
		return err
	}
}
func (s *ReaderService) ReaderPhotoByTicketNumber(ctx context.Context, ticketNumber int64) ([]byte, error) {
	request := &protobuf.ByTicketNumberRequest{
		TicketNumber: ticketNumber,
	}
	response, err := s.client.ReaderPhotoByTicketNumber(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return response.Photo, nil
	default:
		return nil, err
	}
}
func (s *ReaderService) ReaderUploadPhoto(ctx context.Context, ticketNumber int64, photo []byte, user *model.User) error {
	request := &protobuf.ReaderUploadPhotoRequest{
		TicketNumber: ticketNumber,
		Photo:        photo,
		User:         user.ToProto(),
	}
	_, err := s.client.ReaderUploadPhoto(ctx, request)
	switch status.Code(err) {
	case codes.OK:
		return nil
	default:
		return err
	}
}
