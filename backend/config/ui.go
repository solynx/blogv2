package config

import "github.com/google/uuid"

type UIQuery struct {
	Slug       string
	Page       int
	Limit      int
	CategoryId *uuid.UUID `json:"category_id"`
	UserId     *uuid.UUID `json:"user_id"`
	RecordId   uuid.UUID  `json:"record_id"`
}

const UI_LIMIT_PAGE_SIZE = 8

func (query *UIQuery) GetLimit() int {
	if query.Limit > UI_LIMIT_PAGE_SIZE {
		return UI_LIMIT_PAGE_SIZE
	}
	return query.Limit
}

func (query *UIQuery) GetPage() int {
	if query.Page < 1 {
		return 1
	}
	return query.Page
}

func (q *UIQuery) GetOffset() int {
	return (q.GetPage() - 1) * q.GetLimit()
}
