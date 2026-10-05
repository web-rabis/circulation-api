package dto

type IssueOrderRequest struct {
	InventoryId int64  `json:"inventoryId"`
	Barcode     string `json:"barcode"`
}
