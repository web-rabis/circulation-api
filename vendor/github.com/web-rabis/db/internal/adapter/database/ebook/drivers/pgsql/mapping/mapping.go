package mapping

import (
	"github.com/jackc/pgx/v5"
	"github.com/web-rabis/db/ebook"
	"github.com/web-rabis/db/ebook/dictionary"
	"github.com/web-rabis/db/internal/adapter/database/orm"
	"reflect"
)

func MappingObjects(t reflect.Type, v reflect.Value, result pgx.Rows, fdm map[string]int, bson string, isPtr bool) {
	if t.Name() == "Catalog" {
		sc := orm.NewObjectFromResult(&ebook.Catalog{}, result, bson+"_", MappingObjects).(*ebook.Catalog)
		if isPtr {
			v.Set(reflect.ValueOf(sc))
		} else {
			v.Set(reflect.ValueOf(*sc))
		}
	}
	if t.Name() == "DState" {
		sc := orm.NewObjectFromResult(&dictionary.DState{}, result, bson+"_", MappingObjects).(*dictionary.DState)
		if isPtr {
			v.Set(reflect.ValueOf(sc))
		} else {
			v.Set(reflect.ValueOf(*sc))
		}
	}
	if t.Name() == "Dictionary" {
		sc := orm.NewObjectFromResult(&dictionary.Dictionary{}, result, bson+"_", MappingObjects).(*dictionary.Dictionary)
		if isPtr {
			v.Set(reflect.ValueOf(sc))
		} else {
			v.Set(reflect.ValueOf(*sc))
		}
	}
	if t.Name() == "BibliographicLevel" {
		sc := orm.NewObjectFromResult(&dictionary.BibliographicLevel{}, result, bson+"_", MappingObjects).(*dictionary.BibliographicLevel)
		if isPtr {
			v.Set(reflect.ValueOf(sc))
		} else {
			v.Set(reflect.ValueOf(*sc))
		}
	}
	if t.Name() == "TypeDescription" {
		sc := orm.NewObjectFromResult(&dictionary.TypeDescription{}, result, bson+"_", MappingObjects).(*dictionary.TypeDescription)
		if isPtr {
			v.Set(reflect.ValueOf(sc))
		} else {
			v.Set(reflect.ValueOf(*sc))
		}
	}
}
