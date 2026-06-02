package dto

import "github.com/web-rabis/circulation-api/internal/domain/model"

type AuditResponse struct {
	Result []*model.OrderAudit `json:"result"`
}
