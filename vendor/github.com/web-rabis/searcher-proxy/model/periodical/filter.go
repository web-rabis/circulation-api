package periodical

import (
	"github.com/web-rabis/searcher-proxy/protobuf"
)

type PeriodicalSearchFilters struct {
	Query          string `json:"query"`
	Type           string `json:"type"`
	IsSimpleSearch bool   `json:"isSimpleSearch"`
}

func (req *PeriodicalSearchFilters) ToProto() *protobuf.PeriodicalSearchFilters {
	return &protobuf.PeriodicalSearchFilters{
		Query:          req.Query,
		Type:           req.Type,
		IsSimpleSearch: req.IsSimpleSearch,
	}
}
