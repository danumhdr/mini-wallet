package http

import (
	"mini-wallet/src/helpers"
	"mini-wallet/src/interfaces/http/routes"

	"github.com/gin-gonic/gin"
)

func (i *Http) Launch() {
	r := gin.Default()
	r.Use(gin.Recovery())

	v1 := r.Group("api/v1")
	routes := routes.New(&routes.Route{
		Gin: v1,
		DB:  i.DB,
	})
	routes.Ping()
	routes.Wallet()

	port := helpers.GetEnv("PORT", "5500")
	r.Run(":" + port)
}
