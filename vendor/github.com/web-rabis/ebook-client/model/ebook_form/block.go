package ebook_form

import "github.com/web-rabis/ebook-client/protobuf"

type Block struct {
	Id      *int64        `json:"id"`
	EbookId *int64        `json:"ebookId"`
	BlockId int64         `json:"blockId"`
	Values  []*BlockField `json:"values"`
	Deleted bool          `json:"deleted"`
}

func NewBlockFromProto(b *protobuf.EbookFormBlock) *Block {
	if b == nil {
		return nil
	}
	return &Block{
		Id:      b.Id,
		EbookId: b.EbookId,
		BlockId: b.BlockId,
		Values:  NewBlockFieldsFromProto(b.Values),
		Deleted: b.Deleted,
	}
}

func NewBlocksFromProto(b []*protobuf.EbookFormBlock) []*Block {
	if b == nil {
		return nil
	}
	var result []*Block
	for _, v := range b {
		result = append(result, NewBlockFromProto(v))
	}
	return result
}

func (b *Block) ToProto() *protobuf.EbookFormBlock {
	if b == nil {
		return nil
	}
	values := make([]*protobuf.EbookFormBlockField, 0, len(b.Values))
	for _, v := range b.Values {
		values = append(values, v.ToProto())
	}
	return &protobuf.EbookFormBlock{
		Id:      b.Id,
		EbookId: b.EbookId,
		BlockId: b.BlockId,
		Values:  values,
		Deleted: b.Deleted,
	}
}

func BlocksToProto(b []*Block) []*protobuf.EbookFormBlock {
	if b == nil {
		return nil
	}
	result := make([]*protobuf.EbookFormBlock, 0, len(b))
	for _, v := range b {
		result = append(result, v.ToProto())
	}
	return result
}
