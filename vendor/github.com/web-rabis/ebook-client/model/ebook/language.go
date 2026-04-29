package ebook

import (
	"github.com/web-rabis/ebook-client/model/dictionary"
	"github.com/web-rabis/ebook-client/protobuf"
)

type Language struct {
	Id            int64                `json:"id"`
	EbookId       int64                `json:"ebookId"`
	Language      *dictionary.Language `json:"language"`
	SourceRewrite string               `json:"sourceRewrite"`
}

func NewLanguageFromProto(e *protobuf.Language) *Language {
	if e == nil {
		return nil
	}
	return &Language{
		Id:            e.Id,
		EbookId:       e.EbookId,
		Language:      dictionary.NewLanguageFromProto(e.Language),
		SourceRewrite: e.SourceRewrite,
	}
}
