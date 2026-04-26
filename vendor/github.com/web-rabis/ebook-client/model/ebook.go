package model

import "github.com/web-rabis/ebook-client/protobuf"

type Ebook struct {
	Id                 int64               `json:"id"`
	BibliographicLevel *BibliographicLevel `json:"bibliographicLevel"`
	TypeDescription    *TypeDescription    `json:"typeDescription"`
	Catalog            *Catalog            `json:"catalog"`
	Author             string              `json:"author"`
	Title              string              `json:"title"`
	Placement          *EbookPlacement     `json:"placement"`
	Format             *EbookFormat        `json:"format"`
	Sources            []*EbookSource      `json:"sources"`
	ServiceNotes       []*EbookServiceNote `json:"serviceNotes"`
	AuthorMark         *EbookAuthorMark    `json:"authorMark"`
	Krv                bool                `json:"krv"`
}

func NewEbookFromProto(e *protobuf.Ebook) *Ebook {
	if e == nil {
		return nil
	}
	return &Ebook{
		Id:                 e.Id,
		Author:             e.Author,
		Title:              e.Title,
		BibliographicLevel: NewBibliographicLevelFromProto(e.BibliographicLevel),
		TypeDescription:    NewTypeDescriptionFromProto(e.TypeDescription),
		Catalog:            NewCatalogFromProto(e.Catalog),
		Placement:          NewEbookPlacementFromProto(e.Placement),
		Format:             NewEbookFormatFromProto(e.Format),
		Sources:            NewEbookSourcesFromProto(e.Sources),
		ServiceNotes:       NewEbookServiceNotesFromProto(e.ServiceNotes),
		AuthorMark:         NewEbookAuthorMarkFromProto(e.AuthorMark),
		Krv:                e.Krv,
	}
}
