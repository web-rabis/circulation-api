package order

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/web-rabis/circulation-api/internal/adapter/database/drivers"
	"github.com/web-rabis/circulation-api/internal/domain/model"
	"github.com/web-rabis/circulation-api/internal/domain/model/order"
	"gorm.io/gorm"
)

type Repository struct {
	pool *pgxpool.Pool
	db   *gorm.DB
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{pool: pool}
}

func (r *Repository) List(ctx context.Context, filter *model.OrderFilter, paging *model.Paging) ([]*order.Order, error) {
	sql := `select eo.id eorder_id, 
				r.ticket_num reader_ticket_number, coalesce(r.barcode, '') reader_barcode, 
				r.firstname reader_firstname, r.middlename reader_middlename, r.lastname reader_lastname,
				e.id ebook_id, c.id catalog_id, c.code catalog_code, c.name catalog_name, 
				e.author ebook_author, e.title ebook_title, coalesce(e.placement, 0) ebook_placement, coalesce(e.format, '') ebook_format,
				ei.id inv_number_id, ei.ebook_id inv_number_ebook_id, ei.inv_number inv_number_inv_number,
				ds.id state_id, ds.code state_code, ds.name state_name,
				eo.ordered_date eorder_order_date,
				p.id per_id, p.nkr per_nkr, p.title per_title, eo.nomgazjur per_num,eo.dategazjur per_year,
				dld.id department_id, dld.code department_code, dld.name department_name,
				eo.invnumber, eo.barcode
			from eorder eo
				join nlrk_reader.reader r on eo.reader_ticket_num = r.ticket_num
				left join ebook e on eo.ebook_id = e.id
				left join catalog c on e.catalog_id = c.id
				left join ebook_inv ei on eo.ebook_inv_id = ei.id
				join directory_state ds on eo.state_id = ds.id 
				left join nlrk_completion.periodika p on eo.idgazjurnumbers = p.id 
				left join directory_lib_depart dld on dld.id=eo.depart_id `

	if filter != nil {
		sql = sql + filter.Sql()
	}
	if paging != nil {
		sql = sql + paging.Sql()
	}
	println(sql)
	rows, err := r.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*order.Order, 0)

	for rows.Next() {
		values, err := rows.Values()

		if err != nil {
			return nil, err
		}

		o := order.NewOrderFromResult(values)
		orders = append(orders, o)
	}

	return orders, nil
}

func (r *Repository) Count(ctx context.Context, filter *model.OrderFilter) (int64, error) {
	sql := `select count(eo.id) 
			from eorder eo
				join nlrk_reader.reader r on eo.reader_ticket_num = r.ticket_num
				left join ebook e on eo.ebook_id = e.id
				left join catalog c on e.catalog_id = c.id
				left join ebook_inv ei on eo.ebook_inv_id = ei.id
				join directory_state ds on eo.state_id = ds.id `
	if filter != nil {
		sql = sql + filter.Sql()
	}
	println(sql)
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

func (r *Repository) ById(ctx context.Context, id int64) (*order.Order, error) {
	sql := `select eo.id eorder_id, 
				r.ticket_num reader_ticket_number, coalesce(r.barcode, '') reader_barcode, 
				r.firstname reader_firstname, r.middlename reader_middlename, r.lastname reader_lastname,
				e.id ebook_id, c.id catalog_id, c.code catalog_code, c.name catalog_name, 
				e.author ebook_author, e.title ebook_title, coalesce(e.placement, 0) ebook_placement, coalesce(e.format, '') ebook_format,
				ei.id inv_number_id, ei.ebook_id inv_number_ebook_id, ei.inv_number inv_number_inv_number,
				ds.id state_id, ds.code state_code, ds.name state_name,
				eo.ordered_date eorder_order_date,
				p.id per_id, p.nkr per_nkr, p.title per_title, eo.nomgazjur per_num,eo.dategazjur per_year,
				dld.id department_id, dld.code department_code, dld.name department_name,
				eo.invnumber, eo.barcode
			from eorder eo
				join nlrk_reader.reader r on eo.reader_ticket_num = r.ticket_num
				left join ebook e on eo.ebook_id = e.id
				left join catalog c on e.catalog_id = c.id
				left join ebook_inv ei on eo.ebook_inv_id = ei.id
				join directory_state ds on eo.state_id = ds.id 
				left join nlrk_completion.periodika p on eo.idgazjurnumbers = p.id 
				left join directory_lib_depart dld on dld.id=eo.depart_id 
			where eo.id = $1`
	rows, err := r.pool.Query(ctx, sql, id)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, drivers.ErrOrderNotExist
	}
	values, err := rows.Values()

	if err != nil {
		return nil, err
	}

	return order.NewOrderFromResult(values), nil
}
