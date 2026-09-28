package ebook_form

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// BlockFieldValueKind - какой из вариантов значения заполнен.
type BlockFieldValueKind uint8

const (
	BlockFieldValueKindNull BlockFieldValueKind = iota
	BlockFieldValueKindString
	BlockFieldValueKindInt
	BlockFieldValueKindBool
	BlockFieldValueKindDictionary
)

// BlockFieldValue - Go-представление oneof из protobuf.
// Поля закрыты: значение создаётся конструктором и читается аксессором,
// поэтому заполненными не могут оказаться сразу два варианта.
type BlockFieldValue struct {
	kind       BlockFieldValueKind
	str        string
	integer    int64
	boolean    bool
	dictionary *DictionaryRef
}

func NullBlockFieldValue() BlockFieldValue {
	return BlockFieldValue{kind: BlockFieldValueKindNull}
}

func StringBlockFieldValue(s string) BlockFieldValue {
	return BlockFieldValue{kind: BlockFieldValueKindString, str: s}
}

func IntBlockFieldValue(i int64) BlockFieldValue {
	return BlockFieldValue{kind: BlockFieldValueKindInt, integer: i}
}

func BoolBlockFieldValue(b bool) BlockFieldValue {
	return BlockFieldValue{kind: BlockFieldValueKindBool, boolean: b}
}

// DictionaryBlockFieldValue при ref == nil даёт null: справочное поле без
// выбранного значения и отсутствие значения - это одно и то же.
func DictionaryBlockFieldValue(ref *DictionaryRef) BlockFieldValue {
	if ref == nil {
		return NullBlockFieldValue()
	}
	return BlockFieldValue{kind: BlockFieldValueKindDictionary, dictionary: ref}
}

func (v BlockFieldValue) Kind() BlockFieldValueKind { return v.kind }

func (v BlockFieldValue) IsNull() bool { return v.kind == BlockFieldValueKindNull }

// Str и остальные аксессоры возвращают false, если заполнен другой вариант.
func (v BlockFieldValue) Str() (string, bool) {
	return v.str, v.kind == BlockFieldValueKindString
}

func (v BlockFieldValue) Int() (int64, bool) {
	return v.integer, v.kind == BlockFieldValueKindInt
}

func (v BlockFieldValue) Bool() (bool, bool) {
	return v.boolean, v.kind == BlockFieldValueKindBool
}

func (v BlockFieldValue) Dictionary() (*DictionaryRef, bool) {
	return v.dictionary, v.kind == BlockFieldValueKindDictionary
}

// MarshalJSON с приёмником-значением, а не указателем: Value в
// BlockFieldValueEntry лежит значением, и с указательным приёмником
// encoding/json этот метод бы просто не вызвал.
func (v BlockFieldValue) MarshalJSON() ([]byte, error) {
	switch v.kind {
	case BlockFieldValueKindString:
		return json.Marshal(v.str)
	case BlockFieldValueKindInt:
		return json.Marshal(v.integer)
	case BlockFieldValueKindBool:
		return json.Marshal(v.boolean)
	case BlockFieldValueKindDictionary:
		return json.Marshal(v.dictionary)
	default:
		return []byte("null"), nil
	}
}

// UnmarshalJSON разбирает вариант по первому значащему байту: варианты
// не пересекаются как JSON-типы, поэтому этого достаточно и метаданные
// поля не нужны.
func (v *BlockFieldValue) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return errors.New("значение поля блока: пустой JSON")
	}

	switch trimmed[0] {
	case 'n':
		if !bytes.Equal(trimmed, []byte("null")) {
			return fmt.Errorf("значение поля блока: непонятный литерал %s", trimmed)
		}
		*v = NullBlockFieldValue()

	case '"':
		var s string
		if err := json.Unmarshal(trimmed, &s); err != nil {
			return err
		}
		*v = StringBlockFieldValue(s)

	case 't', 'f':
		var b bool
		if err := json.Unmarshal(trimmed, &b); err != nil {
			return err
		}
		*v = BoolBlockFieldValue(b)

	case '{':
		var ref DictionaryRef
		if err := json.Unmarshal(trimmed, &ref); err != nil {
			return err
		}
		// На записи от ссылки нужен только id, и без него она бессмысленна.
		if ref.Id == 0 {
			return errors.New("значение поля блока: у ссылки на справочник нет id")
		}
		*v = DictionaryBlockFieldValue(&ref)

	case '[':
		return errors.New("значение поля блока: массив не поддерживается")

	default:
		var n json.Number
		if err := json.Unmarshal(trimmed, &n); err != nil {
			return fmt.Errorf("значение поля блока: непонятный тип %s", trimmed)
		}
		// В схеме из числовых типов только INT, дробное здесь - ошибка данных,
		// а не повод молча отбросить дробную часть.
		i, err := n.Int64()
		if err != nil {
			return fmt.Errorf("значение поля блока: %s не целое число", n)
		}
		*v = IntBlockFieldValue(i)
	}

	return nil
}
