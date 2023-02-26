package model

type Order struct {
	ID string
	CreateTime string
	TotalCount int32
	TotalAmount float64
	Status int32 //0.未付款 1.已付款 2.等待发货 3.已发货 4.已收货
	UserID string
	UserName string `gorm:"-"`
}

func (order *Order) Unpaid() bool {
	if order.Status == 0 {
		return true
	}
	return false
}

func (order *Order) Paid() bool {
	if order.Status == 1 {
		return true
	}
	return false
}

func (order *Order) WaitingDelivery() bool {
	if order.Status == 2 {
		return true
	}
	return false
}

func (order *Order) Delivering() bool {
	if order.Status == 3 {
		return true
	}
	return false
}

func (order *Order) Signed() bool {
	if order.Status == 4 {
		return true
	}
	return false
}