package utils

import (
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/redis/go-redis/v9"
	"fmt"
	"context"
)

var MDB *mongo.Client
var err error
var Ctx context.Context
var RDB *redis.ClusterClient 
var Cancel context.CancelFunc
var C_books *mongo.Collection
var C_cartitems *mongo.Collection
var C_carts *mongo.Collection
var C_orderitems *mongo.Collection
var C_orders *mongo.Collection
var C_users *mongo.Collection
func init() {
	Ctx, Cancel = context.WithTimeout(context.Background(), time.Second*20)
	MDB, err = mongo.Connect(Ctx, options.Client().ApplyURI("mongodb://192.168.108.170:27017"))
	if err != nil {
		fmt.Println("connect to mongodb error:", err)
		return
	}
	C_books = MDB.Database("bookstore").Collection("books")
	C_cartitems = MDB.Database("bookstore").Collection("cart_items")
	C_carts = MDB.Database("bookstore").Collection("carts")
	C_orderitems = MDB.Database("bookstore").Collection("order_items")
	C_orders = MDB.Database("bookstore").Collection("orders")
	C_users = MDB.Database("bookstore").Collection("users")
	RDB = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.108.165:6381", "192.168.108.165:6382", "192.168.108.165:6383", "192.168.108.165:6384", "192.168.108.165:6385", "192.168.108.165:6386"},
	})
	Ctx = context.Background()
}