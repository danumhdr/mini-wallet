package transactioncontroller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionController struct {
	DB *mongo.Database
}

type iTransactionController interface {
	GetTransaction(*gin.Context)
	AddDeposit(*gin.Context)
	UseDeposit(*gin.Context)
	GetBalance(*gin.Context)
}

func New(transaction *TransactionController) iTransactionController {
	return transaction
}
