package routes

import accountcontroller "mini-wallet/src/interfaces/http/controllers/account"

func (i *Route) Wallet() {
	account := accountcontroller.New(&accountcontroller.AccountController{
		DB: i.DB,
	})
	i.Gin.POST("/init", account.CreateAccount)
}
