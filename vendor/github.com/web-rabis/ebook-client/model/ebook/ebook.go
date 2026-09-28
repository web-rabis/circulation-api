package ebook

import (
	"time"

	"github.com/web-rabis/ebook-client/model/dictionary"
	"github.com/web-rabis/ebook-client/protobuf"
)

const (
	EbookEditingStatus          = "Ebook.Editing"
	EbookWorkStatus             = "Ebook.Work"
	EbookDeletedStatus          = "Ebook.Deleted"
	EbookInStorage              = "Ebook.InStorage"
	EbookImported               = "Ebook.Imported"
	EbookImportedFromCompletion = "Ebook.ImportedFromCompletion"
)

type Ebook struct {
	Id                 int64                          `json:"id"`
	ParentId           int64                          `json:"parentId"`
	CreatedAt          time.Time                      `json:"createdAt"`
	UpdatedAt          *time.Time                     `json:"updatedAt"`
	State              *dictionary.State              `json:"state"`
	BibliographicLevel *dictionary.BibliographicLevel `json:"bibliographicLevel"`
	TypeDescription    *dictionary.TypeDescription    `json:"typeDescription"`
	Krv                bool                           `json:"krv"`
	VolumeNumber       int64                          `json:"volumeNumber"`
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
	Publishings        []*Publishing                  `json:"publishings"`
	Language           *Language                      `json:"language"`
	// RCipher - полочный шифр, вычисляется на сервере, в БД не хранится.
	RCipher string `json:"rCipher"`
}

func NewEbookFromProto(e *protobuf.Ebook) *Ebook {
	if e == nil {
		return nil
	}
	var updatedAt *time.Time
	if e.UpdatedAt != nil {
		t := e.UpdatedAt.AsTime()
		updatedAt = &t
	}
	return &Ebook{
		Id:                 e.Id,
		ParentId:           e.ParentId,
		CreatedAt:          e.CreatedAt.AsTime(),
		UpdatedAt:          updatedAt,
		State:              dictionary.NewStateFromProto(e.State),
		Author:             e.Author,
		Title:              e.Title,
		BibliographicLevel: dictionary.NewBibliographicLevelFromProto(e.BibliographicLevel),
		TypeDescription:    dictionary.NewTypeDescriptionFromProto(e.TypeDescription),
		Krv:                e.Krv,
		VolumeNumber:       e.VolumeNumber,
		Catalog:            dictionary.NewCatalogFromProto(e.Catalog),
		Sources:            NewSourcesFromProto(e.Sources),
		ServiceNotes:       NewServiceNotesFromProto(e.ServiceNotes),
		AuthorMark:         NewAuthorMarkFromProto(e.AuthorMark),
		Inv:                NewInvListFromProto(e.Inv),
		Texts:              NewTextListFromProto(e.Texts),
		Siglas:             NewSiglaListFromProto(e.Siglas),
		Card:               NewEbookCardFromProto(e.Card),
		Publishings:        NewEbookPublishingsFromProto(e.Publishings),
		Language:           NewLanguageFromProto(e.Language),
		RCipher:            e.RCipher,
	}
}
func NewEbookListFromProto(e []*protobuf.Ebook) []*Ebook {
	if e == nil {
		return nil
	}
	var result []*Ebook
	for _, v := range e {
		result = append(result, NewEbookFromProto(v))
	}
	return result
}
