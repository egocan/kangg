package utils

import (
	"math"
)

type Paginator struct {
	CurrentPage int64 `json:"currentPage"` //当前页
	PageSize    int64 `json:"pageSize"`    //每页数量
	TotalPage   int64 `json:"totalPage"`   //总页数
	TotalCount  int64 `json:"totalCount"`  //总数量
	PageRange   []int64
}


func GenPaginator(page, limit, count int64) Paginator {
	var paginator Paginator
	paginator.TotalCount = count
	paginator.TotalPage = int64(math.Ceil(float64(count) / float64(limit)))
	paginator.PageSize = limit
	paginator.CurrentPage = page

	var pages []int64
	total := paginator.TotalPage
	switch {
	case page >=total -4 && total >9 :
		start := total - 9 +1
		pages = make([]int64, 9)
		for i, _ := range pages{
			pages[i] = start + int64(i)
		}
	case page > 5 && total > 9:
		start := page -4
		pages = make([]int64, int64(math.Min(9, float64(page+4+1))))
		for i, _ := range pages {
			pages[i] = start + int64(i)
		}
	default:
		pages = make([]int64, int64(math.Min(9, float64(total))))
		for i, _ := range pages {
			pages[i] = int64(i) + 1
		}
	}
	paginator.PageRange = pages
	return paginator
}
