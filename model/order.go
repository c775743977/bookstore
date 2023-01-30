package model


type Order struct {
	OrderID string
	CreateTime string
	Num int64
	Amount float64
	Status int64 //0.未付款 1.已付款 2.等待发货 3.已发货 4.已收货
	UserID int64
	UserName string
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