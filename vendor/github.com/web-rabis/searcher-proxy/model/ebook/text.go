package ebook

import (
	"github.com/web-rabis/searcher-proxy/protobuf"
)

type Text struct {
	Id      int64  `json:"id"`
	EbookId int64  `json:"ebookId"`
	Text    string `json:"text"`
	Note    string `json:"note"`
	Url     string `json:"url"`
}

func NewTextFromProto(i *protobuf.Text) *Text {
	if i == nil {
		return nil
	}
	return &Text{
		Id:      i.Id,
		EbookId: i.EbookId,
		Text:    i.Text,
		Note:    i.Note,
		Url:     i.Url,
	}
}
func NewTextListFromProto(s []*protobuf.Text) []*Text {
	if s == nil {
		return nil
	}
	var result []*Text
	for _, v := range s {
		result = append(result, NewTextFromProto(v))
	}
	return result
}
