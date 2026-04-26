package model

import "github.com/web-rabis/ebook-client/protobuf"

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
type DictionaryServiceData struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewBibliographicLevelFromProto(b *protobuf.BibliographicLevel) *BibliographicLevel {
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
func NewTypeDescriptionFromProto(t *protobuf.TypeDescription) *TypeDescription {
	if t == nil {
		return nil
	}
	return &TypeDescription{
		Id:   t.Id,
		Code: t.Code,
		Name: t.Name,
	}
}
func NewCatalogFromProto(c *protobuf.Catalog) *Catalog {
	if c == nil {
		return nil
	}
	return &Catalog{
		Id:   c.Id,
		Code: c.Code,
		Name: c.Name,
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
func NewDictionaryServiceDataFromProto(s *protobuf.DictionaryServiceData) *DictionaryServiceData {
	if s == nil {
		return nil
	}
	return &DictionaryServiceData{
		Id:   s.Id,
		Code: s.Code,
		Name: s.Name,
	}
}
