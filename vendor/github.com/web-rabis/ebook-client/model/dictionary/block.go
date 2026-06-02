package dictionary

import "github.com/web-rabis/ebook-client/protobuf"

type Block struct {
	Id            int64
	Name          string
	IsRepeat      bool
	Priority      int64
	ExternalTable *string
	KeyValue      *int64
}
type BlockFilters struct {
}

func NewBlockFromProto(s *protobuf.Block) *Block {
	if s == nil {
		return nil
	}
	var externalTable *string
	if s.ExternalTable != "" {
		externalTable = &s.ExternalTable
	}
	var keyValue *int64
	if s.KeyValue != 0 {
		keyValue = &s.KeyValue
	}
	return &Block{
		Id:            s.Id,
		Name:          s.Name,
		IsRepeat:      s.IsRepeat,
		Priority:      s.Priority,
		ExternalTable: externalTable,
		KeyValue:      keyValue,
	}
}
func (f *BlockFilters) ToProto() *protobuf.BlockFilters {
	if f == nil {
		return nil
	}
	return &protobuf.BlockFilters{}
}
