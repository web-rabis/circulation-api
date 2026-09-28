package ebook_form

import (
	"time"

	"github.com/web-rabis/ebook-client/model/dictionary"
	"github.com/web-rabis/ebook-client/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EbookForm struct {
	Id                 int64                          `json:"id"`
	ParentId           *int64                         `json:"parentId"` // нет у записи верхнего уровня
	CreatedAt          time.Time                      `json:"createdAt"`
	UpdatedAt          time.Time                      `json:"updatedAt"`
	State              *dictionary.State              `json:"state"`
	BibliographicLevel *dictionary.BibliographicLevel `json:"bibliographicLevel"`
	TypeDescription    *dictionary.TypeDescription    `json:"typeDescription"`
	Krv                bool                           `json:"krv"`
	VolumeNumber       int64                          `json:"volumeNumber"`
	Catalog            *dictionary.Catalog            `json:"catalog"`
	Blocks             []*Block                       `json:"blocks"`
}

func NewEbookFormFromProto(e *protobuf.EbookForm) *EbookForm {
	if e == nil {
		return nil
	}
	return &EbookForm{
		Id:                 e.Id,
		ParentId:           e.ParentId,
		CreatedAt:          e.CreatedAt.AsTime(),
		UpdatedAt:          e.UpdatedAt.AsTime(),
		State:              dictionary.NewStateFromProto(e.State),
		BibliographicLevel: dictionary.NewBibliographicLevelFromProto(e.BibliographicLevel),
		TypeDescription:    dictionary.NewTypeDescriptionFromProto(e.TypeDescription),
		Krv:                e.Krv,
		VolumeNumber:       e.VolumeNumber,
		Catalog:            dictionary.NewCatalogFromProto(e.Catalog),
		Blocks:             NewBlocksFromProto(e.Blocks),
	}
}

// ToProto собирает EbookForm целиком для EbookFormSave: create/update
// определяется по Id (0 - новая запись) и, для каждого блока, по
// Block.Id/Block.EbookId.
func (e *EbookForm) ToProto() *protobuf.EbookForm {
	if e == nil {
		return nil
	}
	return &protobuf.EbookForm{
		Id:                 e.Id,
		ParentId:           e.ParentId,
		CreatedAt:          timestamppb.New(e.CreatedAt),
		UpdatedAt:          timestamppb.New(e.UpdatedAt),
		State:              e.State.ToProto(),
		BibliographicLevel: e.BibliographicLevel.ToProto(),
		TypeDescription:    e.TypeDescription.ToProto(),
		Krv:                e.Krv,
		VolumeNumber:       e.VolumeNumber,
		Catalog:            e.Catalog.ToProto(),
		Blocks:             BlocksToProto(e.Blocks),
	}
}
