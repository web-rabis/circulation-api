package reader

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type ReaderUserFilter struct {
	Status string
	Query  string
}
type ReaderFilter struct {
	Status             string
	Query              string
	Iin                string
	ExpirationDateLow  *time.Time
	ExpirationDateHigh *time.Time
	IsEmployee         *bool
}
type ReaderEmployeeFilter struct {
	Status             string
	Query              string
	Iin                string
	ExpirationDateLow  *time.Time
	ExpirationDateHigh *time.Time
}
type ControlFilter struct {
	VisitDateLow   *time.Time
	VisitDateHigh  *time.Time
	ExitDateLow    *time.Time
	ExitDateHigh   *time.Time
	ExitDateIsNull *bool
}

func (f *ReaderUserFilter) Sql() string {
	if f == nil {
		return ""
	}
	var sql = ""
	if f.Status != "" {
		sql = "status = '" + f.Status + "'"
	}
	if f.Query != "" {
		if sql != "" {
			sql += " and "
		}
		sql = sql + `(
				id_number ilike '` + f.Query + `%' or 
				firstname ilike '` + f.Query + `%' or
				middlename ilike '` + f.Query + `%' or
				lastname ilike '` + f.Query + `%' or
				phone ilike '` + f.Query + `%' or
				work_phone ilike '` + f.Query + `%'
				)`
	}
	return sql
}

func ReaderUserFilterParseFromHttp(r *http.Request) (*ReaderUserFilter, error) {
	filter := ReaderUserFilter{}

	if status := r.URL.Query().Get("status"); status != "" && status != "*" {
		filter.Status = status
	}
	if query := r.URL.Query().Get("query"); query != "" {
		filter.Query = query
	}

	if filter.Sql() == "" {
		return nil, nil
	}
	return &filter, nil
}

func (f *ReaderFilter) Sql() string {
	if f == nil {
		return ""
	}
	var sql = ""
	if f.Status != "" {
		sql = "status = '" + f.Status + "'"
	} else {
		sql = "(status is null or status!='deleted') and priznakudal=''"
	}
	if f.Iin != "" {
		if sql != "" {
			sql += " and "
		}
		sql += " id_number='" + f.Iin + "'"
	}
	if f.IsEmployee != nil {
		if sql != "" {
			sql += " and "
		}
		sql += " is_employee=" + strconv.FormatBool(*f.IsEmployee)
	}
	if f.Query != "" {
		if sql != "" {
			sql += " and "
		}
		queryNumWhere := ""
		queryFioWhere := ""
		queryNum := parseNumber(f.Query)
		if len(queryNum) > 0 {
			queryNumWhere = `ticket_num = ` + queryNum + ` or 
				replace(phone,' ','') like '%` + queryNum + `%' or
				replace(work_phone,' ','') like '%` + queryNum + `%' or
				id_number like '` + f.Query + `%'
`
		}
		fio := trimSlice(strings.Split(f.Query, " "))
		if len(fio) == 1 {
			queryFioWhere = `firstname ilike '` + fio[0] + `%' or
				middlename ilike '` + fio[0] + `%' or
				lastname ilike '` + fio[0] + `%'`
		} else if len(fio) == 2 {
			queryFioWhere = `(lastname ilike '` + fio[0] + `%' and
				firstname ilike '` + fio[1] + `%')`
		} else if len(fio) > 2 {
			queryFioWhere = `(lastname ilike '` + fio[0] + `%' and
				firstname ilike '` + fio[1] + `%' and
				middlename ilike '` + fio[2] + `%')`

		}
		sql = sql + "(" + queryFioWhere
		if queryNumWhere != "" {
			sql = sql + " or " + queryNumWhere
		}
		sql = sql + ")"
	}
	if f.ExpirationDateLow != nil || f.ExpirationDateHigh != nil {
		if sql != "" {
			sql += " and "
		}
		if f.ExpirationDateLow != nil && f.ExpirationDateHigh != nil {
			sql = sql + "f48 between '" + f.ExpirationDateLow.Format("2006-01-02") + "' and '" + f.ExpirationDateHigh.Format("2006-01-02") + "'"
		} else if f.ExpirationDateLow != nil {
			sql = sql + "f48 >= '" + f.ExpirationDateLow.Format("2006-01-02") + "'"
		} else if f.ExpirationDateHigh != nil {
			sql = sql + "f48 <= '" + f.ExpirationDateHigh.Format("2006-01-02") + "'"
		}
	}
	return sql
}

func ReaderFilterParseFromHttp(r *http.Request) (*ReaderFilter, error) {
	filter := ReaderFilter{}

	if status := r.URL.Query().Get("status"); status != "" && status != "*" {
		filter.Status = status
	}
	if iin := r.URL.Query().Get("iin"); iin != "" {
		filter.Iin = iin
	}
	if expirationDate := r.URL.Query().Get("expirationDateLow"); expirationDate != "" {
		expDate, err := time.Parse("2006-01-02", expirationDate)
		if err == nil {
			filter.ExpirationDateLow = &expDate
		}
	}
	if expirationDate := r.URL.Query().Get("expirationDateHigh"); expirationDate != "" {
		expDate, err := time.Parse("2006-01-02", expirationDate)
		if err == nil {
			filter.ExpirationDateHigh = &expDate
		}
	}
	if query := r.URL.Query().Get("query"); query != "" {
		filter.Query = query
	}
	if isEmployeeStr := r.URL.Query().Get("isEmployee"); isEmployeeStr != "" {
		isEmployee, err := strconv.ParseBool(isEmployeeStr)
		if err == nil {
			filter.IsEmployee = &isEmployee
		}
	}
	return &filter, nil
}
func (f *ReaderEmployeeFilter) Sql() string {
	if f == nil {
		return ""
	}
	var sql = ""
	if f.Status != "" {
		sql = "status = '" + f.Status + "'"
	} else {
		sql = "(status != 'deleted' or priznakudal=='')"
	}
	if f.Iin != "" {
		if sql != "" {
			sql += " and "
		}
		sql += " id_number='" + f.Iin + "'"
	}
	if f.Query != "" {
		if sql != "" {
			sql += " and "
		}
		queryNumWhere := ""
		queryFioWhere := ""
		queryNum := parseNumber(f.Query)
		if len(queryNum) > 0 {
			queryNumWhere = `ticket_num = ` + queryNum + ` or 
				replace(phone,' ','') ilike '%` + queryNum + `%' or
				replace(work_phone,' ','') ilike '%` + queryNum + `%' or
				id_number ilike '` + f.Query + `%'
`
		}
		fio := trimSlice(strings.Split(f.Query, " "))
		if len(fio) == 1 {
			queryFioWhere = `firstname ilike '` + fio[0] + `%' or
				middlename ilike '` + fio[0] + `%' or
				lastname ilike '` + fio[0] + `%'`
		} else if len(fio) == 2 {
			queryFioWhere = `(lastname ilike '` + fio[0] + `%' and
				firstname ilike '` + fio[1] + `%')`
		} else if len(fio) > 2 {
			queryFioWhere = `(lastname ilike '` + fio[0] + `%' and
				firstname ilike '` + fio[1] + `%' and
				middlename ilike '` + fio[2] + `%')`

		}
		sql = sql + "(" + queryFioWhere
		if queryNumWhere != "" {
			sql = sql + " or " + queryNumWhere
		}
		sql = sql + ")"
	}
	if f.ExpirationDateLow != nil || f.ExpirationDateHigh != nil {
		if sql != "" {
			sql += " and "
		}
		if f.ExpirationDateLow != nil && f.ExpirationDateHigh != nil {
			sql = sql + "f48 between '" + f.ExpirationDateLow.Format("2006-01-02") + "' and '" + f.ExpirationDateHigh.Format("2006-01-02") + "'"
		} else if f.ExpirationDateLow != nil {
			sql = sql + "f48 >= '" + f.ExpirationDateLow.Format("2006-01-02") + "'"
		} else if f.ExpirationDateHigh != nil {
			sql = sql + "f48 <= '" + f.ExpirationDateHigh.Format("2006-01-02") + "'"
		}
	}
	return sql
}

func ReaderEmployeeFilterParseFromHttp(r *http.Request) (*ReaderEmployeeFilter, error) {
	filter := ReaderEmployeeFilter{}

	if status := r.URL.Query().Get("status"); status != "" && status != "*" {
		filter.Status = status
	}
	if iin := r.URL.Query().Get("iin"); iin != "" {
		filter.Iin = iin
	}
	if expirationDate := r.URL.Query().Get("expirationDateLow"); expirationDate != "" {
		expDate, err := time.Parse("2006-01-02", expirationDate)
		if err == nil {
			filter.ExpirationDateLow = &expDate
		}
	}
	if expirationDate := r.URL.Query().Get("expirationDateHigh"); expirationDate != "" {
		expDate, err := time.Parse("2006-01-02", expirationDate)
		if err == nil {
			filter.ExpirationDateHigh = &expDate
		}
	}

	if query := r.URL.Query().Get("query"); query != "" {
		filter.Query = query
	}
	return &filter, nil
}
func parseNumber(s string) string {
	// Оставляем только цифры
	digits := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		}
	}
	// Превращаем в строку с цифрами
	return string(digits)
}
func trimSlice(s []string) []string {
	result := make([]string, 0, len(s))
	for _, v := range s {
		if strings.TrimSpace(v) != "" {
			result = append(result, v)
		}
	}
	return result
}

func (f *ControlFilter) Sql() string {
	if f == nil {
		return ""
	}
	var sql = ""
	if f.VisitDateLow != nil || f.VisitDateHigh != nil {
		if sql != "" {
			sql += " and "
		}
		if f.VisitDateLow != nil && f.VisitDateHigh != nil {
			sql = sql + "vizit_date between '" + f.VisitDateLow.Format("2006-01-02 15:04:05") + "' and '" + f.VisitDateHigh.Format("2006-01-02 15:04:05") + "'"
		} else if f.VisitDateLow != nil {
			sql = sql + "vizit_date >= '" + f.VisitDateLow.Format("2006-01-02 15:04:05") + "'"
		} else if f.VisitDateHigh != nil {
			sql = sql + "vizit_date <= '" + f.VisitDateHigh.Format("2006-01-02 15:04:05") + "'"
		}
	}
	if f.ExitDateLow != nil || f.ExitDateHigh != nil {
		if sql != "" {
			sql += " and "
		}
		if f.ExitDateLow != nil && f.ExitDateHigh != nil {
			sql = sql + "vyhod between '" + f.ExitDateLow.Format("2006-01-02 15:04:05") + "' and '" + f.ExitDateHigh.Format("2006-01-02 15:04:05") + "'"
		} else if f.ExitDateLow != nil {
			sql = sql + "vyhod >= '" + f.ExitDateLow.Format("2006-01-02 15:04:05") + "'"
		} else if f.ExitDateHigh != nil {
			sql = sql + "vyhod <= '" + f.ExitDateHigh.Format("2006-01-02 15:04:05") + "'"
		}
	}
	if f.ExitDateIsNull != nil {
		if sql != "" {
			sql += " and "
		}
		if *f.ExitDateIsNull {
			sql = sql + "vyhod is null"
		} else {
			sql = sql + "vyhod is not null"
		}
	}
	return sql
}
