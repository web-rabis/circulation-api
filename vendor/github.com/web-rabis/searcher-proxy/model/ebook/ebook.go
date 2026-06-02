package ebook

import (
	"github.com/web-rabis/searcher-proxy/model/dictionary"
	"github.com/web-rabis/searcher-proxy/protobuf"
)

type Ebook struct {
	Id                 int64                          `json:"id"`
	ParentId           *int64                         `json:"parentId"`
	VolumeNumber       int64                          `json:"volumeNumber"`
	BibliographicLevel *dictionary.BibliographicLevel `json:"bibliographicLevel"`
	TypeDescription    *dictionary.TypeDescription    `json:"typeDescription"`
	Krv                bool                           `json:"krv"`
	Digest             bool                           `json:"digest"`
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
	var parentId *int64
	if e.ParentId != 0 {
		parentId = &e.ParentId
	}
	return &Ebook{
		Id:                 e.Id,
		ParentId:           parentId,
		VolumeNumber:       e.VolumeNumber,
		Author:             e.Author,
		Title:              e.Title,
		BibliographicLevel: dictionary.NewBibliographicLevelFromProto(e.BibliographicLevel),
		TypeDescription:    dictionary.NewTypeDescriptionFromProto(e.TypeDescription),
		Krv:                e.Krv,
		Digest:             e.Digest,
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
