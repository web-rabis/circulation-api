package drivers

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/web-rabis/db/model"
	"github.com/web-rabis/db/reader"
)

type DataStore interface {
	Base
}

type Base interface {
	Name() string
	Ping(ctx context.Context) error
	Close(ctx context.Context) error
	Connect(ctx context.Context) error

	Reader() Reader
	ReaderUser() ReaderUser
	ReaderEmployee() ReaderEmployee
	DTypeCard() DTypeCard
	DSex() DSex
	DSocialStatus() DSocialStatus
	DNationality() DNationality
	DSpeciality() DSpeciality
	DEducation() DEducation
	DInstitution() DInstitution
	DAcademicDegree() DAcademicDegree
	DFaculty() DFaculty
	Control() Control
}
type Reader interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)
	List(ctx context.Context, filter *reader.ReaderFilter, paging *model.Paging) ([]*reader.Reader, error)
	Count(ctx context.Context, filter *reader.ReaderFilter) (int64, error)
	ByTicketNumber(ctx context.Context, ticketNumber int64) (*reader.Reader, error)
	ByBarcode(ctx context.Context, barcode string) (*reader.Reader, error)
	ByIin(ctx context.Context, iin string) (*reader.Reader, error)
	Create(ctx context.Context, reader *reader.Reader) error
	CreateTx(ctx context.Context, tx pgx.Tx, reader_ *reader.Reader) error
	Update(ctx context.Context, reader *reader.Reader) error
	Delete(ctx context.Context, ticketNumber int64) error
	NextTicketNumber(ctx context.Context) (int64, error)
	NextTemporaryTicketNumber(ctx context.Context) (int64, error)
	UploadPhoto(ctx context.Context, ticketNumber int64, photo []byte, user string) error
	UploadPhotoTx(ctx context.Context, tx pgx.Tx, ticketNumber int64, photo []byte, user string) error
	PhotoById(ctx context.Context, ticketNumber int64) ([]byte, error)
	UpdateStatus(ctx context.Context, ticketNumber int64, status string) error
	UpdateTicketNumber(ctx context.Context, oldTicketNumber, newTicketNumber int64) error
}
type ReaderEmployee interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)
	List(ctx context.Context, filter *reader.ReaderEmployeeFilter, paging *model.Paging) ([]*reader.ReaderEmployee, error)
	Count(ctx context.Context, filter *reader.ReaderEmployeeFilter) (int64, error)
	ByTicketNumber(ctx context.Context, ticketNumber int64) (*reader.ReaderEmployee, error)
	ByIin(ctx context.Context, iin string) (*reader.ReaderEmployee, error)
	Create(ctx context.Context, reader *reader.ReaderEmployee) error
	CreateTx(ctx context.Context, tx pgx.Tx, reader_ *reader.ReaderEmployee) error
	Update(ctx context.Context, reader *reader.ReaderEmployee) error
	Delete(ctx context.Context, ticketNumber int64) error
	NextTicketNumber(ctx context.Context) (int64, error)
	UploadPhoto(ctx context.Context, ticketNumber int64, photo []byte, user string) error
	UploadPhotoTx(ctx context.Context, tx pgx.Tx, ticketNumber int64, photo []byte, user string) error
	PhotoById(ctx context.Context, ticketNumber int64) ([]byte, error)
	UpdateStatus(ctx context.Context, ticketNumber int64, status string) error
}
type ReaderUser interface {
	List(ctx context.Context, filter *reader.ReaderUserFilter, paging *model.Paging) ([]*reader.ReaderUser, error)
	Count(ctx context.Context, filter *reader.ReaderUserFilter) (int64, error)
	ById(ctx context.Context, id int) (*reader.ReaderUser, error)
	ByIin(ctx context.Context, iin string) (*reader.ReaderUser, error)
	Create(ctx context.Context, reader_ *reader.ReaderUser) (int64, error)
	Update(ctx context.Context, reader *reader.ReaderUser) error
	Delete(ctx context.Context, id int64) error
	UploadPhoto(ctx context.Context, id int, photo []byte) error
	PhotoById(ctx context.Context, id int) ([]byte, error)
	UpdateStatus(ctx context.Context, id int, status string) error
}
type DTypeCard interface {
	List(ctx context.Context, paging *model.Paging) ([]*reader.DTypeCard, error)
	ById(ctx context.Context, id int64) (*reader.DTypeCard, error)
	ByCode(ctx context.Context, code string) (*reader.DTypeCard, error)
}
type DSex interface {
	List(ctx context.Context, paging *model.Paging) ([]*reader.Dictionary, error)
}
type DSocialStatus interface {
	List(ctx context.Context, paging *model.Paging) ([]*reader.DSocialStatus, error)
}
type DNationality interface {
	List(ctx context.Context, paging *model.Paging) ([]*reader.Dictionary, error)
}
type DSpeciality interface {
	List(ctx context.Context, paging *model.Paging) ([]*reader.Dictionary, error)
}
type DEducation interface {
	List(ctx context.Context, paging *model.Paging) ([]*reader.Dictionary, error)
}
type DInstitution interface {
	List(ctx context.Context, paging *model.Paging) ([]*reader.Dictionary, error)
}
type DAcademicDegree interface {
	List(ctx context.Context, paging *model.Paging) ([]*reader.Dictionary, error)
}
type DFaculty interface {
	List(ctx context.Context, paging *model.Paging) ([]*reader.Dictionary, error)
}
type Control interface {
	List(ctx context.Context, filter *reader.ControlFilter, paging *model.Paging) ([]*reader.Control, error)
	Count(ctx context.Context, filter *reader.ControlFilter) (int64, error)
	Create(ctx context.Context, control *reader.Control) error
	Update(ctx context.Context, control *reader.Control) error
}
