package walletcontroller

import "github.com/gin-gonic/gin"

func (i *WalletController) Disable(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
