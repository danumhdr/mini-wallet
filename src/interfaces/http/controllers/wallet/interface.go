package walletcontroller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type WalletController struct {
	DB *mongo.Database
}

type iWalletController interface {
	Enable(*gin.Context)
	Disable(*gin.Context)
}

func New(wallet *WalletController) iWalletController {
	return wallet
}
