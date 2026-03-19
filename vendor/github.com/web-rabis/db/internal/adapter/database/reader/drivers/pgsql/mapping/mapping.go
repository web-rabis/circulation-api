package mapping

import (
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/web-rabis/db/internal/adapter/database/orm"
	"github.com/web-rabis/db/reader"
)

func MappingObjects(t reflect.Type, v reflect.Value, result pgx.Rows, fdm map[string]int, bson string, isPtr bool) {
	if t.Name() == "DSocialStatus" {
		sc := orm.NewObjectFromResult(&reader.DSocialStatus{}, result, bson+"_", MappingObjects).(*reader.DSocialStatus)
		if isPtr {
			v.Set(reflect.ValueOf(sc))
		} else {
			v.Set(reflect.ValueOf(*sc))
		}
	} else if t.Name() == "Dictionary" {
		sc := orm.NewObjectFromResult(&reader.Dictionary{}, result, bson+"_", MappingObjects).(*reader.Dictionary)
		if isPtr {
			v.Set(reflect.ValueOf(sc))
		} else {
			v.Set(reflect.ValueOf(*sc))
		}
	} else if t.Name() == "DTypeCard" {
		sc := orm.NewObjectFromResult(&reader.DTypeCard{}, result, bson+"_", MappingObjects).(*reader.DTypeCard)
		if isPtr {
			v.Set(reflect.ValueOf(sc))
		} else {
			v.Set(reflect.ValueOf(*sc))
		}
	}
}
