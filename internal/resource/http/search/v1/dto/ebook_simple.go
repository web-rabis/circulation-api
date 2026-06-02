package dto

import (
	ebookModel "github.com/web-rabis/searcher-proxy/model/ebook"
)

type EbookSimpleSearchRequest struct {
	Query string `json:"query"`
}
type EbookSimpleSearchResponse struct {
	Result []*ebookModel.Ebook `json:"result"`
	Count  int64               `json:"count"`
}
