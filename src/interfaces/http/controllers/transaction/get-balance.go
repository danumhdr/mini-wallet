package transactioncontroller

import "github.com/gin-gonic/gin"

func (i *TransactionController) GetBalance(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
