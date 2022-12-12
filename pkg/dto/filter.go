package dto

import "time"

type PaginationFilter struct {
	Limit  int
	Offset int
}

type DateFilter struct {
	StartDate time.Time
	EndDate   time.Time
}
