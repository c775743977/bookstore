package model

import (
	"time"
)

type Order struct {
	OrderID string
	CreateTime time.Time
	Num int64
	Amount float64
	OrderStatus int64 //0.未付款 1.已付款 2.等待发货 3.已发货 4.已收货
	UserID int64
}