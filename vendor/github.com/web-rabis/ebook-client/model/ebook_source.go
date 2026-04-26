package model

import "github.com/web-rabis/ebook-client/protobuf"

type EbookSource struct {
	Id      int64  `json:"id"`
	EbookId int64  `json:"ebookId"`
	Title   string `json:"title"`
}

func NewEbookSourceFromProto(s *protobuf.EbookSource) *EbookSource {
	if s == nil {
		return nil
	}
	return &EbookSource{
		Id:      s.Id,
		EbookId: s.EbookId,
		Title:   s.Title,
	}
}
func NewEbookSourcesFromProto(s []*protobuf.EbookSource) []*EbookSource {
	if s == nil {
		return nil
	}
	var result []*EbookSource
	for _, v := range s {
		result = append(result, NewEbookSourceFromProto(v))
	}
	return result
}
