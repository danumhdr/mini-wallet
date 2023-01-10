package routes

import "github.com/gin-gonic/gin"

func (i *Route) Ping() {
	i.Gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"token":   c.GetHeader("Authorization"),
		})
	})
}
