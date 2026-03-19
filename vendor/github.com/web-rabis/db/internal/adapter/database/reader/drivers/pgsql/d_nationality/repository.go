package d_nationality

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/web-rabis/db/internal/adapter/database/orm"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/mapping"
	"github.com/web-rabis/db/model"
	"github.com/web-rabis/db/reader"
	"gorm.io/gorm"
	"strings"
)

type DNationality struct {
	pool *pgxpool.Pool
	db   *gorm.DB
}

func New(pool *pgxpool.Pool) *DNationality {
	return &DNationality{
		pool: pool,
	}
}

func (r *DNationality) List(ctx context.Context, paging *model.Paging) ([]*reader.Dictionary, error) {
	var list []*reader.Dictionary
	var f = strings.Join(orm.Fields(reader.Dictionary{}).SqlFields(""), ",")

	var sql = "select " + f + " from nlrk_reader.d_nationality "
	if paging != nil {
		sql = sql + paging.Sql()
	}
	result, err := r.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		item := orm.NewObjectFromResult(&reader.Dictionary{}, result, "", mapping.MappingObjects).(*reader.Dictionary)
		list = append(list, item)
	}
	return list, nil
}
