package eorder

import (
	"net/http"
	"strconv"
	"time"
)

type EOrderFilter struct {
	TicketNumber  int64
	Status        string
	OrderedAtLow  *time.Time
	OrderedAtHigh *time.Time
}

func (f *EOrderFilter) Sql() string {
	if f == nil {
		return ""
	}
	var sql = ""
	if f.Status != "" {
		sql = "state_id = '" + f.Status + "'"
	}
	if f.TicketNumber != 0 {
		if sql != "" {
			sql += " and "
		}
		sql += "reader_ticket_num = '" + strconv.FormatInt(int64(f.TicketNumber), 10) + "'"
	}
	if f.OrderedAtLow != nil || f.OrderedAtHigh != nil {
		if sql != "" {
			sql += " and "
		}
		if f.OrderedAtLow != nil && f.OrderedAtHigh != nil {
			sql = sql + "ordered_date between '" + f.OrderedAtLow.Format("2006-01-02") + "' and '" + f.OrderedAtHigh.Format("2006-01-02") + "'"
		} else if f.OrderedAtLow != nil {
			sql = sql + "ordered_date >= '" + f.OrderedAtLow.Format("2006-01-02") + "'"
		} else if f.OrderedAtHigh != nil {
			sql = sql + "ordered_date <= '" + f.OrderedAtHigh.Format("2006-01-02") + "'"
		}
	}
	return sql
}

func EOrderFilterParseFromHttp(r *http.Request) (*EOrderFilter, error) {
	filter := EOrderFilter{}

	if status := r.URL.Query().Get("status"); status != "" && status != "*" {
		filter.Status = status
	}
	if ticketNumber := r.URL.Query().Get("ticketNumber"); ticketNumber != "" {
		filter.TicketNumber, _ = strconv.ParseInt(ticketNumber, 10, 64)
	}

	if filter.Sql() == "" {
		return nil, nil
	}
	return &filter, nil
}
