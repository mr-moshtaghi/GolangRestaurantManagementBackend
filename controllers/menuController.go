package controller

import (
	"restaurant-management/database"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var menuCollection *mongo.Collection = database.OpenCollection(database.Client, "menu")

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func inTimeSpan(start, end, check time.Time) bool {
	return start.After(time.Now()) && end.After(start)
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
