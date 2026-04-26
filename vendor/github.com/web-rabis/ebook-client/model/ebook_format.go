package model

import "github.com/web-rabis/ebook-client/protobuf"

type EbookFormat struct {
	Format string `json:"format"`
}

func NewEbookFormatFromProto(f *protobuf.EbookFormat) *EbookFormat {
	if f == nil {
		return nil
	}
	return &EbookFormat{Format: f.Format}
}
