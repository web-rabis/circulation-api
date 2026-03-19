package user

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/web-rabis/db/internal/adapter/database/user/drivers/pgsql/mapping"
	"gorm.io/gorm"

	"github.com/web-rabis/db/internal/adapter/database/orm"
	"github.com/web-rabis/db/internal/adapter/database/user/drivers"
	"github.com/web-rabis/db/model"
	"github.com/web-rabis/db/user"
)

type User struct {
	pool *pgxpool.Pool
	db   *gorm.DB
}

func New(pool *pgxpool.Pool) *User {
	return &User{
		pool: pool,
	}
}

func (r *User) List(ctx context.Context, paging *model.Paging) ([]*user.User, error) {
	var users []*user.User
	var f = strings.Join(orm.Fields(user.User{}).SqlFields("users"), ",")

	var sql = "select " + f + " from public.users user " +
		" left join directory_lib_depart depart_id on depart_id.id=users.depart_id "
	if paging != nil {
		sql = sql + paging.Sql()
	}
	result, err := r.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		rr := orm.NewObjectFromResult(&user.User{}, result, "", mapping.MappingObjects).(*user.User)
		users = append(users, rr)
	}
	return users, nil
}
func (r *User) ById(ctx context.Context, id int) (*user.User, error) {
	var f = strings.Join(orm.Fields(user.User{}).SqlFields("users"), ",")
	var sql = "select " + f + " from public.users users " +
		" left join directory_lib_depart depart_id on depart_id.id=users.depart_id " +
		" where users.id=$1"
	result, err := r.pool.Query(ctx, sql, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return orm.NewObjectFromResult(&user.User{}, result, "", mapping.MappingObjects).(*user.User), nil
	}
	return nil, drivers.ErrUserNotExist

}
func (r *User) ByUsername(ctx context.Context, username string) (*user.User, error) {
	var f = strings.Join(orm.Fields(user.User{}).SqlFields("users"), ",")
	var sql = "select " + f + " from public.users users " +
		" left join directory_lib_depart depart_id on depart_id.id=users.depart_id " +
		" where lower(users.username)=$1"
	result, err := r.pool.Query(ctx, sql, strings.ToLower(username))
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return orm.NewObjectFromResult(&user.User{}, result, "", mapping.MappingObjects).(*user.User), nil
	}
	return nil, drivers.ErrUserNotExist

}
