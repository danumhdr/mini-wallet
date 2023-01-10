package accountcontroller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountController struct {
	DB *mongo.Database
}

type iAccountController interface {
	CreateAccount(*gin.Context)
}

func New(account *AccountController) iAccountController {
	return account
}
