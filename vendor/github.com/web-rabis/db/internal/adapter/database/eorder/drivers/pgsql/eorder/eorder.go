package eorder

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/web-rabis/db/eorder"
	"github.com/web-rabis/db/internal/adapter/database/eorder/drivers/pgsql/mapping"
	"github.com/web-rabis/db/internal/adapter/database/orm"
	"github.com/web-rabis/db/model"
	"gorm.io/gorm"
)

type EOrder struct {
	pool *pgxpool.Pool
	db   *gorm.DB
}

func New(pool *pgxpool.Pool) *EOrder {
	return &EOrder{
		pool: pool,
	}
}

func (r *EOrder) List(ctx context.Context, filter *eorder.EOrderFilter, paging *model.Paging) ([]*eorder.EOrder, error) {
	var eorders []*eorder.EOrder
	var f = strings.Join(orm.Fields(eorder.EOrder{}).SqlFields("eorder"), ",")

	var sql = "select " + f + " from eorder " +
		"join directory_lib_depart depart_id on depart_id.id=eorder.depart_id " +
		"join directory_state state_id on state_id.id=eorder.state_id "
	sqlWhere := filter.Sql()
	if sqlWhere != "" {
		sql += " where " + sqlWhere
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
		rr := orm.NewObjectFromResult(&eorder.EOrder{}, result, "", mapping.MappingObjects).(*eorder.EOrder)
		eorders = append(eorders, rr)
	}
	return eorders, nil

}
func (r *EOrder) Count(ctx context.Context, filter *eorder.EOrderFilter) (int64, error) {

	var sql = "select count(eorder.id) from eorder"
	sqlWhere := filter.Sql()
	if sqlWhere != "" {
		sql += " where " + sqlWhere
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
