package model

type CartItems struct {
	// ID int
	Book *Book
	Amount float64
	Num int
	CartID int
	IsThis bool
}

type CartItemsByBookID struct {
	// ID int
	BookID int
	Amount float64
	Num int
	CartID int
}

func (c *CartItems) GetAmount() float64 {
	return float64(c.Num) * c.Book.Price
}