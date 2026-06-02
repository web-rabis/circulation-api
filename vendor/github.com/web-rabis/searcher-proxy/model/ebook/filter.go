package ebook

import (
	"github.com/web-rabis/searcher-proxy/protobuf"
)

type EbookSearchFilters struct {
	Query               string  `json:"query"`
	Type                string  `json:"type"`
	IsSimpleSearch      bool    `json:"isSimpleSearch"`
	Catalogs            []int64 `json:"catalogs"`
	Languages           []int64 `json:"languages"`
	Fields              []int64 `json:"fields"`
	TypeDescriptions    []int64 `json:"typeDescriptions"`
	BibliographicLevels []int64 `json:"bibliographicLevels"`
	FieldsQuery         string  `json:"fieldsQuery"`
	YearEditionLow      int64   `json:"yearEditionLow"`
	YearEditionHigh     int64   `json:"yearEditionHigh"`
}

func (req *EbookSearchFilters) ToProto() *protobuf.EbookSearchFilters {
	return &protobuf.EbookSearchFilters{
		Query:               req.Query,
		Type:                req.Type,
		IsSimpleSearch:      req.IsSimpleSearch,
		Catalogs:            req.Catalogs,
		Languages:           req.Languages,
		Fields:              req.Fields,
		TypeDescriptions:    req.TypeDescriptions,
		BibliographicLevels: req.BibliographicLevels,
		FieldsQuery:         req.FieldsQuery,
		YearEditionLow:      req.YearEditionLow,
		YearEditionHigh:     req.YearEditionHigh,
	}
}
