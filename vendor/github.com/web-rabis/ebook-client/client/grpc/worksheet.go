package grpc

import (
	"context"

	"github.com/web-rabis/ebook-client/client"
	"github.com/web-rabis/ebook-client/model"
	"github.com/web-rabis/ebook-client/model/worksheet"
	"github.com/web-rabis/ebook-client/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorksheetService struct {
	client protobuf.WorksheetSvcClient
}

var _ client.WorksheetService = &WorksheetService{}

func NewWorksheetServiceClient(client protobuf.WorksheetSvcClient) client.WorksheetService {
	return &WorksheetService{
		client: client,
	}
}

func (s *WorksheetService) UserWorksheet(ctx context.Context, userId, bLevelId, typeDescId int64) (*worksheet.UserWorksheet, error) {
	resp, err := s.client.UserWorksheet(ctx, &protobuf.UserWorksheetRequest{
		UserId:     userId,
		BLevelId:   bLevelId,
		TypeDescId: typeDescId,
	})
	switch status.Code(err) {
	case codes.OK:
		return worksheet.NewUserWorksheetFromProto(resp), nil
	default:
		return nil, err
	}
}

func (s *WorksheetService) UserWorksheets(ctx context.Context, userId int64) ([]*worksheet.UserWorksheet, error) {
	resp, err := s.client.UserWorksheets(ctx, &protobuf.UserWorksheetsRequest{UserId: userId})
	switch status.Code(err) {
	case codes.OK:
		return model.NewListFromProto(resp.Result, worksheet.NewUserWorksheetFromProto), nil
	default:
		return nil, err
	}
}

func (s *WorksheetService) UserWorksheetSave(ctx context.Context, userId int64, uw *worksheet.UserWorksheet) (*worksheet.UserWorksheet, error) {
	resp, err := s.client.UserWorksheetSave(ctx, &protobuf.UserWorksheetSaveRequest{
		UserId:    userId,
		Worksheet: uw.ToProto(),
	})
	switch status.Code(err) {
	case codes.OK:
		return worksheet.NewUserWorksheetFromProto(resp), nil
	default:
		return nil, err
	}
}
