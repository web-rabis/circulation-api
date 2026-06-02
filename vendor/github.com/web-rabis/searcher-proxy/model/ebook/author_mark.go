package ebook

import "github.com/web-rabis/searcher-proxy/protobuf"

type AuthorMark struct {
	AuthorMark string `json:"authorMark"`
}

func NewAuthorMarkFromProto(a *protobuf.AuthorMark) *AuthorMark {
	if a == nil {
		return nil
	}
	return &AuthorMark{AuthorMark: a.AuthorMark}
}
