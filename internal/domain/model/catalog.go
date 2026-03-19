package model

type Catalog struct {
	Id            int64   `json:"id"`
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	ReaderCatalog *string `json:"readerCatalog"`
}

type CatalogWeb struct {
	Id             int64       `json:"id"`
	CatalogId      int64       `json:"catalogId"`
	Name           LangOptions `json:"name"`
	Priority       int         `json:"priority"`
	IsDownloadText bool        `json:"isDownloadText"`
}
type LangOptions struct {
	EN string `json:"en" bson:"en"`
	RU string `json:"ru" bson:"ru"`
	KZ string `json:"kz" bson:"kz"`
}

func NewCatalogFromResult(v []any) Catalog {
	c := Catalog{}
	c.Id = int64(v[0].(int32))
	if v[1] != nil {
		c.Code = v[1].(string)
	}
	if v[2] != nil {
		c.Name = v[2].(string)
	}
	return c
}
func NewCatalogFromResultElk(v interface{}) Catalog {
	if v == nil {
		return Catalog{}
	}
	vv := v.(map[string]interface{})
	readerCatalog := vv["readerCatalog"].(string)
	return Catalog{
		Id:            int64(vv["id"].(float64)),
		Code:          vv["code"].(string),
		Name:          vv["name"].(string),
		ReaderCatalog: &readerCatalog,
	}
}

func NewCatalogWebFromResult(v []any) CatalogWeb {
	c := CatalogWeb{}
	c.Id = int64(v[0].(int32))
	if v[1] != nil {
		c.Name.RU = v[1].(string)
	}
	if v[2] != nil {
		c.Name.KZ = v[2].(string)
	}
	if v[3] != nil {
		c.Name.EN = v[2].(string)
	}
	c.CatalogId = int64(v[4].(int32))
	c.Priority = int(v[5].(int32))
	if v[6] != nil {
		c.IsDownloadText = v[5].(int32) == 1
	}
	return c
}
