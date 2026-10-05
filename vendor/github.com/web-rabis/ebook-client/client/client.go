package client

import (
	"context"

	"github.com/web-rabis/ebook-client/model"
	"github.com/web-rabis/ebook-client/model/dictionary"
	"github.com/web-rabis/ebook-client/model/ebook"
	"github.com/web-rabis/ebook-client/model/ebook_form"
	"github.com/web-rabis/ebook-client/model/worksheet"
)

type Base interface {
	Connect() error
	Close() error
	EbookSvc() EbookService
	EbookFormSvc() EbookFormService
	DictionarySvc() DictionaryService
	WorksheetSvc() WorksheetService
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name EbookService
type EbookService interface {
	EbookList(ctx context.Context, filters *model.EbookFilters, paging *model.Paging) (int64, []*ebook.Ebook, error)
	EbookById(ctx context.Context, id int64, withCard bool) (*ebook.Ebook, error)
	EbookCardById(ctx context.Context, id int64) (*ebook.EbookCard, error)
	InvList(ctx context.Context, filters *model.InvFilters, paging *model.Paging) (int64, []*ebook.Inv, error)
	// InvBarcodeSave сохраняет barcode экземпляра id, не трогая остальные
	// поля. Возвращает экземпляр с уже применённым изменением.
	InvBarcodeSave(ctx context.Context, id int64, barcode string) (*ebook.Inv, error)
	// EbookDelete - "мягкое" удаление: запись не удаляется физически, а
	// переводится в статус Ebook.Deleted; userId - кто удалил, записывается
	// в ebook_audit. Возвращает запись уже в новом статусе.
	EbookDelete(ctx context.Context, id, userId int64) (*ebook.Ebook, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name EbookFormService
type EbookFormService interface {
	EbookFormById(ctx context.Context, id int64) (*ebook_form.EbookForm, error)
	// EbookFormSave создаёт или обновляет EbookForm целиком (заголовок и
	// блоки). Create/update определяется по form.Id (0 - новая запись) и,
	// для каждого блока, по Block.Id/Block.EbookId. Возвращает сохранённую
	// форму с проставленными id.
	// userId - кто сохраняет: при создании пишется в ebook.create_user, при
	// редактировании туда не пишется - в обоих случаях добавляется запись в
	// ebook_audit.
	EbookFormSave(ctx context.Context, form *ebook_form.EbookForm, userId int64) (*ebook_form.EbookForm, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.53 --name DictionaryService
type DictionaryService interface {
	DictionaryList(ctx context.Context, filters *dictionary.DictionaryFilters, paging *model.Paging) (int64, []*dictionary.Dictionary, error)
	SearchDictionary(ctx context.Context, filters *dictionary.SearchDictionaryFilters, paging *model.Paging) (int64, []*dictionary.DictionaryBase, error)
	CatalogList(ctx context.Context, filters *dictionary.CatalogFilters, paging *model.Paging) (int64, []*dictionary.Catalog, error)
	TypeDescriptionList(ctx context.Context, filters *dictionary.TypeDescriptionFilters, paging *model.Paging) (int64, []*dictionary.TypeDescription, error)
	BibliographicLevelList(ctx context.Context, filters *dictionary.BibliographicLevelFilters, paging *model.Paging) (int64, []*dictionary.BibliographicLevel, error)
	LanguageList(ctx context.Context, filters *dictionary.LanguageFilters, paging *model.Paging) (int64, []*dictionary.Language, error)
	BlockList(ctx context.Context, filters *dictionary.BlockFilters, paging *model.Paging) (int64, []*dictionary.Block, error)
	BlockFieldList(ctx context.Context, filters *dictionary.BlockFieldFilters, paging *model.Paging) (int64, []*dictionary.BlockField, error)
	StateList(ctx context.Context, filters *dictionary.StateFilters, paging *model.Paging) (int64, []*dictionary.State, error)
}

// WorksheetService - рабочие листы пользователей: персональная настройка,
// какие поля блоков показывать при добавлении/редактировании записи для
// конкретной пары bLevelId/typeDescId. Отдельно от DictionaryService, так
// как это не общий справочник, а данные конкретного пользователя, которые
// он сам изменяет (UserWorksheetSave).
//
//go:generate go run github.com/vektra/mockery/v2@v2.53 --name WorksheetService
type WorksheetService interface {
	// UserWorksheet возвращает рабочий лист пользователя userId — набор
	// block_field_id, которые нужно показывать при добавлении/редактировании
	// записи с данным bLevelId/typeDescId.
	UserWorksheet(ctx context.Context, userId, bLevelId, typeDescId int64) (*worksheet.UserWorksheet, error)
	// UserWorksheets возвращает все рабочие листы пользователя userId сразу для
	// всех комбинаций b_level_id/type_desc_id.
	UserWorksheets(ctx context.Context, userId int64) ([]*worksheet.UserWorksheet, error)
	// UserWorksheetSave сохраняет рабочий лист пользователя userId целиком:
	// старый набор block_field_id для worksheet.BLevelId/worksheet.TypeDescId
	// полностью заменяется новым (worksheet.BlockFieldId). Пустой
	// BlockFieldId - легитимный случай ("очистить рабочий лист"). Возвращает
	// сохранённый рабочий лист.
	UserWorksheetSave(ctx context.Context, userId int64, worksheet *worksheet.UserWorksheet) (*worksheet.UserWorksheet, error)
}
