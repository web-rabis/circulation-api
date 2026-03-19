package control

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/gorm"

	"github.com/web-rabis/db/internal/adapter/database/orm"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/mapping"
	"github.com/web-rabis/db/model"
	"github.com/web-rabis/db/reader"
)

type Control struct {
	pool *pgxpool.Pool
	db   *gorm.DB
}

func New(pool *pgxpool.Pool) *Control {
	return &Control{
		pool: pool,
	}
}
func (r *Control) List(ctx context.Context, filter *reader.ControlFilter, paging *model.Paging) ([]*reader.Control, error) {
	var list []*reader.Control
	var f = strings.Join(orm.Fields(reader.Control{}).SqlFields(""), ",")

	var sql = "select " + f + " from nlrk_reader.control "
	if filter != nil {
		sql = sql + " where " + filter.Sql()
	}
	if paging != nil {
		sql = sql + paging.Sql()
	}

	result, err := r.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		item := orm.NewObjectFromResult(&reader.Control{}, result, "", mapping.MappingObjects).(*reader.Control)
		list = append(list, item)
	}
	return list, nil
}
func (r *Control) Count(ctx context.Context, filter *reader.ControlFilter) (int64, error) {

	var sql = "select count(control.ticket_num) from nlrk_reader.control"
	if filter != nil {
		sql += " where " + filter.Sql()
	}
	result, err := r.pool.Query(ctx, sql)
	if err != nil {
		return 0, err
	}
	defer result.Close()

	count := int64(0)

	for result.Next() {
		err = result.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}
func (r *Control) Create(ctx context.Context, control *reader.Control) error {
	var rFields = orm.Fields(reader.Control{})
	var rf = reflect.ValueOf(control).Elem()
	var sql = "insert into nlrk_reader.control(%s) values(%s)"
	fields, values := rFields.FieldsValues(rf)
	var vv []string
	for i, _ := range values {
		vv = append(vv, "$"+strconv.Itoa(i+1))
	}
	sql = fmt.Sprintf(sql, strings.Join(fields, ","), strings.Join(vv, ","))
	var err error
	_, err = r.pool.Exec(ctx, sql, values...)
	if err != nil {
		return err
	}
	return nil
}
func (r *Control) Update(ctx context.Context, control *reader.Control) error {
	var sql = "update nlrk_reader.control set " +
		"f51=$1, vyhod=$2, sotrudnik_vyhod=$3, zakazy=$4 " +
		"where ticket_num=$5 and f51=$6 and vizit_date>=$7"
	_, err := r.pool.Exec(ctx, sql,
		control.DepartmentF51,
		control.ExitDate,
		control.UserExit,
		control.OrdersCount,
		control.TicketNumber,
		"КР",
		control.ExitDate.Format("2006-01-02"))
	if err != nil {
		return err
	}
	return nil
}
