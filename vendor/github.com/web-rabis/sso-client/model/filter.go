package model

import (
	"net/http"

	"github.com/web-rabis/sso-client/protobuf"
)

type UserFilters struct {
	Name string
}

func UserFiltersParseFromHttp(r *http.Request) *UserFilters {
	return &UserFilters{
		Name: r.URL.Query().Get("name"),
	}
}

func (f *UserFilters) ToProto() *protobuf.UserFilters {
	if f == nil {
		return nil
	}
	return &protobuf.UserFilters{
		Name: f.Name,
	}
}
