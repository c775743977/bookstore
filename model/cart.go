package model

type Cart struct {
	ID int
	Items []*CartItem `gorm:"-"`
	Amount float64
	Num int
	// 个人认为可以直接把ID设置为userID
	UserID int
	UserName string `gorm:"-"`
}

func (c *Cart) GetNum() int {
	var total int
	for _, k := range c.Items {
		total += k.Num
	}
	return total
}

func (c *Cart) GetAmount() float64 {
	var total float64
	for _, k := range c.Items {
		total += k.Amount
	}
	return total
}