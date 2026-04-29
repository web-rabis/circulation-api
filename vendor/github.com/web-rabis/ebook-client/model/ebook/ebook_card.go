package ebook

import (
	"github.com/web-rabis/ebook-client/protobuf"
)

type EbookCard struct {
	VolumeNumber int64       `json:"volumeNumber"`
	Author       string      `json:"author"`
	Title        string      `json:"title"`
	RAvt         string      `json:"rAvt"`
	Main         string      `json:"main"`
	RCipher      string      `json:"rCipher"`
	AuthorMark   *AuthorMark `json:"authorMark"`
	Language     *Language   `json:"language"`
	Indexes      []string    `json:"indexes"`
	Inv          []*Inv      `json:"inv"`
}

func NewEbookCardFromProto(f *protobuf.EbookCard) *EbookCard {
	if f == nil {
		return nil
	}
	return &EbookCard{
		VolumeNumber: f.VolumeNumber,
		Author:       f.Author,
		Title:        f.Title,
		RAvt:         f.RAvt,
		Main:         f.Main,
		RCipher:      f.RCipher,
		AuthorMark:   NewAuthorMarkFromProto(f.GetAuthorMark()),
		Language:     NewLanguageFromProto(f.Language),
		Indexes:      f.Indexes,
		Inv:          NewInvListFromProto(f.Inv),
	}
}
