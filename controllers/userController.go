package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"restaurant-management/database"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

//func HashPassword(password string) string {
//
//}

//func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
//}
