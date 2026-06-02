package dictionary

import "github.com/web-rabis/ebook-client/protobuf"

type BlockField struct {
	Id          int64  `json:"id§"`
	BlockId     int64  `json:"blockId"`
	Name        string `json:"name"`
	FieldName   string `json:"fieldName"`
	FieldType   string `json:"fieldType"`
	DirectoryId *int64 `json:"directoryId"`
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
		DirectoryId: directoryId,
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
