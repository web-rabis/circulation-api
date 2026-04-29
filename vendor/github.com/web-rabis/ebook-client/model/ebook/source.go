package ebook

import "github.com/web-rabis/ebook-client/protobuf"

type Source struct {
	Id                   int64  `json:"id"`
	EbookId              int64  `json:"ebookId"`
	Title                string `json:"title"`
	Data                 string `json:"data"`
	Responsibility       string `json:"responsibility"`
	Authors              string `json:"authors"`
	DataEdition          string `json:"dataEdition"`
	DataOut              string `json:"dataOut"`
	SpecialCharacter     string `json:"specialCharacter"`
	CompositeInformation string `json:"compositeInformation"`
	Series               string `json:"series"`
	Isbn                 string `json:"isbn"`
	Issn                 string `json:"issn"`
	ParallelName         string `json:"parallelName"`
	Year                 string `json:"year"`
	Pages                string `json:"pages"`
}

func NewSourceFromProto(s *protobuf.Source) *Source {
	if s == nil {
		return nil
	}
	return &Source{
		Id:                   s.Id,
		EbookId:              s.EbookId,
		Title:                s.Title,
		Data:                 s.Data,
		Responsibility:       s.Responsibility,
		Authors:              s.Authors,
		DataEdition:          s.DataEdition,
		DataOut:              s.DataOut,
		SpecialCharacter:     s.SpecialCharacter,
		CompositeInformation: s.CompositeInformation,
		Series:               s.Series,
		Isbn:                 s.Isbn,
		Issn:                 s.Issn,
		ParallelName:         s.ParallelName,
		Year:                 s.Year,
		Pages:                s.Pages,
	}
}
func NewSourcesFromProto(s []*protobuf.Source) []*Source {
	if s == nil {
		return nil
	}
	var result []*Source
	for _, v := range s {
		result = append(result, NewSourceFromProto(v))
	}
	return result
}
