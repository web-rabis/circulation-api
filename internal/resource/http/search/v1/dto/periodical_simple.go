package dto

import (
	periodicalModel "github.com/web-rabis/searcher-proxy/model/periodical"
)

type PeriodicalSimpleSearchRequest struct {
	Query string `json:"query"`
}
type PeriodicalSimpleSearchResponse struct {
	Result []*periodicalModel.Periodical `json:"result"`
	Count  int64                         `json:"count"`
}
