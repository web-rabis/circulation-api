package dictionary

import "github.com/web-rabis/ebook-client/protobuf"

type Catalog struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
type CatalogFilters struct {
	IsSearch *bool
}

func NewCatalogFromProto(c *protobuf.Catalog) *Catalog {
	if c == nil {
		return nil
	}
	return &Catalog{
		Id:   c.Id,
		Code: c.Code,
		Name: c.Name,
	}
}
func (f *CatalogFilters) ToProto() *protobuf.CatalogFilters {
	if f == nil {
		return nil
	}
	var isSearch int64 = -1
	if f.IsSearch != nil {
		if *f.IsSearch {
			isSearch = 1
		} else {
			isSearch = 0
		}
	}
	return &protobuf.CatalogFilters{
		IsSearch: isSearch,
	}
}
