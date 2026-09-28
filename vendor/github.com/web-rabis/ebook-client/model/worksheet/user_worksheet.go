package worksheet

import "github.com/web-rabis/ebook-client/protobuf"

// UserWorksheet - персональная настройка пользователя: какие block_field_id
// показывать при добавлении/редактировании записи для пары
// BLevelId/TypeDescId. В отличие от справочников в model/dictionary, это не
// общие данные, а состояние конкретного пользователя - отдельный домен, а
// не часть dictionary, хоть и ссылается на её id.
type UserWorksheet struct {
	BLevelId     int64   `json:"bLevelId"`
	TypeDescId   int64   `json:"typeDescId"`
	BlockFieldId []int64 `json:"blockFieldId"`
}

func NewUserWorksheetFromProto(s *protobuf.UserWorksheet) *UserWorksheet {
	if s == nil {
		return nil
	}
	return &UserWorksheet{
		BLevelId:     s.BLevelId,
		TypeDescId:   s.TypeDescId,
		BlockFieldId: s.BlockFieldId,
	}
}

// ToProto - для UserWorksheetSave: клиент отправляет рабочий лист целиком,
// в отличие от чтения (UserWorksheet/UserWorksheets), где сервер его только
// возвращает.
func (s *UserWorksheet) ToProto() *protobuf.UserWorksheet {
	if s == nil {
		return nil
	}
	return &protobuf.UserWorksheet{
		BLevelId:     s.BLevelId,
		TypeDescId:   s.TypeDescId,
		BlockFieldId: s.BlockFieldId,
	}
}
