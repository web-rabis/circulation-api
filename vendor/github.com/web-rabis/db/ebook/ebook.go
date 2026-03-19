package ebook

import (
	"time"

	"github.com/web-rabis/db/ebook/dictionary"
)

type Ebook struct {
	Id                 int64                          `json:"id" bson:"id"`
	CreateDate         time.Time                      `json:"create_date" bson:"create_date"`                        // Дата создания
	EditDate           time.Time                      `json:"edit_date" bson:"edit_date"`                            // Дата редактирования
	Catalog            *Catalog                       `json:"catalog" bson:"catalog_id" foreignTable:"catalog"`      // Код каталога
	State              *dictionary.DState             `json:"state" bson:"state_id" foreignTable:"dictionary_state"` // Код состояния
	Krv                bool                           `json:"krv" bson:"krv"`                                        // Признак краеведения
	Digest             bool                           `json:"digest" bson:"digest"`                                  // Признак краеведения
	BibliographicLevel *dictionary.BibliographicLevel `json:"b_level_id" bson:"b_level_id"`                          // Код библиографического уровня
	TypeDescriptionId  *dictionary.TypeDescription    `json:"type_descr_id" bson:"type_descr_id"`                    // Код типа описания
	VolumeNumber       int64                          `json:"volume_number" bson:"volume_number"`                    // Номер тома
	CreateUser         int64                          `json:"create_user" bson:"create_user"`                        // ID пользователя создавшего
	CreateUserFam      string                         `json:"create_user_fam" bson:"create_user_fam"`                // Fam пользователя создавшего
	Author             string                         `json:"author" bson:"author"`                                  // Автор
	Title              string                         `json:"title" bson:"title"`                                    // Заглавие
	//OtherAuthors       []*EbookAuthor                 `json:"other_authors" bson:"other_authors"`                    // Другие авторы
	//RegularAuthors       *Authors              `json:"regular_authors"`                                  // Постоянные авторы
	//TemporaryAuthors     *Authors              `json:"temporary_authors"`                                // Временные авторы
	//Titles               []*Title              `json:"titles"`                                           // Заглавия
	//InfoPublishing       *InfoPublishing       `json:"info_publishing"`                                  // Сведения об издательстве
	//Publishes            []*Publishing         `json:"publishes"`                                        // Издательство
	//Volume               *Volume               `json:"volume"`                                           // Обьем
	//Series               []*Series             `json:"series"`                                           // История
	//Sources              []*Source             `json:"sources"`                                          // Источник
	//Notes                []*Note               `json:"notes"`                                            // Примечания
	//Summary              string                `json:"summary" bson:"summary"`                           // Реферат
	//DataOuts             []*DataOut            `json:"data_outs" bson:"data_outs"`                       // Выходные данные
	//PeriodicalIssn       string                `json:"periodical_issn" bson:"periodical_issn"`           // ISSN
	//AddCards             []string              `json:"add_cards" bson:"add_cards"`                       // Добавочная карточка
	//Rubrics              []*Rubric             `json:"rubrics" bson:"rubrics"`                           // Рубрика
	//Keywords             []string              `json:"keywords" bson:"keywords"`                         // Ключевые слова
	//Udk                  []string              `json:"udk" bson:"udk"`                                   // УДК
	//BbkM                 []string              `json:"bbk_m" bson:"bbk_m"`                               //ББК Научный
	//BbkN                 []*BbkN               `json:"bbk_n" bson:"bbk_n"`                               // ББК Массовый
	//Vak                  []string              `json:"vak" bson:"vak"`                                   // ВАК
	//PersonalClassifiers  []string              `json:"personal_classifiers" json:"personal_classifiers"` // Индекс лич. классифик-ра
	//Sdk                  []string              `json:"sdk" bson:"sdk"`                                   // СДК
	//AuthorMark           string                `json:"author_mark" bson:"author_mark"`                   // Авторский знак
	//Format               *Format               `json:"format" bson:"format"`                             // Формат
	//PlacementId          int64                 `json:"placement" bson:"placement"`                       // Расстановка
	//AddRefs              []*AddRef             `json:"add_refs" bson:"add_refs"`                         // Дополнительные ссылки
	//Specifications       []string              `json:"specifications" bson:"specifications"`             // Спецификация
	//InvNumbers           []*Inv                `json:"inv_numbers" bson:"inv_numbers"`                   // Инвертарный номер
	//Language             *Language             `json:"language" bson:"language"`                         // Язык
	//OtherLanguages       []Language `json:"other_languages" bson:"other_languages"`           // Другой язык(Спр)
	//Scale                string                `json:"scale" bson:"scale"`                               // Масштаб
	//ServiceNotes         []*primitive.ObjectID `json:"service_notes" bson:"service_notes"`               // Служебные отметки(Спр)
	//Texts                []*Text               `json:"texts" bson:"texts"`                               // Текст
	//Siglas               []*primitive.ObjectID `json:"siglas" bson:"siglas"`                             // Сигла(Спр)
	//Orders               []*Order              `json:"orders" bson:"orders"`                             // СП.СВ
	//CopyCount            string                `json:"copy_count" bson:"copy_count"`                     // Количество экземпляров
	//PartNumbers          []string              `json:"part_numbers" bson:"part_numbers"`                 // Номер партии
	//Editions             []*Edition            `json:"editions" bson:"editions"`                         // Вид издательства
	//Booking              *Booking              `json:"booking" bson:"booking"`                           // Заказ
	//StorageSiglas        []*primitive.ObjectID `json:"storage_siglas" bson:"storage_siglas"`             // Сигла Хранения
	//Transfers            []*Transfer           `json:"transfers" bson:"transfers"`                       // Перевод
	//Country              *primitive.ObjectID   `json:"country" bson:"country"`                           // Страна(Спр)
	//Original             *Original             `json:"original" bson:"original"`                         // Оригинал
	//Educationals         []*Educational        `json:"educationals" bson:"educationals"`                 // Учебная литература
}
