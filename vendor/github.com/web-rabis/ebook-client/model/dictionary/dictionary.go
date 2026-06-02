package dictionary

import "github.com/web-rabis/ebook-client/protobuf"

type Dictionary struct {
	Id        int64  `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	TableName string `json:"tableName"`
	Base      bool   `json:"base"`
}
type DictionaryBase struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
type DictionaryFilters struct{}
type SearchDictionaryFilters struct {
	Dictionary *Dictionary `json:"dictionary"`
	Query      string      `json:"query"`
}

func (d *Dictionary) ToProto() *protobuf.Dictionary {
	if d == nil {
		return nil
	}
	return &protobuf.Dictionary{
		Id:        d.Id,
		Code:      d.Code,
		Name:      d.Name,
		TableName: d.TableName,
		Base:      d.Base,
	}
}
func (f *DictionaryFilters) ToProto() *protobuf.DictionaryFilters {
	if f == nil {
		return nil
	}
	return &protobuf.DictionaryFilters{}
}
func (f *SearchDictionaryFilters) ToProto() *protobuf.SearchDictionaryFilters {
	if f == nil {
		return nil
	}
	return &protobuf.SearchDictionaryFilters{
		Dictionary: f.Dictionary.ToProto(),
		Query:      f.Query,
	}
}
func NewDictionaryFromProto(c *protobuf.Dictionary) *Dictionary {
	if c == nil {
		return nil
	}
	return &Dictionary{
		Id:        c.Id,
		Code:      c.Code,
		Name:      c.Name,
		TableName: c.TableName,
		Base:      c.Base,
	}
}
func NewDictionaryBaseFromProto(c *protobuf.DictionaryBase) *DictionaryBase {
	if c == nil {
		return nil
	}
	return &DictionaryBase{
		Id:   c.Id,
		Code: c.Code,
		Name: c.Name,
	}
}
