package model

import (
	"bookstore1.4/utils"
	// "fmt"
)

type Page struct {
	Books []*Book
	PageNo int64
	PageSize int64
	TotalPages int64
	TotalBooks int64
	MaxPrice float64
	MinPrice float64
	Username string
}

func (page *Page) IsFirstPage() bool {
	if page.PageNo == 1 {
		return false
	} else {
		return true
	}
}

func (page *Page) IsLastPage() bool {
	if page.PageNo == page.TotalPages {
		return false
	} else {
		return true
	}
}

func (page *Page) PrePage() int64 {
	return page.PageNo - 1
}

func (page *Page) SubPage() int64 {
	return page.PageNo + 1
}

func (page *Page) Test() bool {
	if page.MaxPrice == 0 {
		return false
	} else {
		return true
	}
}

func (page *Page) IsRoot() bool {
	var res string
	utils.DB.Model(&User{}).Where("name = ?", page.Username).Select("privilege").Find(&res)
	if res == "Y" {
		return true
	}
	return false
}