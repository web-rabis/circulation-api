package reader_user

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/web-rabis/db/internal/adapter/database/orm"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/mapping"
	"github.com/web-rabis/db/model"
	"github.com/web-rabis/db/reader"
	"gorm.io/gorm"
)

type ReaderUser struct {
	pool *pgxpool.Pool
	db   *gorm.DB
}

func New(pool *pgxpool.Pool) *ReaderUser {
	return &ReaderUser{
		pool: pool,
	}
}

func (r *ReaderUser) List(ctx context.Context, filter *reader.ReaderUserFilter, paging *model.Paging) ([]*reader.ReaderUser, error) {
	var readers []*reader.ReaderUser
	var f = strings.Join(orm.Fields(reader.ReaderUser{}).SqlFields("reader_user"), ",")

	var sql = "select " + f + " from nlrk_reader.reader_user" +
		" left join nlrk_reader.d_sex sex on sex.id=reader_user.sex" +
		" left join nlrk_reader.d_socialstatus social_status on social_status.id=reader_user.social_status" +
		" left join nlrk_reader.d_nationality nationality on nationality.id=reader_user.nationality" +
		" left join nlrk_reader.d_speciality speciality on speciality.id=reader_user.speciality" +
		" left join nlrk_reader.d_education education on education.id=reader_user.education" +
		" left join nlrk_reader.d_institution institution on institution.id=reader_user.institution" +
		" left join nlrk_reader.d_academic_degree academic_degree on academic_degree.id=reader_user.academic_degree" +
		" left join nlrk_reader.d_faculty faculty_id on faculty_id.id=reader_user.faculty_id"
	if filter != nil {
		sql += " where " + filter.Sql()
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
		rr := orm.NewObjectFromResult(&reader.ReaderUser{}, result, "", mapping.MappingObjects).(*reader.ReaderUser)
		readers = append(readers, rr)
	}
	return readers, nil
}
func (r *ReaderUser) Count(ctx context.Context, filter *reader.ReaderUserFilter) (int64, error) {

	var sql = "select count(reader_user.id) from nlrk_reader.reader_user" +
		" left join nlrk_reader.d_sex sex on sex.id=reader_user.sex" +
		" left join nlrk_reader.d_socialstatus social_status on social_status.id=reader_user.social_status" +
		" left join nlrk_reader.d_nationality nationality on nationality.id=reader_user.nationality" +
		" left join nlrk_reader.d_speciality speciality on speciality.id=reader_user.speciality" +
		" left join nlrk_reader.d_education education on education.id=reader_user.education" +
		" left join nlrk_reader.d_institution institution on institution.id=reader_user.institution" +
		" left join nlrk_reader.d_academic_degree academic_degree on academic_degree.id=reader_user.academic_degree" +
		" left join nlrk_reader.d_faculty faculty_id on faculty_id.id=reader_user.faculty_id"
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

func (r *ReaderUser) ById(ctx context.Context, id int) (*reader.ReaderUser, error) {
	var f = strings.Join(orm.Fields(reader.ReaderUser{}).SqlFields("reader_user"), ",")
	var sql = "select " + f + " from nlrk_reader.reader_user" +
		" left join nlrk_reader.d_sex sex on sex.id=reader_user.sex" +
		" left join nlrk_reader.d_socialstatus social_status on social_status.id=reader_user.social_status" +
		" left join nlrk_reader.d_nationality nationality on nationality.id=reader_user.nationality" +
		" left join nlrk_reader.d_speciality speciality on speciality.id=reader_user.speciality" +
		" left join nlrk_reader.d_education education on education.id=reader_user.education" +
		" left join nlrk_reader.d_institution institution on institution.id=reader_user.institution" +
		" left join nlrk_reader.d_academic_degree academic_degree on academic_degree.id=reader_user.academic_degree" +
		" left join nlrk_reader.d_faculty faculty_id on faculty_id.id=reader_user.faculty_id" +
		" where reader_user.id=$1"
	result, err := r.pool.Query(ctx, sql, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return orm.NewObjectFromResult(&reader.ReaderUser{}, result, "", mapping.MappingObjects).(*reader.ReaderUser), nil
	}
	return nil, drivers.ErrReaderUserNotExist

}
func (r *ReaderUser) ByIin(ctx context.Context, iin string) (*reader.ReaderUser, error) {
	var f = strings.Join(orm.Fields(reader.ReaderUser{}).SqlFields("reader_user"), ",")
	var sql = "select " + f + " from nlrk_reader.reader_user" +
		" left join nlrk_reader.d_sex sex on sex.id=reader_user.sex" +
		" left join nlrk_reader.d_socialstatus social_status on social_status.id=reader_user.social_status" +
		" left join nlrk_reader.d_nationality nationality on nationality.id=reader_user.nationality" +
		" left join nlrk_reader.d_speciality speciality on speciality.id=reader_user.speciality" +
		" left join nlrk_reader.d_education education on education.id=reader_user.education" +
		" left join nlrk_reader.d_institution institution on institution.id=reader_user.institution" +
		" left join nlrk_reader.d_academic_degree academic_degree on academic_degree.id=reader_user.academic_degree" +
		" left join nlrk_reader.d_faculty faculty_id on faculty_id.id=reader_user.faculty_id" +
		" where reader_user.id_number=$1"
	result, err := r.pool.Query(ctx, sql, iin)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return orm.NewObjectFromResult(&reader.ReaderUser{}, result, "", mapping.MappingObjects).(*reader.ReaderUser), nil
	}
	return nil, drivers.ErrReaderUserNotExist

}
func (r *ReaderUser) Create(ctx context.Context, reader_ *reader.ReaderUser) (int64, error) {
	var err error
	if reader_.Id == 0 {
		reader_.Id, err = r.nextId(ctx)
		if err != nil {
			return 0, err
		}
	}
	var rFields = orm.Fields(reader.ReaderUser{})
	var rf = reflect.ValueOf(reader_).Elem()
	var sql = "insert into nlrk_reader.reader_user(%s) values(%s)"
	fields, values := rFields.FieldsValues(rf)
	var vv []string
	for i, _ := range values {
		vv = append(vv, "$"+strconv.Itoa(i+1))
	}
	sql = fmt.Sprintf(sql, strings.Join(fields, ","), strings.Join(vv, ","))
	_, err = r.pool.Exec(ctx, sql, values...)
	if err != nil {
		return 0, err
	}
	return reader_.Id, nil
}

func (r *ReaderUser) nextId(ctx context.Context) (int64, error) {
	var sql = "SELECT nextval('nlrk_reader.reader_user_id_seq') nextval"
	var id int64

	result, err := r.pool.Query(ctx, sql)
	if err != nil {
		return 0, err
	}
	defer result.Close()

	if result.Next() {
		err = result.Scan(&id)
		if err != nil {
			return 0, err
		}
	} else {
		return 0, errors.New("no nextval")
	}
	return id, nil
}

func (r *ReaderUser) Update(ctx context.Context, reader_ *reader.ReaderUser) error {
	var rFields = orm.Fields(reader.ReaderUser{})
	var rf = reflect.ValueOf(reader_).Elem()
	var sql = "update nlrk_reader.reader_user set %s where id=%s"
	fields, values := rFields.FieldsValues(rf)
	var vv []string
	for i, field := range fields {
		vv = append(vv, field+"=$"+strconv.Itoa(i+1))
	}
	sql = fmt.Sprintf(sql, strings.Join(vv, ","), strconv.Itoa(int(reader_.Id)))
	_, err := r.pool.Exec(ctx, sql, values...)
	if err != nil {
		return err
	}
	return nil
}
func (r *ReaderUser) UpdateStatus(ctx context.Context, id int, status string) error {
	var sql = "update nlrk_reader.reader_user set status=$1 where id=$2"
	_, err := r.pool.Exec(ctx, sql, status, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReaderUser) UploadPhoto(ctx context.Context, id int, photo []byte) error {
	var sql = "select id from nlrk_reader.reader_user_photo where id=$1"
	result, err := r.pool.Query(ctx, sql, id)
	if err != nil {
		return err
	}
	defer result.Close()
	sql = "insert into nlrk_reader.reader_user_photo(photo,id) values($1,$2)"
	if result.Next() {
		sql = "update nlrk_reader.reader_user_photo set photo=$1 where id=$2"
	}
	_, err = r.pool.Exec(ctx, sql, photo, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *ReaderUser) PhotoById(ctx context.Context, id int) ([]byte, error) {
	var sql = "select photo from nlrk_reader.reader_user_photo where id=$1"
	result, err := r.pool.Query(ctx, sql, id)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return result.RawValues()[0], nil
	}
	return nil, nil
}
func (r *ReaderUser) Delete(ctx context.Context, id int64) error {
	var sql = "delete from nlrk_reader.reader_user where id=$1"
	_, err := r.pool.Exec(ctx, sql, id)
	if err != nil {
		return err
	}
	return nil

}
