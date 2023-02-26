package model

type CartItem struct {
	// ID int
	Book *Book `gorm:"-"`
	Amount float64
	Num int32
	CartID string `gorm:"column:cart_id"`
	BookID string `gorm:"column:book_id"`
	IsThis bool `gorm:"-"`
}

type CartItemsByBookID struct {
	// ID int
	BookID int
	Amount float64
	Num int
	CartID int
}

func (c *CartItem) GetAmount() float64 {
	return float64(c.Num) * c.Book.Price
}