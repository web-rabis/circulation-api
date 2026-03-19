package model

import (
	"net/http"
	"strconv"
	"strings"
)

type Paging struct {
	Skip    int
	Limit   int
	SortKey string
	SortVal int
}

func (p *Paging) Sql() string {
	if p == nil {
		return ""
	}

	sql := ""

	if p.SortKey != "" {
		sql = sql + " order by " + p.SortKey

		if p.SortVal == -1 {
			sql = sql + " desc"
		} else {
			sql = sql + " asc"
		}
	}

	if p.Limit != 0 {
		sql = sql + "  limit " + strconv.Itoa(p.Limit)
	}

	sql = sql + " offset " + strconv.Itoa(p.Skip)

	return sql
}

func (p *Paging) Elk() map[string]string {
	var sortOrders = map[string]string{}
	var order = "ASC"
	if p.SortVal == -1 {
		order = "DESC"
	}
	// Список разрешённых полей для сортировки
	allowedFields := map[string]bool{
		"id":     true,
		"author": true,
		"title":  true,
	}
	for _, orderField := range strings.Split(p.SortKey, ",") {
		orderField = strings.ToLower(orderField)
		if _, ok := allowedFields[orderField]; ok {
			f := orderField
			if f != "id" {
				f += ".keyword"
			}
			sortOrders[f] = order
		}
	}
	return sortOrders
}

func PagingParseFromHttp(r *http.Request) (*Paging, error) {
	paging := Paging{}

	if limitParam := r.URL.Query().Get("limit"); limitParam != "" {
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			return nil, err
		}
		paging.Limit = limit
	}

	if pageParam := r.URL.Query().Get("skip"); pageParam != "" {
		page, err := strconv.Atoi(pageParam)
		if err != nil {
			return nil, err
		}
		skip := page - 1
		if paging.Limit != 0 {
			skip = paging.Limit * skip
		} else {
			skip = skip * 10
		}
		paging.Skip = skip
	}

	if orderParam := r.URL.Query().Get("order"); orderParam != "" {
		sortVal := 1
		if orderParam == "desc" {
			sortVal = -1
		}

		paging.SortVal = sortVal
	}

	if sortKeyFromQuery := r.URL.Query().Get("orderBy"); sortKeyFromQuery != "" {
		paging.SortKey = sortKeyFromQuery
	}

	return &paging, nil
}
