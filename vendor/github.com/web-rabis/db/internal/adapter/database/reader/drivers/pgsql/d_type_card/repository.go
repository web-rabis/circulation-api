package d_type_card

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/web-rabis/db/internal/adapter/database/orm"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/mapping"
	"github.com/web-rabis/db/model"
	"github.com/web-rabis/db/reader"
	"gorm.io/gorm"
)

type DTypeCardRepository struct {
	pool *pgxpool.Pool
	db   *gorm.DB
}

func New(pool *pgxpool.Pool) *DTypeCardRepository {
	return &DTypeCardRepository{
		pool: pool,
	}
}

func (r *DTypeCardRepository) List(ctx context.Context, paging *model.Paging) ([]*reader.DTypeCard, error) {
	var list []*reader.DTypeCard
	var f = strings.Join(orm.Fields(reader.DTypeCard{}).SqlFields(""), ",")

	var sql = "select " + f + " from nlrk_reader.d_typeorder "
	if paging != nil {
		sql = sql + paging.Sql()
	}
	result, err := r.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		item := orm.NewObjectFromResult(&reader.DTypeCard{}, result, "", mapping.MappingObjects).(*reader.DTypeCard)
		list = append(list, item)
	}
	return list, nil
}
func (r *DTypeCardRepository) ById(ctx context.Context, id int64) (*reader.DTypeCard, error) {
	var f = strings.Join(orm.Fields(reader.DTypeCard{}).SqlFields(""), ",")

	var sql = "select " + f + " from nlrk_reader.d_typeorder where id=$1"
	result, err := r.pool.Query(ctx, sql, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return orm.NewObjectFromResult(&reader.DTypeCard{}, result, "", mapping.MappingObjects).(*reader.DTypeCard), nil
	}
	return nil, drivers.ErrTypeCardNotExist
}
func (r *DTypeCardRepository) ByCode(ctx context.Context, code string) (*reader.DTypeCard, error) {
	var f = strings.Join(orm.Fields(reader.DTypeCard{}).SqlFields(""), ",")

	var sql = "select " + f + " from nlrk_reader.d_typeorder where code=$1"
	result, err := r.pool.Query(ctx, sql, code)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return orm.NewObjectFromResult(&reader.DTypeCard{}, result, "", mapping.MappingObjects).(*reader.DTypeCard), nil
	}
	return nil, drivers.ErrTypeCardNotExist

}
