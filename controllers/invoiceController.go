package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"restaurant-management/database"
)

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
