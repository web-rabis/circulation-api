package model

import "github.com/web-rabis/ebook-client/protobuf"

type EbookServiceNote struct {
	Id          int64                  `json:"id"`
	EbookId     int64                  `json:"ebookId"`
	ServiceData *DictionaryServiceData `json:"serviceData"`
}

func NewEbookServiceNoteFromProto(s *protobuf.EbookServiceNote) *EbookServiceNote {
	if s == nil {
		return nil
	}
	return &EbookServiceNote{
		Id:          s.Id,
		EbookId:     s.EbookId,
		ServiceData: NewDictionaryServiceDataFromProto(s.ServiceData),
	}
}
func NewEbookServiceNotesFromProto(s []*protobuf.EbookServiceNote) []*EbookServiceNote {
	if s == nil {
		return nil
	}
	var result []*EbookServiceNote
	for _, v := range s {
		result = append(result, NewEbookServiceNoteFromProto(v))
	}
	return result
}
