package ebook

import (
	"github.com/web-rabis/ebook-client/model/dictionary"
	"github.com/web-rabis/ebook-client/protobuf"
)

type EbookBrief struct {
	Id                 int64                          `json:"id"`
	BibliographicLevel *dictionary.BibliographicLevel `json:"bibliographicLevel"`
	TypeDescription    *dictionary.TypeDescription    `json:"typeDescription"`
	Krv                bool                           `json:"krv"`
	Catalog            *dictionary.Catalog            `json:"catalog"`
	Author             string                         `json:"author"`
	Title              string                         `json:"title"`
	Sources            []*Source                      `json:"sources"`
	ServiceNotes       []*ServiceNote                 `json:"serviceNotes"`
	AuthorMark         *AuthorMark                    `json:"authorMark"`
	RCipher            string                         `json:"rCipher"`
}

func NewEbookBriefProto(e *protobuf.EbookBrief) *EbookBrief {
	if e == nil {
		return nil
	}
	return &EbookBrief{
		Id:                 e.Id,
		Author:             e.Author,
		Title:              e.Title,
		BibliographicLevel: dictionary.NewBibliographicLevelFromProto(e.BibliographicLevel),
		TypeDescription:    dictionary.NewTypeDescriptionFromProto(e.TypeDescription),
		Krv:                e.Krv,
		Catalog:            dictionary.NewCatalogFromProto(e.Catalog),
		Sources:            NewSourcesFromProto(e.Sources),
		ServiceNotes:       NewServiceNotesFromProto(e.ServiceNotes),
		AuthorMark:         NewAuthorMarkFromProto(e.AuthorMark),
		RCipher:            e.RCipher,
	}
}
