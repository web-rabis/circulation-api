package ebook

import (
	"github.com/web-rabis/ebook-client/model/dictionary"
	"github.com/web-rabis/ebook-client/protobuf"
)

type ServiceNote struct {
	Id          int64                   `json:"id"`
	EbookId     int64                   `json:"ebookId"`
	ServiceData *dictionary.ServiceData `json:"serviceData"`
}

func NewServiceNoteFromProto(s *protobuf.ServiceNote) *ServiceNote {
	if s == nil {
		return nil
	}
	return &ServiceNote{
		Id:          s.Id,
		EbookId:     s.EbookId,
		ServiceData: dictionary.NewServiceDataFromProto(s.ServiceData),
	}
}
func NewServiceNotesFromProto(s []*protobuf.ServiceNote) []*ServiceNote {
	if s == nil {
		return nil
	}
	var result []*ServiceNote
	for _, v := range s {
		result = append(result, NewServiceNoteFromProto(v))
	}
	return result
}
