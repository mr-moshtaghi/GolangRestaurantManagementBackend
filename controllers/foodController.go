package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"math"
	"restaurant-management/database"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}


