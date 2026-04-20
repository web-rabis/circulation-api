package dto

type RedirectOrderRequest struct {
	Ids          []int64 `json:"ids"`
	DepartmentId int64   `json:"departmentId"`
}
