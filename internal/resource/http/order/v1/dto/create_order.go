package dto

type CreateOrderRequest struct {
	TicketNumber int64 `json:"ticketNumber"`
	EbookId      int64 `json:"ebookId"`
	DepartmentId int64 `json:"departmentId"`
	InventoryId  int64 `json:"inventoryId"`
}
