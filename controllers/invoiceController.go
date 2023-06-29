package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"restaurant-management/database"
	"time"
)

var invoiceCollection *mongo.Collection = database.OpenCollection(database.Client, "invoice")

type InvoiceViewFormat struct {
	InvoiceId      string
	PaymentMethod  string
	OrderId        string
	PaymentStatus  *string
	PaymentDue     interface{}
	TableNumber    interface{}
	PaymentDueDate time.Time
	OrderDetails   interface{}
}

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
