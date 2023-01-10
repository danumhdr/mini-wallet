package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Route struct {
	Gin *gin.RouterGroup
	DB  *mongo.Database
}

type iRoute interface {
	Ping()
	Wallet()
}

func New(route *Route) iRoute {
	return route
}
