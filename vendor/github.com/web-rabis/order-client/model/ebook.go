package model

import "github.com/web-rabis/order-client/protobuf"

type Ebook struct {
	Id                 int64               `json:"id"`
	BibliographicLevel *BibliographicLevel `json:"bibliographicLevel"`
	TypeDescription    *TypeDescription    `json:"typeDescription"`
	Catalog            *Catalog            `json:"catalog"`
	Author             string              `json:"author"`
	Title              string              `json:"title"`
	Placement          *EbookPlacement     `json:"placement"`
	Format             *EbookFormat        `json:"format"`
}

type EbookInv struct {
	Id         int64       `json:"id"`
	EbookId    int64       `json:"ebookId"`
	InvNumber  string      `json:"invNumber"`
	Barcode    string      `json:"barcode"`
	Department *Department `json:"department"`
	State      *State      `json:"state"`
}

type EbookPlacement struct {
	Placement int64 `json:"placement"`
}

type EbookFormat struct {
	Format string `json:"format"`
}

type BibliographicLevel struct {
	Id         int64  `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	TypeEbooks string `json:"typeEbooks"`
}

type TypeDescription struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type Catalog struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewEbookFromProto(e *protobuf.Ebook) *Ebook {
	if e == nil {
		return nil
	}
	return &Ebook{
		Id:                 e.Id,
		BibliographicLevel: NewBibliographicLevel(e.BibliographicLevel),
		TypeDescription:    NewTypeDescription(e.TypeDescription),
		Catalog:            NewCatalog(e.Catalog),
		Author:             e.Author,
		Title:              e.Title,
		Placement:          NewEbookPlacementFromProto(e.Placement),
		Format:             NewEbookFormatFromProto(e.Format),
	}
}
func NewEbookInvFromProto(e *protobuf.EbookInv) *EbookInv {
	if e == nil {
		return nil

	}
	return &EbookInv{
		Id:         e.Id,
		EbookId:    e.EbookId,
		InvNumber:  e.InvNumber,
		Barcode:    e.Barcode,
		Department: NewDepartmentFromProto(e.Department),
		State:      NewStateFromProto(e.State),
	}
}
func NewEbookPlacementFromProto(e *protobuf.EbookPlacement) *EbookPlacement {
	if e == nil {
		return nil
	}
	return &EbookPlacement{
		Placement: e.Placement,
	}
}
func NewEbookFormatFromProto(e *protobuf.EbookFormat) *EbookFormat {
	if e == nil {
		return nil
	}
	return &EbookFormat{
		Format: e.Format,
	}
}
func NewBibliographicLevel(b *protobuf.BibliographicLevel) *BibliographicLevel {
	if b == nil {
		return nil
	}
	return &BibliographicLevel{
		Id:         b.Id,
		Code:       b.Code,
		Name:       b.Name,
		TypeEbooks: b.TypeEbooks,
	}
}
func NewTypeDescription(t *protobuf.TypeDescription) *TypeDescription {
	if t == nil {
		return nil
	}
	return &TypeDescription{
		Id:   t.Id,
		Code: t.Code,
		Name: t.Name,
	}
}
func NewCatalog(c *protobuf.Catalog) *Catalog {
	if c == nil {
		return nil
	}
	return &Catalog{
		Id:   c.Id,
		Code: c.Code,
		Name: c.Name,
	}
}
