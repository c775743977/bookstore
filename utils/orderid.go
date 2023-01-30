package utils

import (
	"fmt"
	"time"
	"math/rand"
)

func CreateOrderID() (orderid string) {
	t := time.Now()
	rand.Seed(time.Now().UnixNano())
	d := t.Day() + rand.Intn(99999)
	h := t.Hour() + rand.Intn(99999)
	m := t.Minute() + rand.Intn(99999)
	s := t.Second() + rand.Intn(99999)
	ns := t.Nanosecond() + rand.Intn(99999)
	return fmt.Sprint(d)+fmt.Sprint(s)+fmt.Sprint(ns)+fmt.Sprint(m)+fmt.Sprint(h)
}