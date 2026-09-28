package ebook

import "github.com/web-rabis/ebook-client/protobuf"

type Volume struct {
	Id                    int64  `json:"id"`
	EbookId               int64  `json:"ebookId"`
	Volume                string `json:"volume"`
	PhysicCharacter       string `json:"physicCharacter"`
	Size                  string `json:"size"`
	SupportingMaterial    string `json:"supportingMaterial"`
	TypeUnit              string `json:"typeUnit"`
	SizeUnit              string `json:"sizeUnit"`
	SpecificationMaterial string `json:"specificationMaterial"`
}

func NewVolumeFromProto(i *protobuf.Volume) *Volume {
	if i == nil {
		return nil
	}
	return &Volume{
		Id:                    i.Id,
		EbookId:               i.EbookId,
		Volume:                i.Volume,
		PhysicCharacter:       i.PhysicCharacter,
		Size:                  i.Size,
		SupportingMaterial:    i.SupportingMaterial,
		TypeUnit:              i.TypeUnit,
		SizeUnit:              i.SizeUnit,
		SpecificationMaterial: i.SpecificationMaterial,
	}
}
