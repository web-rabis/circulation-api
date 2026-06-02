package ebook

import (
	"github.com/web-rabis/ebook-client/model/dictionary"
	"github.com/web-rabis/ebook-client/protobuf"
)

type Ebook struct {
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
	Inv                []*Inv                         `json:"inv"`
	Texts              []*Text                        `json:"texts"`
	Siglas             []*Sigla                       `json:"siglas"`
	Card               *EbookCard                     `json:"card"`
}

func NewEbookFromProto(e *protobuf.Ebook) *Ebook {
	if e == nil {
		return nil
	}
	return &Ebook{
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
		Inv:                NewInvListFromProto(e.Inv),
		Texts:              NewTextListFromProto(e.Texts),
		Siglas:             NewSiglaListFromProto(e.Siglas),
		Card:               NewEbookCardFromProto(e.Card),
	}
}
