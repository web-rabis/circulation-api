package model

import "github.com/web-rabis/ebook-client/protobuf"

type EbookBbkM struct {
	Id      int64  `json:"id"`
	EbookId int64  `json:"ebookId"`
	Index   string `json:"index"`
}

type EbookBbkN struct {
	Id      int64  `json:"id"`
	EbookId int64  `json:"ebookId"`
	Index   string `json:"index"`
	Source  string `json:"source"`
}

func NewEbookBbkMFromResult(v []any) *EbookBbkM {
	e := &EbookBbkM{}
	e.Id = int64(v[0].(int32))
	e.EbookId = int64(v[1].(int32))
	if v[2] != nil {
		e.Index = v[2].(string)
	}
	return e
}
func NewEbookBbkMFromProto(s *protobuf.EbookBbkM) *EbookBbkM {
	if s == nil {
		return nil
	}
	return &EbookBbkM{
		Id:      s.Id,
		EbookId: s.EbookId,
		Index:   s.Index,
	}
}
func NewEbookBbkMsFromProto(s []*protobuf.EbookBbkM) []*EbookBbkM {
	if s == nil {
		return nil
	}
	var result []*EbookBbkM
	for _, v := range s {
		result = append(result, NewEbookBbkMFromProto(v))
	}
	return result
}
func NewEbookBbkNFromProto(s *protobuf.EbookBbkN) *EbookBbkN {
	if s == nil {
		return nil
	}
	return &EbookBbkN{
		Id:      s.Id,
		EbookId: s.EbookId,
		Index:   s.Index,
	}
}
func NewEbookBbkNsFromProto(s []*protobuf.EbookBbkN) []*EbookBbkN {
	if s == nil {
		return nil
	}
	var result []*EbookBbkN
	for _, v := range s {
		result = append(result, NewEbookBbkNFromProto(v))
	}
	return result
}
