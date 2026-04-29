package model

import (
	"github.com/web-rabis/ebook-client/protobuf"
)

type Paging struct {
	Offset  int64  `json:"offset"`
	Limit   int64  `json:"limit"`
	SortKey string `json:"sort_key"`
	SortVal int32  `json:"sort_val"`
}

func (p *Paging) ToProto() *protobuf.Paging {
	if p == nil {
		return nil
	}
	return &protobuf.Paging{
		Offset:  p.Offset,
		Limit:   p.Limit,
		SortKey: p.SortKey,
		SortVal: p.SortVal,
	}
}
