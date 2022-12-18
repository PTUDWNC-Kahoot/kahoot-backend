package response

import (
	"encoding/json"
	"math"
)

type Collection struct {
	Results    interface{} `json:"results,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Data struct {
	Result interface{} `json:"result,omitempty"`
}

type Pagination struct {
	TotalRecords int64 `json:"totalRecords"`
	TotalPages   int64 `json:"totalPages"`
	PageSize     int64 `json:"pageSize"`
	CurrentPage  int64 `json:"currentPage"`
}

func (p *Pagination) MarshalJSON() ([]byte, error) {
	if p.TotalPages == 0 {
		p.TotalPages = int64(math.Ceil(float64(p.TotalRecords) / float64(p.PageSize)))
	}

	return json.Marshal(p)
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	TraceID string      `json:"traceId"`
}

type FailureResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	TraceID string `json:"traceId"`
}
