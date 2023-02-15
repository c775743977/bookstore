package model

type Book struct {
	ID int64 `form:"bookID"`
	Title string `form:"title" binding:"required"`
	Author string `form:"author" binding:"required"`
	Price float64 `form:"price"`
	Sales int64 `form:"sales"`
	Stock int64 `form:"stock"`
	Img_path string
	Err string `gorm:"-"`
}