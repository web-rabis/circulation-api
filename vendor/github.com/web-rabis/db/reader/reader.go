package reader

import (
	"time"
)

type Reader struct {
	TicketNumber       int64      `json:"ticketNumber" bson:"ticket_num"`
	RFID               string     `json:"rfid" bson:"rfid_metka"`
	Barcode            string     `json:"barcode" bson:"barcode"`
	TypeCard           *DTypeCard `json:"typeCard" bson:"typecard_id"`
	RegistrationDate   *time.Time `json:"registrationDate" bson:"registration_date"`
	ReregistrationDate *time.Time `json:"reregistrationDate" bson:"reregistration_date"`
	ExpirationDate     *time.Time `json:"expirationDate" bson:"f48"`
	LastVisitDate      *time.Time `json:"lastVisitDate" bson:"last_visit_date"`
	DateCorrection     *time.Time `json:"dateCorrection" bson:"datecorection"`
	ReadersHall        string     `json:"readersHall" bson:"readers_hall"`
	Status             string     `json:"status" bson:"status"`
	BranchId           string     `json:"branchId" bson:"codefilial"`

	Iin             string         `json:"iin" bson:"id_number"`
	Firstname       string         `json:"firstname" bson:"firstname"`
	Middlename      string         `json:"middlename" bson:"middlename"`
	Lastname        string         `json:"lastname" bson:"lastname"`
	YearBirth       int64          `json:"yearBirth" bson:"year_birth"`
	BirthDate       *time.Time     `json:"birthDate" bson:"birth_date"`
	Sex             *Dictionary    `json:"sex" bson:"sex"`
	SocialStatus    *DSocialStatus `json:"socialStatus" bson:"social_status"`
	Nationality     *Dictionary    `json:"nationality" bson:"nationality"`
	Speciality      *Dictionary    `json:"speciality" bson:"speciality"`
	Address         string         `json:"address" bson:"address"`
	Email           string         `json:"email" bson:"email"`
	Phone           string         `json:"phone" bson:"phone"`
	WorkPhone       string         `json:"workPhone" bson:"work_phone"`
	Education       *Dictionary    `json:"education" bson:"education"`
	PlaceEmployment string         `json:"placeEmployment" bson:"place_employment"`
	Position        string         `json:"position" bson:"position"`
	Institution     *Dictionary    `json:"institution" bson:"institution"`
	AcademicDegree  *Dictionary    `json:"academicDegree" bson:"academic_degree"`
	Course          string         `json:"course" bson:"course" gorm:"column:course"`
	Faculty         *Dictionary    `json:"faculty" bson:"faculty_id" gorm:"column:faculty"`
	Note            string         `json:"note" bson:"note" gorm:"column:note"`
	PriznakUdal     string         `json:"priznakUdal" bson:"priznakudal"`
	Department      string         `json:"department" bson:"f51"`
	IsEmployee      bool           `json:"isEmployee" bson:"is_employee"`
}

type ReaderEmployee struct {
	TicketNumber       int64      `json:"ticketNumber" bson:"ticket_num"`
	RFID               string     `json:"rfid" bson:"rfid_metka"`
	Barcode            string     `json:"barcode" bson:"barcode"`
	TypeCard           *DTypeCard `json:"typeCard" bson:"typecard_id"`
	RegistrationDate   *time.Time `json:"registrationDate" bson:"registration_date"`
	ReregistrationDate *time.Time `json:"reregistrationDate" bson:"reregistration_date"`
	ExpirationDate     *time.Time `json:"expirationDate" bson:"f48"`
	LastVisitDate      *time.Time `json:"lastVisitDate" bson:"last_visit_date"`
	DateCorrection     *time.Time `json:"dateCorrection" bson:"datecorection"`
	ReadersHall        string     `json:"readersHall" bson:"readers_hall"`
	Status             string     `json:"status" bson:"status"`
	BranchId           string     `json:"branchId" bson:"codefilial"`

	Iin             string         `json:"iin" bson:"id_number"`
	Firstname       string         `json:"firstname" bson:"firstname"`
	Middlename      string         `json:"middlename" bson:"middlename"`
	Lastname        string         `json:"lastname" bson:"lastname"`
	YearBirth       int64          `json:"yearBirth" bson:"year_birth"`
	BirthDate       *time.Time     `json:"birthDate" bson:"birth_date"`
	Sex             *Dictionary    `json:"sex" bson:"sex"`
	SocialStatus    *DSocialStatus `json:"socialStatus" bson:"social_status"`
	Nationality     *Dictionary    `json:"nationality" bson:"nationality"`
	Speciality      *Dictionary    `json:"speciality" bson:"speciality"`
	Address         string         `json:"address" bson:"address"`
	Email           string         `json:"email" bson:"email"`
	Phone           string         `json:"phone" bson:"phone"`
	WorkPhone       string         `json:"workPhone" bson:"work_phone"`
	Education       *Dictionary    `json:"education" bson:"education"`
	PlaceEmployment string         `json:"placeEmployment" bson:"place_employment"`
	Position        string         `json:"position" bson:"position"`
	Institution     *Dictionary    `json:"institution" bson:"institution"`
	AcademicDegree  *Dictionary    `json:"academicDegree" bson:"academic_degree"`
	Course          string         `json:"course" bson:"course" gorm:"column:course"`
	Faculty         *Dictionary    `json:"faculty" bson:"faculty_id" gorm:"column:faculty"`
	Note            string         `json:"note" bson:"note" gorm:"column:note"`
	PriznakUdal     string         `json:"priznakUdal" bson:"priznakudal"`
	Department      string         `json:"department" bson:"f51"`
}

type ReaderUser struct {
	Id               int64      `json:"id" bson:"id"`
	Status           string     `json:"status" bson:"status"`
	RegistrationDate *time.Time `json:"registrationDate" bson:"registration_date"`
	ReadersHall      string     `json:"readersHall" bson:"readers_hall"`

	Iin             string         `json:"iin" bson:"id_number"`
	Firstname       string         `json:"firstname" bson:"firstname"`
	Middlename      string         `json:"middlename" bson:"middlename"`
	Lastname        string         `json:"lastname" bson:"lastname"`
	YearBirth       int64          `json:"yearBirth" bson:"year_birth"`
	BirthDate       time.Time      `json:"birthDate" bson:"birth_date"`
	Sex             *Dictionary    `json:"sex" bson:"sex"`
	SocialStatus    *DSocialStatus `json:"socialStatus" bson:"social_status"`
	Nationality     *Dictionary    `json:"nationality" bson:"nationality"`
	Speciality      *Dictionary    `json:"speciality" bson:"speciality"`
	Address         string         `json:"address" bson:"address"`
	Email           string         `json:"email" bson:"email"`
	Phone           string         `json:"phone" bson:"phone"`
	WorkPhone       string         `json:"workPhone" bson:"work_phone"`
	Education       *Dictionary    `json:"education" bson:"education"`
	PlaceEmployment string         `json:"placeEmployment" bson:"place_employment"`
	Position        string         `json:"position" bson:"position"`
	Institution     *Dictionary    `json:"institution" bson:"institution"`
	AcademicDegree  *Dictionary    `json:"academicDegree" bson:"academic_degree"`
	Course          string         `json:"course" bson:"course" gorm:"column:course"`
	Faculty         *Dictionary    `json:"faculty" bson:"faculty_id" gorm:"column:faculty"`
	Note            string         `json:"note" bson:"note" gorm:"column:note"`
}

type Control struct {
	TicketNumber  int64      `json:"ticketNumber" bson:"ticket_num"`
	VisitDate     *time.Time `json:"visitDate" bson:"vizit_date"`
	ExitDate      *time.Time `json:"exitDate" bson:"vyhod"`
	UserVisit     string     `json:"userVisit" bson:"sotrudnik_vizit"`
	UserExit      string     `json:"userExit" bson:"sotrudnik_vyhod"`
	DepartmentF51 string     `json:"departmentF51" bson:"f51"`
	Department    string     `json:"department" bson:"otdel"`
	OrdersCount   int64      `json:"ordersCount" bson:"zakazy"`
	Count         int64      `json:"count" bson:"count"` // непонятно что, навсякий случай добавил
}
