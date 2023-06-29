package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"restaurant-management/database"
)

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {

}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
