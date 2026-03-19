package eorder

import "time"

type EOrder struct {
	Id                 int64       `json:"id,omitempty" bson:"id"`
	EbookId            int64       `json:"ebookId,omitempty" bson:"ebook_id"`
	InvNumber          string      `json:"invNumber,omitempty" bson:"invnumber"`
	Barcode            string      `json:"barcode,omitempty" bson:"barcode"`
	ReaderTicketNumber int64       `json:"readerTicketNumber,omitempty" bson:"reader_ticket_num"`
	OrderedAt          time.Time   `json:"ordered_date" bson:"ordered_date"`
	State              *DState     `json:"state,omitempty" bson:"state_id"`
	Department         *Department `json:"department,omitempty" bson:"depart_id"`
	EbookInvId         int64       `json:"ebookInvId,omitempty" bson:"ebook_inv_id"`
	MagazineId         int64       `json:"magazineId,omitempty" bson:"idgazjurnumbers"`
	MagazineNkr        int64       `json:"magazineNkr,omitempty" bson:"gazjurnkr"`
	MagazineTitle      string      `json:"magazineTitle,omitempty" bson:"gazjurtitle"`
	MagazineNumber     string      `json:"magazineNumber,omitempty" bson:"nomgazjur"`
	MagazineYear       string      `json:"magazineYear,omitempty" bson:"dategazjur"`
}
