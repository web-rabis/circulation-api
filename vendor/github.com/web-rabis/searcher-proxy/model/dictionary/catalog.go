package dictionary

import "github.com/web-rabis/searcher-proxy/protobuf"

type Catalog struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
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
