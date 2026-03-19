package mapping

import (
	"reflect"

	"github.com/jackc/pgx/v5"
	
	"github.com/web-rabis/db/internal/adapter/database/orm"
	"github.com/web-rabis/db/user"
)

func MappingObjects(t reflect.Type, v reflect.Value, result pgx.Rows, fdm map[string]int, bson string, isPtr bool) {

	if t.Name() == "Department" {
		sc := orm.NewObjectFromResult(&user.Department{}, result, bson+"_", MappingObjects).(*user.Department)
		if isPtr {
			v.Set(reflect.ValueOf(sc))
		} else {
			v.Set(reflect.ValueOf(*sc))
		}
	}

}
