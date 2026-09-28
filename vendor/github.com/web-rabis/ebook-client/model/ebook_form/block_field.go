package ebook_form

import (
	"fmt"

	"github.com/web-rabis/ebook-client/protobuf"
	"google.golang.org/protobuf/types/known/structpb"
)

type BlockField struct {
	FieldId int64           `json:"fieldId"`
	Value   BlockFieldValue `json:"value"`
}

func NewBlockFieldsFromProto(v []*protobuf.EbookFormBlockField) []*BlockField {
	result := make([]*BlockField, 0, len(v))
	for _, e := range v {
		entry, err := NewBlockFieldFromProto(e)
		if err != nil {
			continue
		}
		result = append(result, entry)
	}
	return result
}

// NewBlockFieldValueEntryFromProto разбирает oneof.
// Незаданный вариант - это не null: null приезжает явным null_value, а пустой
// oneof означает вариант, добавленный в проту позже и этой версии неизвестный.
// Молча превратить такое в null нельзя - данные бы исчезли, - поэтому ошибка.
func NewBlockFieldFromProto(v *protobuf.EbookFormBlockField) (*BlockField, error) {
	if v == nil {
		return nil, nil
	}

	entry := &BlockField{FieldId: v.FieldId}

	switch value := v.Value.(type) {
	case *protobuf.EbookFormBlockField_NullValue:
		entry.Value = NullBlockFieldValue()
	case *protobuf.EbookFormBlockField_StringValue:
		entry.Value = StringBlockFieldValue(value.StringValue)
	case *protobuf.EbookFormBlockField_IntValue:
		entry.Value = IntBlockFieldValue(value.IntValue)
	case *protobuf.EbookFormBlockField_BoolValue:
		entry.Value = BoolBlockFieldValue(value.BoolValue)
	case *protobuf.EbookFormBlockField_DictionaryValue:
		entry.Value = DictionaryBlockFieldValue(NewDictionaryRefFromProto(value.DictionaryValue))
	default:
		return nil, fmt.Errorf("поле %d: неизвестный вариант значения %T", v.FieldId, v.Value)
	}

	return entry, nil
}

func (e *BlockField) ToProto() *protobuf.EbookFormBlockField {
	if e == nil {
		return nil
	}

	result := &protobuf.EbookFormBlockField{FieldId: e.FieldId}

	switch e.Value.kind {
	case BlockFieldValueKindString:
		result.Value = &protobuf.EbookFormBlockField_StringValue{StringValue: e.Value.str}
	case BlockFieldValueKindInt:
		result.Value = &protobuf.EbookFormBlockField_IntValue{IntValue: e.Value.integer}
	case BlockFieldValueKindBool:
		result.Value = &protobuf.EbookFormBlockField_BoolValue{BoolValue: e.Value.boolean}
	case BlockFieldValueKindDictionary:
		result.Value = &protobuf.EbookFormBlockField_DictionaryValue{DictionaryValue: e.Value.dictionary.ToProto()}
	default:
		result.Value = &protobuf.EbookFormBlockField_NullValue{NullValue: structpb.NullValue_NULL_VALUE}
	}

	return result
}
