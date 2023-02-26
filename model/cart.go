package model

type Cart struct {
	ID string
	Items []*CartItem `gorm:"-"`
	Amount float64
	Num int32
	// 个人认为可以直接把ID设置为userID
	UserID string
	UserName string `gorm:"-"`
}

func (c *Cart) GetNum() int32 {
	var total int32
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