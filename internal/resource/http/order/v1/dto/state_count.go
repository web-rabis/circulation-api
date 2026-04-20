package dto

import orderModel "github.com/web-rabis/order-client/model"

type StateCount struct {
	StateCode string `json:"stateCode"`
	Total     int64  `json:"total"`
}

func NewStateCounts(stateCounts []*orderModel.StateCount) []StateCount {
	var result = make([]StateCount, len(stateCounts))
	for i, stateCount := range stateCounts {
		result[i] = StateCount{
			StateCode: stateCount.State.Code,
			Total:     stateCount.Total,
		}
	}
	return result
}
