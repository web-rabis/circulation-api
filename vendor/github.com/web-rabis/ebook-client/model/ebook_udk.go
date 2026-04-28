package model

import "github.com/web-rabis/ebook-client/protobuf"

type EbookUdk struct {
	Id      int64  `json:"id"`
	EbookId int64  `json:"ebookId"`
	Index   string `json:"index"`
}

func NewEbookUdkFromProto(s *protobuf.EbookUdk) *EbookUdk {
	if s == nil {
		return nil
	}
	return &EbookUdk{
		Id:      s.Id,
		EbookId: s.EbookId,
		Index:   s.Index,
	}
}
func NewEbookUdksFromProto(s []*protobuf.EbookUdk) []*EbookUdk {
	if s == nil {
		return nil
	}
	var result []*EbookUdk
	for _, v := range s {
		result = append(result, NewEbookUdkFromProto(v))
	}
	return result
}
