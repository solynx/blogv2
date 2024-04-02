package config

type Pagination struct {
	Page int `json:"page"`
	Total int64 `json:"total"`
}