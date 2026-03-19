package order

import (
	"time"
)

const (
	OrderStateOrdered         = "Order.Ordered"
	OrderStateInStorage       = "Order.InStorage"
	OrderStateInReadingHall   = "Order.InReadingHall"
	OrderStateInHands         = "Order.InHands"
	OrderStatePostponed       = "Order.Postponed"
	OrderStateReturnToStorage = "Order.ReturnToStorage"
	OrderStateProcessed       = "Order.Processed"
	OrderStateRejected        = "Order.Rejected"
	OrderStateReaderReturned  = "Order.ReaderReturned"
	OrderStateDeleted         = "Order.Deleted"
)

type Order struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
	//Reader *reader.Reader `json:"reader"`
	//Ebook  *ebook.Ebook   `json:"ebook"`
	//InvNumber  *ebook.Inv                `json:"invNumber"`
	//State     *dictionary.State `json:"state"`
	OrderDate time.Time `json:"orderDate"`
	//Periodical *periodical.Periodical    `json:"periodical"`
	//Department *dictionary.LibDepartment `json:"department"`
}

func NewOrderFromResult(v []any) *Order {
	e := &Order{}
	e.Id = int(v[0].(int64))
	//e.Reader = &reader.Reader{
	//	TicketNumber: int(v[1].(int32)),
	//	Barcode:      v[2].(string),
	//	Lastname:     v[3].(string),
	//	Firstname:    v[4].(string),
	//	Middlename:   v[5].(string),
	//}
	if v[6] != nil {
		e.Type = "ebook"
		//e.Ebook = &ebook.Ebook{
		//	Id: int64(v[6].(int32)),
		//}
		//e.Ebook.Catalog = model.Catalog{
		//	Id:   int64(v[7].(int32)),
		//	Code: v[8].(string),
		//	Name: v[9].(string),
		//}
		//e.Ebook.Author = v[10].(string)
		//e.Ebook.Title = v[11].(string)
		//placement := int(v[12].(int32))
		//e.Ebook.Placement = &ebook.Placement{
		//	Placement: &placement,
		//}
		//e.Ebook.Format = &ebook.Format{
		//	Format: v[13].(string),
		//}
		//if v[14] != nil {
		//	e.InvNumber = &ebook.Inv{
		//		Id: int64(v[14].(int32)),
		//	}
		//	e.InvNumber.EbookId = int64(v[15].(int32))
		//	if v[16] != nil {
		//		e.InvNumber.InvNumber = v[16].(string)
		//	}
		//}
	}
	if v[17] != nil {
		//e.State = &dictionary.State{
		//	Id:   v[17].(int64),
		//	Code: v[18].(string),
		//	Name: v[19].(string),
		//}
	}
	if v[20] != nil {
		e.OrderDate = v[20].(time.Time)
	}
	if v[21] != nil {
		e.Type = "periodical"
		//e.Periodical = &periodical.Periodical{
		//	Id: v[21].(pgtype.Numeric).Int.Int64(),
		//}
		//e.Periodical.Nkr = v[22].(pgtype.Numeric).Int.Int64()
		//e.Periodical.Title = v[23].(string)
		//e.Periodical.Number = v[24].(string)
		//e.Periodical.YearEdition = v[25].(string)
	}
	//if v[26] != nil {
	//	e.Department = &dictionary.LibDepartment{
	//		Id: int64(v[26].(int32)),
	//	}
	//	e.Department.Code = v[27].(string)
	//	e.Department.Name = v[28].(string)
	//}
	//if e.InvNumber == nil && v[29] != nil {
	//	e.InvNumber = &ebook.Inv{
	//		Id:        -1,
	//		InvNumber: v[29].(string),
	//	}
	//	if v[30] != nil {
	//		e.InvNumber.Barcode = v[30].(string)
	//	}
	//}
	return e
}
