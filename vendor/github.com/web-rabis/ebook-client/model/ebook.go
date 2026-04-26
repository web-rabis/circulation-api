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

type Department struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Type string
}
type State struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewEbookFromProto(i *protobuf.EbookInv) *EbookInv {
	if i == nil {
		return nil
	}
	return &EbookInv{
		Id:         i.Id,
		EbookId:    i.EbookId,
		InvNumber:  i.InvNumber,
		Barcode:    i.Barcode,
		Department: NewDepartmentFromProto(i.Department),
		State:      NewStateFromProto(i.State),
	}
}
func NewDepartmentFromProto(d *protobuf.Department) *Department {
	if d == nil {
		return nil
	}
	return &Department{
		Id:   d.Id,
		Code: d.Code,
		Name: d.Name,
		Type: d.Type,
	}
}
func NewStateFromProto(s *protobuf.State) *State {
	if s == nil {
		return nil
	}
	return &State{
		Id:   s.Id,
		Code: s.Code,
		Name: s.Name,
	}
}
