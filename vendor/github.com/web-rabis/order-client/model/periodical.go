package model

import "github.com/web-rabis/order-client/protobuf"

type Periodical struct {
	Id                 int64  `json:"id"`
	Nkr                int64  `json:"nkr"`
	Type               string `json:"type"`
	Title              string `json:"title"`
	Number             string `json:"number"`
	Index              string `json:"index"`
	DataResponsibility string `json:"dataResponsibility"`
	Language           string `json:"language"`
	YearCount          int64  `json:"yearCount"`
	Industry           string `json:"industry"`
	PlaceEdition       string `json:"placeEdition"`
	Publishing         string `json:"publishing"`
	YearEdition        string `json:"yearEdition"`
	TitleInformation   string `json:"titleInformation"`
}

func NewPeriodicalFromProto(p *protobuf.Periodical) *Periodical {
	if p == nil {
		return nil
	}
	return &Periodical{
		Id:                 p.Id,
		Nkr:                p.Nkr,
		Type:               p.Type,
		Title:              p.Title,
		Number:             p.Number,
		Index:              p.Index,
		DataResponsibility: p.DataResponsibility,
		Language:           p.Language,
		YearCount:          p.YearCount,
		Industry:           p.Industry,
		PlaceEdition:       p.PlaceEdition,
		Publishing:         p.Publishing,
		YearEdition:        p.YearEdition,
		TitleInformation:   p.TitleInformation,
	}
}
