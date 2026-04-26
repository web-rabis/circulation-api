package model

import "github.com/web-rabis/ebook-client/protobuf"

type EbookPlacement struct {
	Placement string `json:"placement"`
}

func NewEbookPlacementFromProto(p *protobuf.EbookPlacement) *EbookPlacement {
	if p == nil {
		return nil
	}
	return &EbookPlacement{Placement: p.Placement}
}
