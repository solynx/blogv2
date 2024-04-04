package config

import "github.com/google/uuid"

type Query struct {
	Limit  int       `json:"limit"`
	Page   int       `json:"page"`
	Offset int       `json:"offset"`
	Sort   string    `json:"sort"`
	Title  string    `json:"title"`
	Id     uuid.UUID `json:"id"`
}

func (q *Query) GetLimit() int {
	if q.Limit < 1 {
		q.Limit = 12
	}
	if q.Limit > 100 {
		q.Limit = 100
	}
	return q.Limit
}

func (q *Query) GetPage() int {
	if q.Page < 1 {
		q.Page = 1
	}
	return q.Page
}

func (q *Query) GetOffset() int {
	return (q.GetPage() - 1) * q.GetLimit()
}

func (q *Query) GetSort() string {
	if q.Sort != "ASC" {
		q.Sort = "DESC"
	}
	return q.Sort
}
