package model

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/web-rabis/ebook-client/protobuf"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type InvFilters struct {
	EbookId int64
}
type EbookFilters struct {
	Statuses             []string
	StatusIds            []int64
	CreatedAtLow         *time.Time
	CreatedAtHigh        *time.Time
	UpdatedAtLow         *time.Time
	UpdatedAtHigh        *time.Time
	CatalogId            *int64
	BibliographicLevelId *int64
	TypeDescriptionId    *int64
}

func EbookFiltersParseFromHttp(r *http.Request) *EbookFilters {
	var statuses []string
	if s := r.URL.Query().Get("statuses"); s != "" {
		statuses = strings.Split(s, ",")
	}
	return &EbookFilters{
		Statuses:             statuses,
		StatusIds:            make([]int64, 0),
		CreatedAtLow:         parseTimeQueryParam(r, "createdAtLow"),
		CreatedAtHigh:        parseTimeQueryParam(r, "createdAtHigh"),
		UpdatedAtLow:         parseTimeQueryParam(r, "updatedAtLow"),
		UpdatedAtHigh:        parseTimeQueryParam(r, "updatedAtHigh"),
		CatalogId:            parseInt64QueryParam(r, "catalogId"),
		BibliographicLevelId: parseInt64QueryParam(r, "bibliographicLevelId"),
		TypeDescriptionId:    parseInt64QueryParam(r, "typeDescriptionId"),
	}
}

func (f *InvFilters) ToProto() *protobuf.InvFilters {
	if f == nil {
		return nil
	}
	return &protobuf.InvFilters{
		EbookId: f.EbookId,
	}
}
func (f *EbookFilters) ToProto() *protobuf.EbookFilters {
	var catalogId int64
	var bibliographicLevelId int64
	var typeDescriptionId int64
	var createdAtLow *timestamppb.Timestamp
	var createdAtHigh *timestamppb.Timestamp
	var updatedAtLow *timestamppb.Timestamp
	var updatedAtHigh *timestamppb.Timestamp
	if f.CatalogId != nil {
		catalogId = *f.CatalogId
	}
	if f.BibliographicLevelId != nil {
		bibliographicLevelId = *f.BibliographicLevelId
	}
	if f.TypeDescriptionId != nil {
		typeDescriptionId = *f.TypeDescriptionId
	}
	if f.CreatedAtLow != nil {
		createdAtLow = timestamppb.New(*f.CreatedAtLow)
	}
	if f.CreatedAtHigh != nil {
		createdAtHigh = timestamppb.New(*f.CreatedAtHigh)
	}
	if f.UpdatedAtLow != nil {
		updatedAtLow = timestamppb.New(*f.UpdatedAtLow)
	}
	if f.UpdatedAtHigh != nil {
		updatedAtHigh = timestamppb.New(*f.UpdatedAtHigh)
	}
	return &protobuf.EbookFilters{
		Statuses:             f.Statuses,
		StatusIds:            f.StatusIds,
		CreatedAtLow:         createdAtLow,
		CreatedAtHigh:        createdAtHigh,
		UpdatedAtLow:         updatedAtLow,
		UpdatedAtHigh:        updatedAtHigh,
		CatalogId:            catalogId,
		BibliographicLevelId: bibliographicLevelId,
		TypeDescriptionId:    typeDescriptionId,
	}
}
func parseInt64QueryParam(r *http.Request, name string) *int64 {
	id := r.URL.Query().Get(name)
	if id == "" {
		return nil
	}
	intValue, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil
	}
	return &intValue
}
func parseTimeQueryParam(r *http.Request, name string) *time.Time {
	id := r.URL.Query().Get(name)
	if id == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", id)
	if err != nil {
		return nil
	}
	return &t
}
