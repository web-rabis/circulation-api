package d_social_status

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

type DSocialStatusRepository struct {
	pool *pgxpool.Pool
	db   *gorm.DB
}

func New(pool *pgxpool.Pool) *DSocialStatusRepository {
	return &DSocialStatusRepository{
		pool: pool,
	}
}

func (r *DSocialStatusRepository) List(ctx context.Context, paging *model.Paging) ([]*reader.DSocialStatus, error) {
	var list []*reader.DSocialStatus
	var f = strings.Join(orm.Fields(reader.DSocialStatus{}).SqlFields(""), ",")

	var sql = "select " + f + " from nlrk_reader.d_socialstatus "
	if paging != nil {
		sql = sql + paging.Sql()
	}
	result, err := r.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		item := orm.NewObjectFromResult(&reader.DSocialStatus{}, result, "", mapping.MappingObjects).(*reader.DSocialStatus)
		list = append(list, item)
	}
	return list, nil
}
