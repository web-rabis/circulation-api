package model

import "github.com/web-rabis/ebook-client/protobuf"

type EbookAuthorMark struct {
	AuthorMark string `json:"format"`
}

func NewEbookAuthorMarkFromProto(a *protobuf.EbookAuthorMark) *EbookAuthorMark {
	if a == nil {
		return nil
	}
	return &EbookAuthorMark{AuthorMark: a.AuthorMark}
}
