package model

type Book struct {
	ID string `form:"bookID"`
	Title string `form:"title" binding:"required"`
	Author string `form:"author" binding:"required"`
	Price float64 `form:"price"`
	Sales int32 `form:"sales"`
	Stock int32 `form:"stock"`
	Img_path string
	Err string `gorm:"-"`
}