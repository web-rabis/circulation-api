package ebook

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/web-rabis/db/ebook"
	"github.com/web-rabis/db/internal/adapter/database/ebook/drivers"
	"github.com/web-rabis/db/internal/adapter/database/ebook/drivers/pgsql/mapping"
	"github.com/web-rabis/db/internal/adapter/database/orm"
	"github.com/web-rabis/db/model"
	"gorm.io/gorm"
)

type Repository struct {
	pool *pgxpool.Pool
	db   *gorm.DB
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) List(ctx context.Context, paging *model.Paging) ([]*ebook.Ebook, error) {
	var ebooks []*ebook.Ebook
	var f = strings.Join(orm.Fields(ebook.Ebook{}).SqlFields("ebook"), ",")

	var sql = "select " + f + " from ebook " +
		"join catalog catalog_id on catalog_id.id=ebook.catalog_id " +
		"join directory_state state_id on state_id.id=ebook.state_id " +
		"join bibliographic_level b_level_id on b_level_id.id=ebook.b_level_id " +
		"join type_description type_descr_id on type_descr_id.id=ebook.type_descr_id "
	if paging != nil {
		sql = sql + paging.Sql()
	}

	result, err := r.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		rr := orm.NewObjectFromResult(&ebook.Ebook{}, result, "", mapping.MappingObjects).(*ebook.Ebook)
		ebooks = append(ebooks, rr)
	}
	return ebooks, nil

}
func (r *Repository) ById(ctx context.Context, id int64) (*ebook.Ebook, error) {
	var f = strings.Join(orm.Fields(ebook.Ebook{}).SqlFields("ebook"), ",")

	var sql = "select " + f + " from ebook " +
		"join catalog catalog_id on catalog_id.id=ebook.catalog_id " +
		"join directory_state state_id on state_id.id=ebook.state_id " +
		"join bibliographic_level b_level_id on b_level_id.id=ebook.b_level_id " +
		"join type_description type_descr_id on type_descr_id.id=ebook.type_descr_id " +
		"where ebook.id=$1"
	result, err := r.pool.Query(ctx, sql, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return orm.NewObjectFromResult(&ebook.Ebook{}, result, "", mapping.MappingObjects).(*ebook.Ebook), nil
	}
	return nil, drivers.ErrEbookNotExist
}
func (r *Repository) Create(ctx context.Context, ebook *ebook.Ebook) error {
	return nil
}
func (r *Repository) Update(ctx context.Context, ebook *ebook.Ebook) error {
	return nil
}
func (r *Repository) Delete(ctx context.Context, id int64) error {
	return nil
}
