package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"restaurant-management/database"
	"restaurant-management/models"
	"time"
)

type OrderItemPack struct {
	TableId    *string
	OrderItems []models.OrderItem
}

var orderItemCollection *mongo.Collection = database.OpenCollection(database.Client, "orderItem")

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		result, err := orderItemCollection.Find(context.TODO(), bson.M{})

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing ordered items"})
			return
		}
		var allOrderItems []bson.M
		if err = result.All(ctx, &allOrderItems); err != nil {
			log.Fatal(err)
			return
		}
		c.JSON(http.StatusOK, allOrderItems)
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
