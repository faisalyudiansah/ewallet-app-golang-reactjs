package httpdto

import "time"

type GetTransactionListRequest struct {
	SortBy          []string   `form:"sort_by[]"`
	SortDir         []string   `form:"sort_dir[]"`
	Limit           int        `form:"limit" binding:"numeric,gte=0"`
	Page            int        `form:"page" binding:"numeric,gte=0"`
	Search          string     `form:"s"`
	FilterType      int        `form:"transaction_type" binding:"required,numeric,gte=0"`
	FilterStartDate *time.Time `form:"start_date" time_format:"2006-01-02"`
	FilterEndDate   *time.Time `form:"end_date" time_format:"2006-01-02"`
	Category        string     `form:"category"`
}
