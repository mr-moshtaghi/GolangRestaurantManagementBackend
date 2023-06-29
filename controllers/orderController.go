package controller

import (
	"context"
	"restaurant-management/database"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
)

var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

//func OrderItemOrderCreator(order models.Order) string {
//
//}
