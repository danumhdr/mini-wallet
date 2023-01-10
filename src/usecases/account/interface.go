package accountusecase

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountUseCase struct {
	DB *mongo.Database
}

type iAccountUseCase interface {
	CreateAccount(CustomerID string) error
}

func New(account *AccountUseCase) iAccountUseCase {
	return account
}
