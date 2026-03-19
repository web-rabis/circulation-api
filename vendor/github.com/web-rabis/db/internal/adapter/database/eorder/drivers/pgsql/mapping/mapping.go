package mapping

import (
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/web-rabis/db/eorder"
	"github.com/web-rabis/db/internal/adapter/database/orm"
)

func MappingObjects(t reflect.Type, v reflect.Value, result pgx.Rows, fdm map[string]int, bson string, isPtr bool) {

	if t.Name() == "Department" {
		sc := orm.NewObjectFromResult(&eorder.Department{}, result, bson+"_", MappingObjects).(*eorder.Department)
		if isPtr {
			v.Set(reflect.ValueOf(sc))
		} else {
			v.Set(reflect.ValueOf(*sc))
		}
	}
	if t.Name() == "DState" {
		sc := orm.NewObjectFromResult(&eorder.DState{}, result, bson+"_", MappingObjects).(*eorder.DState)
		if isPtr {
			v.Set(reflect.ValueOf(sc))
		} else {
			v.Set(reflect.ValueOf(*sc))
		}
	}
}
