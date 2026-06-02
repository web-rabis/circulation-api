package periodical

import "github.com/web-rabis/searcher-proxy/protobuf"

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
}

func NewPeriodicalFromProto(s *protobuf.Periodical) *Periodical {
	if s == nil {
		return nil
	}
	return &Periodical{
		Id:                 s.Id,
		Nkr:                s.Nkr,
		Type:               s.Type,
		Title:              s.Title,
		Number:             s.Number,
		Index:              s.Index,
		DataResponsibility: s.DataResponsibility,
		Language:           s.Language,
		YearCount:          s.YearCount,
		Industry:           s.Industry,
		PlaceEdition:       s.PlaceEdition,
		Publishing:         s.Publishing,
		YearEdition:        s.YearEdition,
	}
}
