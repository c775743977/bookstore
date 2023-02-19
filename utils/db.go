package utils

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/redis/go-redis/v9"
	"fmt"
	"context"
)

var DBrr DBRR
var WDB *gorm.DB
var RDB1 *gorm.DB
var RDB2 *gorm.DB
var err error
var Ctx context.Context
var RDB *redis.ClusterClient 
func init() {
	WDB, err = gorm.Open(mysql.Open("root:Chen@123@tcp(192.168.108.166:3307)/bookstore"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to WDB error:", err)
	}
	RDB1, err = gorm.Open(mysql.Open("root:Chen@123@tcp(192.168.108.166:3308)/bookstore"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to RDB1 error:", err)
	}
	RDB2, err = gorm.Open(mysql.Open("root:Chen@123@tcp(192.168.108.166:3308)/bookstore"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to RDB2 error:", err)
	}
	DBrr.Add(RDB1)
	DBrr.Add(RDB2)
	RDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.108.165:6381", "192.168.108.165:6382", "192.168.108.165:6383", "192.168.108.165:6384", "192.168.108.165:6385", "192.168.108.165:6386"},
	})
	Ctx = context.Background()
}