package model

import "github.com/web-rabis/reader-client/protobuf"

type Reader struct {
	TicketNumber int64  `json:"ticketNumber"`
	Barcode      string `json:"barcode"`
	Firstname    string `json:"firstname"`
	Middlename   string `json:"middlename"`
	Lastname     string `json:"lastname"`
	Department   string `json:"department"`
	IsEmployee   bool   `json:"isEmployee"`
}

func NewReaderFromProto(r *protobuf.Reader) *Reader {
	if r == nil {
		return nil
	}
	return &Reader{
		TicketNumber: r.TicketNumber,
		Barcode:      r.Barcode,
		Firstname:    r.Firstname,
		Middlename:   r.Middlename,
		Lastname:     r.Lastname,
		Department:   r.Department,
		IsEmployee:   r.IsEmployee,
	}
}
