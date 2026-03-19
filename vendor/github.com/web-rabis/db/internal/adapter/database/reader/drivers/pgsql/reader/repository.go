package reader

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/web-rabis/db/internal/adapter/database/orm"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/mapping"
	"github.com/web-rabis/db/model"
	"github.com/web-rabis/db/reader"
	"gorm.io/gorm"
)

type Reader struct {
	pool *pgxpool.Pool
	db   *gorm.DB
}

func New(pool *pgxpool.Pool) *Reader {
	return &Reader{
		pool: pool,
	}
}
func (r *Reader) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return r.pool.Begin(ctx)
}
func (r *Reader) List(ctx context.Context, filter *reader.ReaderFilter, paging *model.Paging) ([]*reader.Reader, error) {
	var readers []*reader.Reader
	var f = strings.Join(orm.Fields(reader.Reader{}).SqlFields("reader"), ",")

	var sql = "select " + f + " from nlrk_reader.reader" +
		" left join nlrk_reader.d_typeorder typecard_id on typecard_id.id=reader.typecard_id" +
		" left join nlrk_reader.d_sex sex on sex.id=reader.sex" +
		" left join nlrk_reader.d_socialstatus social_status on social_status.id=reader.social_status" +
		" left join nlrk_reader.d_nationality nationality on nationality.id=reader.nationality" +
		" left join nlrk_reader.d_speciality speciality on speciality.id=reader.speciality" +
		" left join nlrk_reader.d_education education on education.id=reader.education" +
		" left join nlrk_reader.d_institution institution on institution.id=reader.institution" +
		" left join nlrk_reader.d_academic_degree academic_degree on academic_degree.id=reader.academic_degree" +
		" left join nlrk_reader.d_faculty faculty_id on faculty_id.id=reader.faculty_id"
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
		rr := orm.NewObjectFromResult(&reader.Reader{}, result, "", mapping.MappingObjects).(*reader.Reader)
		readers = append(readers, rr)
	}
	return readers, nil
}
func (r *Reader) Count(ctx context.Context, filter *reader.ReaderFilter) (int64, error) {

	var sql = "select count(reader.ticket_num) from nlrk_reader.reader" +
		" left join nlrk_reader.d_typeorder typecard_id on typecard_id.id=reader.typecard_id" +
		" left join nlrk_reader.d_sex sex on sex.id=reader.sex" +
		" left join nlrk_reader.d_socialstatus social_status on social_status.id=reader.social_status" +
		" left join nlrk_reader.d_nationality nationality on nationality.id=reader.nationality" +
		" left join nlrk_reader.d_speciality speciality on speciality.id=reader.speciality" +
		" left join nlrk_reader.d_education education on education.id=reader.education" +
		" left join nlrk_reader.d_institution institution on institution.id=reader.institution" +
		" left join nlrk_reader.d_academic_degree academic_degree on academic_degree.id=reader.academic_degree" +
		" left join nlrk_reader.d_faculty faculty_id on faculty_id.id=reader.faculty_id"
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
func (r *Reader) ByTicketNumber(ctx context.Context, ticketNumber int64) (*reader.Reader, error) {
	var f = strings.Join(orm.Fields(reader.Reader{}).SqlFields("reader"), ",")
	var sql = "select " + f + " from nlrk_reader.reader" +
		" left join nlrk_reader.d_typeorder typecard_id on typecard_id.id=reader.typecard_id" +
		" left join nlrk_reader.d_sex sex on sex.id=reader.sex" +
		" left join nlrk_reader.d_socialstatus social_status on social_status.id=reader.social_status" +
		" left join nlrk_reader.d_nationality nationality on nationality.id=reader.nationality" +
		" left join nlrk_reader.d_speciality speciality on speciality.id=reader.speciality" +
		" left join nlrk_reader.d_education education on education.id=reader.education" +
		" left join nlrk_reader.d_institution institution on institution.id=reader.institution" +
		" left join nlrk_reader.d_academic_degree academic_degree on academic_degree.id=reader.academic_degree" +
		" left join nlrk_reader.d_faculty faculty_id on faculty_id.id=reader.faculty_id" +
		" where ticket_num=$1"
	result, err := r.pool.Query(ctx, sql, ticketNumber)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return orm.NewObjectFromResult(&reader.Reader{}, result, "", mapping.MappingObjects).(*reader.Reader), nil
	}
	return nil, drivers.ErrReaderNotExist

}
func (r *Reader) ByBarcode(ctx context.Context, barcode string) (*reader.Reader, error) {
	var f = strings.Join(orm.Fields(reader.Reader{}).SqlFields("reader"), ",")
	var sql = "select " + f + " from nlrk_reader.reader" +
		" left join nlrk_reader.d_typeorder typecard_id on typecard_id.id=reader.typecard_id" +
		" left join nlrk_reader.d_sex sex on sex.id=reader.sex" +
		" left join nlrk_reader.d_socialstatus social_status on social_status.id=reader.social_status" +
		" left join nlrk_reader.d_nationality nationality on nationality.id=reader.nationality" +
		" left join nlrk_reader.d_speciality speciality on speciality.id=reader.speciality" +
		" left join nlrk_reader.d_education education on education.id=reader.education" +
		" left join nlrk_reader.d_institution institution on institution.id=reader.institution" +
		" left join nlrk_reader.d_academic_degree academic_degree on academic_degree.id=reader.academic_degree" +
		" left join nlrk_reader.d_faculty faculty_id on faculty_id.id=reader.faculty_id" +
		" where barcode=$1"
	result, err := r.pool.Query(ctx, sql, barcode)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return orm.NewObjectFromResult(&reader.Reader{}, result, "", mapping.MappingObjects).(*reader.Reader), nil
	}
	return nil, drivers.ErrReaderNotExist

}

func (r *Reader) ByIin(ctx context.Context, iin string) (*reader.Reader, error) {
	var f = strings.Join(orm.Fields(reader.Reader{}).SqlFields("reader"), ",")
	var sql = "select " + f + " from nlrk_reader.reader" +
		" left join nlrk_reader.d_typeorder typecard_id on typecard_id.id=reader.typecard_id" +
		" left join nlrk_reader.d_sex sex on sex.id=reader.sex" +
		" left join nlrk_reader.d_socialstatus social_status on social_status.id=reader.social_status" +
		" left join nlrk_reader.d_nationality nationality on nationality.id=reader.nationality" +
		" left join nlrk_reader.d_speciality speciality on speciality.id=reader.speciality" +
		" left join nlrk_reader.d_education education on education.id=reader.education" +
		" left join nlrk_reader.d_institution institution on institution.id=reader.institution" +
		" left join nlrk_reader.d_academic_degree academic_degree on academic_degree.id=reader.academic_degree" +
		" left join nlrk_reader.d_faculty faculty_id on faculty_id.id=reader.faculty_id" +
		" where reader.id_number=$1"
	result, err := r.pool.Query(ctx, sql, iin)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return orm.NewObjectFromResult(&reader.Reader{}, result, "", mapping.MappingObjects).(*reader.Reader), nil
	}
	return nil, drivers.ErrReaderNotExist

}
func (r *Reader) Create(ctx context.Context, reader_ *reader.Reader) error {
	return r.CreateTx(ctx, nil, reader_)
}
func (r *Reader) CreateTx(ctx context.Context, tx pgx.Tx, reader_ *reader.Reader) error {
	var rFields = orm.Fields(reader.Reader{})
	var rf = reflect.ValueOf(reader_).Elem()
	var sql = "insert into nlrk_reader.reader(%s) values(%s)"
	fields, values := rFields.FieldsValues(rf)
	var vv []string
	for i, _ := range values {
		vv = append(vv, "$"+strconv.Itoa(i+1))
	}
	sql = fmt.Sprintf(sql, strings.Join(fields, ","), strings.Join(vv, ","))
	var err error
	if tx != nil {
		_, err = tx.Exec(ctx, sql, values...)
	} else {
		_, err = r.pool.Exec(ctx, sql, values...)
	}
	if err != nil {
		return err
	}
	return nil
}
func (r *Reader) Update(ctx context.Context, reader_ *reader.Reader) error {
	var rFields = orm.Fields(reader.Reader{})
	var rf = reflect.ValueOf(reader_).Elem()
	var sql = "update nlrk_reader.reader set %s where ticket_num=%v"
	fields, values := rFields.FieldsValues(rf)

	var vv []string
	for i, field := range fields {
		vv = append(vv, field+"=$"+strconv.Itoa(i+1))
	}
	sql = fmt.Sprintf(sql, strings.Join(vv, ","), reader_.TicketNumber)
	_, err := r.pool.Exec(ctx, sql, values...)
	if err != nil {
		return err
	}
	return nil
}
func (r *Reader) Delete(ctx context.Context, ticketNumber int64) error {
	var sql = "delete from nlrk_reader.reader where ticket_num=$1"
	_, err := r.pool.Exec(ctx, sql, ticketNumber)
	if err != nil {
		return err
	}
	return nil

}
func (r *Reader) NextTicketNumber(ctx context.Context) (int64, error) {
	return r.nextSequence(ctx, "nlrk_reader.reader_id_seq")
}
func (r *Reader) NextTemporaryTicketNumber(ctx context.Context) (int64, error) {
	return r.nextSequence(ctx, "nlrk_reader.reader_id_temp_seq")
}
func (r *Reader) nextSequence(ctx context.Context, seqName string) (int64, error) {
	var sql = "select nextval('" + seqName + "') as Ticket_Num_seq"
	result, err := r.pool.Query(ctx, sql)
	if err != nil {
		return 0, err
	}
	defer result.Close()
	var ticketNum int64
	if result.Next() {
		err = result.Scan(&ticketNum)
		if err != nil {
			return 0, err
		}
		return ticketNum, nil
	}
	return 0, drivers.ErrReaderNotExist
}

func (r *Reader) UploadPhoto(ctx context.Context, ticketNumber int64, photo []byte, user string) error {
	return r.UploadPhotoTx(ctx, nil, ticketNumber, photo, user)
}
func (r *Reader) UploadPhotoTx(ctx context.Context, tx pgx.Tx, ticketNumber int64, photo []byte, user string) error {
	var sql = "select ticket_num from nlrk_reader.reader_photo where ticket_num=$1"
	result, err := r.pool.Query(ctx, sql, ticketNumber)
	if err != nil {
		return err
	}
	defer result.Close()
	sql = "insert into nlrk_reader.reader_photo(photo,sotrudnik,datecorection,ticket_num) values($1,$2,now(),$3)"
	if result.Next() {
		sql = "update nlrk_reader.reader_photo set photo=$1,sotrudnik=$2,datecorection=now() where ticket_num=$3"
	}
	if tx != nil {
		_, err = tx.Exec(ctx, sql, photo, user, ticketNumber)
	} else {
		_, err = r.pool.Exec(ctx, sql, photo, user, ticketNumber)
	}
	if err != nil {
		return err
	}
	return nil
}
func (r *Reader) PhotoById(ctx context.Context, ticketNumber int64) ([]byte, error) {
	var sql = "select photo from nlrk_reader.reader_photo where ticket_num=$1"
	result, err := r.pool.Query(ctx, sql, ticketNumber)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		return result.RawValues()[0], nil
	}
	return nil, nil
}
func (r *Reader) UpdateStatus(ctx context.Context, ticketNumber int64, status string) error {
	var sql = "update nlrk_reader.reader set status=$1 where ticket_num=$2"
	_, err := r.pool.Exec(ctx, sql, status, ticketNumber)
	if err != nil {
		return err
	}
	return nil
}
func (r *Reader) UpdateTicketNumber(ctx context.Context, oldTicketNumber, newTicketNumber int64) error {
	var sql = "update nlrk_reader.reader set ticket_num=$1 where ticket_num=$2"
	_, err := r.pool.Exec(ctx, sql, newTicketNumber, oldTicketNumber)
	if err != nil {
		return err
	}
	sql = "update nlrk_reader.reader_photo set ticket_num=$1 where ticket_num=$2"
	_, err = r.pool.Exec(ctx, sql, newTicketNumber, oldTicketNumber)
	if err != nil {
		return err
	}
	return nil
}
