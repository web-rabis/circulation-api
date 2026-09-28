package ebook

import (
	"github.com/web-rabis/ebook-client/model/dictionary"
	"github.com/web-rabis/ebook-client/protobuf"
)

type Publishing struct {
	Id              int64                  `json:"id,omitempty"`
	EbookId         int64                  `json:"ebookId,omitempty"`
	PlaceEdition    *dictionary.City       `json:"placeEdition,omitempty"`
	PublishingHouse *dictionary.Publishing `json:"publishingHouse,omitempty"`
	YearEdition     string                 `json:"yearEdition,omitempty"`
}

func NewEbookPublishingFromProto(e *protobuf.Publishing) *Publishing {
	if e == nil {
		return nil
	}
	return &Publishing{
		Id:              e.Id,
		EbookId:         e.EbookId,
		PlaceEdition:    dictionary.NewCityFromProto(e.PlaceEdition),
		PublishingHouse: dictionary.NewPublishingFromProto(e.Publishing),
		YearEdition:     e.YearEdition,
	}
}
func NewEbookPublishingsFromProto(s []*protobuf.Publishing) []*Publishing {
	if s == nil {
		return nil
	}
	var result []*Publishing
	for _, v := range s {
		result = append(result, NewEbookPublishingFromProto(v))
	}
	return result
}
