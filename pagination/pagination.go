package pagination

import (
	"math"

	"gorm.io/gorm"
)

type Param struct {
	DB      *gorm.DB
	Limit   int
	Offset  int
	OrderBy []string
	ShowSQL bool
}

type Paginator struct {
	TotalRecord int64       `json:"total_record"`
	TotalPage   int         `json:"total_page"`
	Records     interface{} `json:"records"`
	Offset      int         `json:"offset"`
	Limit       int         `json:"limit"`
	Page        int         `json:"page"`
	PrevPage    int         `json:"prev_page"`
	NextPage    int         `json:"next_page"`
}

func Paging(p *Param, result interface{}) *Paginator {
	db := p.DB

	if p.Limit == 0 {
		p.Limit = 10
	}
	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}

	done := make(chan bool, 1)
	var paginator Paginator
	var count int64
	var offset int
	var page int

	go countRecords(db, result, done, &count)

	offset = p.Offset

	db.Debug().Limit(p.Limit).Offset(offset).Find(result)

	<-done

	if offset == 0 {
		page = 1
	} else {
		page = 2
	}

	paginator.TotalRecord = count
	paginator.Records = result
	paginator.Page = page

	paginator.Offset = offset
	paginator.Limit = p.Limit
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	if page > 1 {
		paginator.PrevPage = page - 1
	} else {
		paginator.PrevPage = page
	}

	if page == paginator.TotalPage {
		paginator.NextPage = page
	} else {
		paginator.NextPage = page + 1
	}
	return &paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int64) {
	db.Model(anyType).Count(count)
	done <- true
}
