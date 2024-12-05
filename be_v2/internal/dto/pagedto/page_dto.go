package pagedto

import "time"

type SortDto struct {
	Column   string `json:"column"`
	OrderDir string `json:"OrderDir"`
}

type PageSortDto struct {
	Search          string     `json:"s"`
	Page            int        `json:"page"`
	Limit           int        `json:"limit"`
	Sorts           []SortDto  `json:"sorts"`
	FilterType      int        `json:"transaction_type_id"`
	FilterStartDate *time.Time `json:"start_date"`
	FilterEndDate   *time.Time `json:"end_date"`
}

type PageInfoDto struct {
	Search          string     `json:"s"`
	Page            int        `json:"page"`
	Limit           int        `json:"limit"`
	Sorts           []SortDto  `json:"sorts"`
	FilterType      int        `json:"transaction_type_id"`
	FilterStartDate *time.Time `json:"start_date"`
	FilterEndDate   *time.Time `json:"end_date"`
	TotalRow        int64      `json:"total_row"`
}
