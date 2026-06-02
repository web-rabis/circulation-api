package ebook

import "github.com/web-rabis/searcher-proxy/protobuf"

type EbookCard struct {
	Author  string   `json:"author"`
	Title   string   `json:"title"`
	RAvt    string   `json:"rAvt"`
	Main    string   `json:"main"`
	RCipher string   `json:"rCipher"`
	Indexes []string `json:"indexes"`
}

func NewEbookCardFromProto(c *protobuf.EbookCard) *EbookCard {
	if c == nil {
		return nil
	}
	return &EbookCard{
		Author:  c.Author,
		Title:   c.Title,
		RAvt:    c.RAvt,
		Main:    c.Main,
		RCipher: c.RCipher,
		Indexes: c.Indexes,
	}
}
