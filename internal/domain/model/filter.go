package model

import "strconv"

type LibDepartmentFilter struct {
	Type string
}
type CatalogFilter struct {
	ReaderCatalog string
}
type BlockFilter struct {
}
type BlockFieldFilter struct {
	IsSearch *bool
}
type OrderFilter struct {
	EbookId           *int64
	PeriodicalId      *int64
	NotEmptyInvNumber bool
	InStateCodes      []string
	TicketNumber      *int
}

func (f *BlockFieldFilter) Sql() string {
	sql := ""
	if f.IsSearch != nil {
		if *f.IsSearch {
			sql = "search=1"
		} else {
			sql = "search=0"
		}
	}
	if sql != "" {
		sql = "where " + sql
	}
	return sql
}

func (f *BlockFilter) Sql() string {

	return ""
}
func (f *CatalogFilter) Sql() string {
	sql := ""
	if f.ReaderCatalog != "" {
		sql = "readercatalog='" + f.ReaderCatalog + "'"
	}
	if sql != "" {
		sql = "where " + sql
	}
	return sql
}

func (f *LibDepartmentFilter) Sql() string {
	sql := ""
	if f.Type != "" {
		sql = "type='" + f.Type + "'"
	}
	if sql != "" {
		sql = "where " + sql
	}
	return sql
}

func (f *OrderFilter) Sql() string {
	sql := ""
	if f.EbookId != nil {
		sql = sql + " and eo.ebook_id=" + strconv.FormatInt(*f.EbookId, 10)
	}
	if f.PeriodicalId != nil {
		sql = sql + " and eo.idgazjurnumbers=" + strconv.FormatInt(*f.PeriodicalId, 10)
	}
	if f.NotEmptyInvNumber {
		sql = sql + " and eo.inv_number!=''"
	}
	if f.TicketNumber != nil {
		sql = sql + " and eo.reader_ticket_num=" + strconv.Itoa(*f.TicketNumber)
	}
	if len(f.InStateCodes) > 0 {
		sqlState := ""
		for _, inState := range f.InStateCodes {
			sqlState = sqlState + " or ds.code='" + inState + "'"
		}
		sql = sql + " and (" + sqlState[3:] + ")"
	}
	if sql != "" {
		sql = "where " + sql[4:]
	}
	return sql
}
