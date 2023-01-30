package model

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
	if page.Username == "root" {
		return true
	} else {
		return false
	}
}