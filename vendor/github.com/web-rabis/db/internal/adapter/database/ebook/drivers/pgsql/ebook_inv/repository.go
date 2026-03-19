package ebook_inv

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

func (r *Repository) List(ctx context.Context, paging *model.Paging) ([]*ebook.EbookInv, error) {
	var ebookInvs []*ebook.EbookInv
	var f = strings.Join(orm.Fields(ebook.EbookInv{}).SqlFields("ebook_inv"), ",")

	var sql = "select " + f + " from ebook_inv "
	if paging != nil {
		sql = sql + paging.Sql()
	}

	result, err := r.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		rr := orm.NewObjectFromResult(&ebook.EbookInv{}, result, "", mapping.MappingObjects).(*ebook.EbookInv)
		ebookInvs = append(ebookInvs, rr)
	}
	return ebookInvs, nil

}
func (r *Repository) ById(ctx context.Context, id int64) (*ebook.EbookInv, error) {
	var f = strings.Join(orm.Fields(ebook.EbookInv{}).SqlFields("ebook_inv"), ",")

	var sql = "select " + f + " from ebook_inv where ebook_inv.id=$1"
	result, err := r.pool.Query(ctx, sql, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return orm.NewObjectFromResult(&ebook.EbookInv{}, result, "", mapping.MappingObjects).(*ebook.EbookInv), nil
	}
	return nil, drivers.ErrEbookNotExist
}
func (r *Repository) Create(ctx context.Context, ebook *ebook.EbookInv) error {
	return nil
}
func (r *Repository) Update(ctx context.Context, ebook *ebook.EbookInv) error {
	return nil
}
func (r *Repository) Delete(ctx context.Context, id int64) error {
	return nil
}
