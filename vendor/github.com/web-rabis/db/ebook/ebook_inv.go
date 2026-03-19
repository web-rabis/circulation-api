package ebook

type EbookInv struct {
	Id        int64  `json:"id,omitempty" bson:"id"`
	EbookId   int64  `json:"ebookId,omitempty" bson:"ebook_id"`
	InvNumber string `json:"invNumber,omitempty" bson:"inv_number"`
	Barcode   string `json:"barcode,omitempty" bson:"barcode"`
}
