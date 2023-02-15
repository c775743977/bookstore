package utils

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/redis/go-redis/v9"
	"fmt"
	"context"
)

var DB *gorm.DB
var err error
var Ctx context.Context
var RDB *redis.ClusterClient 
func init() {
	DB, err = gorm.Open(mysql.Open("root:Chen@123@/bookstore"), &gorm.Config{})
	if err != nil {
		fmt.Println("connect to db error:", err)
	}
	RDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.108.165:6381", "192.168.108.165:6388", "192.168.108.165:6383", "192.168.108.165:6384", "192.168.108.165:6385", "192.168.108.165:6386"},
	})
	Ctx = context.Background()
}