package dictionary

import "github.com/web-rabis/ebook-client/protobuf"

type BlockField struct {
	Id          int64  `json:"id§"`
	BlockId     int64  `json:"blockId"`
	Name        string `json:"name"`
	FieldName   string `json:"fieldName"`
	FieldType   string `json:"fieldType"`
	Priority    int64  `json:"priority"`
	DirectoryId *int64 `json:"directoryId"`
	Required    bool   `json:"required"`
}
type BlockFieldFilters struct {
	IsSearch *bool
}

func NewBlockFieldFromProto(s *protobuf.BlockField) *BlockField {
	if s == nil {
		return nil
	}
	var directoryId *int64
	if s.DirectoryId != 0 {
		directoryId = &s.DirectoryId
	}
	return &BlockField{
		Id:          s.Id,
		BlockId:     s.BlockId,
		Name:        s.Name,
		FieldName:   s.FieldName,
		FieldType:   s.FieldType,
		Priority:    s.Priority,
		DirectoryId: directoryId,
		Required:    s.Required,
	}
}
func (f *BlockFieldFilters) ToProto() *protobuf.BlockFieldFilters {
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
	return &protobuf.BlockFieldFilters{
		IsSearch: isSearch,
	}
}
