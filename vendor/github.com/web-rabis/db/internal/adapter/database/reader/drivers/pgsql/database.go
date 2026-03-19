package pgsql

import (
	"context"
	"time"

	drivers2 "github.com/web-rabis/db/internal/adapter/database/reader/drivers"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/control"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/d_academic_degree"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/d_education"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/d_faculty"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/d_institution"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/d_nationality"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/d_sex"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/d_social_status"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/d_speciality"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/d_type_card"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/reader"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/reader_employee"
	"github.com/web-rabis/db/internal/adapter/database/reader/drivers/pgsql/reader_user"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	connectionTimeout = 100 * time.Second
	ensureIdxTimeout  = 10 * time.Second
)

type PgSql struct {
	connURL string
	dbName  string

	client *pgconn.PgConn
	pool   *pgxpool.Pool
	config *pgxpool.Config

	reader            drivers2.Reader
	readerUser        drivers2.ReaderUser
	readerEmployee    drivers2.ReaderEmployee
	dTypeCard         drivers2.DTypeCard
	dSex              drivers2.DSex
	dSocialStatus     drivers2.DSocialStatus
	dNationality      drivers2.DNationality
	dSpeciality       drivers2.DSpeciality
	dEducation        drivers2.DEducation
	dInstitution      drivers2.DInstitution
	dFaculty          drivers2.DFaculty
	dAcademicDegree   drivers2.DAcademicDegree
	control           drivers2.Control
	connectionTimeout time.Duration
	ensureIdxTimeout  time.Duration
}

func New(conf drivers2.DataStoreConfig) (drivers2.DataStore, error) {
	if conf.URL == "" {
		return nil, drivers2.ErrInvalidConfigStruct
	}

	if conf.DataBaseName == "" {
		return nil, drivers2.ErrInvalidConfigStruct
	}

	config, err := pgxpool.ParseConfig(conf.URL)
	if err != nil {
		return nil, err
	}
	config.MaxConns = 100
	return &PgSql{
		connURL:           conf.URL,
		dbName:            conf.DataBaseName,
		config:            config,
		connectionTimeout: connectionTimeout,
		ensureIdxTimeout:  ensureIdxTimeout,
	}, nil
}

func (m *PgSql) Name() string { return "PgSql" }

func (m *PgSql) Connect(ctx context.Context) error {
	ctxWT, cancel := context.WithTimeout(ctx, m.connectionTimeout)
	defer cancel()

	var err error
	m.pool, err = pgxpool.NewWithConfig(ctxWT, m.config)
	if err != nil {
		return err
	}
	return nil
}

func (m *PgSql) Ping(ctx context.Context) error {
	conn, err := m.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	return conn.Ping(ctx)
}

func (m *PgSql) Close(ctx context.Context) error {
	m.pool.Close()
	return nil
}

func (m *PgSql) ReaderUser() drivers2.ReaderUser {
	if m.readerUser == nil {
		m.readerUser = reader_user.New(m.pool)
	}

	return m.readerUser
}

func (m *PgSql) Reader() drivers2.Reader {
	if m.reader == nil {
		m.reader = reader.New(m.pool)
	}

	return m.reader
}

func (m *PgSql) ReaderEmployee() drivers2.ReaderEmployee {
	if m.readerEmployee == nil {
		m.readerEmployee = reader_employee.New(m.pool)
	}

	return m.readerEmployee
}

func (m *PgSql) DTypeCard() drivers2.DTypeCard {
	if m.dTypeCard == nil {
		m.dTypeCard = d_type_card.New(m.pool)
	}

	return m.dTypeCard
}

func (m *PgSql) DSex() drivers2.DSex {
	if m.dSex == nil {
		m.dSex = d_sex.New(m.pool)
	}

	return m.dSex
}

func (m *PgSql) DSocialStatus() drivers2.DSocialStatus {
	if m.dSocialStatus == nil {
		m.dSocialStatus = d_social_status.New(m.pool)
	}

	return m.dSocialStatus
}
func (m *PgSql) DNationality() drivers2.DNationality {
	if m.dNationality == nil {
		m.dNationality = d_nationality.New(m.pool)
	}

	return m.dNationality
}
func (m *PgSql) DSpeciality() drivers2.DSpeciality {
	if m.dSpeciality == nil {
		m.dSpeciality = d_speciality.New(m.pool)
	}

	return m.dSpeciality
}
func (m *PgSql) DEducation() drivers2.DEducation {
	if m.dEducation == nil {
		m.dEducation = d_education.New(m.pool)
	}

	return m.dEducation
}
func (m *PgSql) DInstitution() drivers2.DInstitution {
	if m.dInstitution == nil {
		m.dInstitution = d_institution.New(m.pool)
	}

	return m.dInstitution
}
func (m *PgSql) DFaculty() drivers2.DFaculty {
	if m.dFaculty == nil {
		m.dFaculty = d_faculty.New(m.pool)
	}

	return m.dFaculty
}
func (m *PgSql) DAcademicDegree() drivers2.DAcademicDegree {
	if m.dAcademicDegree == nil {
		m.dAcademicDegree = d_academic_degree.New(m.pool)
	}

	return m.dAcademicDegree
}
func (m *PgSql) Control() drivers2.Control {
	if m.control == nil {
		m.control = control.New(m.pool)
	}

	return m.control
}
