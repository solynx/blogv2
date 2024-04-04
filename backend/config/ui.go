package config

type UIQuery struct {
	Slug string
	Page   int  
	Limit int
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
