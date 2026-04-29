package model

import (
	"github.com/web-rabis/ebook-client/protobuf"
)

type InvFilters struct {
	EbookId int64
}

func (f *InvFilters) ToProto() *protobuf.InvFilters {
	if f == nil {
		return nil
	}
	return &protobuf.InvFilters{
		EbookId: f.EbookId,
	}
}
