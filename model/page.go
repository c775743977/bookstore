package model

import (
	"bookstore1.4/utils"
	"go.mongodb.org/mongo-driver/bson"
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
	var res bson.D
	users := utils.MDB.Database("bookstore").Collection("users")
	data := users.FindOne(utils.Ctx, bson.D{{"name",page.Username},})
	data.Decode(&res)
	if res[4].Value == "Y" {
		return true
	} else {
		return false
	}
}