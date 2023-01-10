package accountcontroller

import (
	"errors"
	accountusecase "mini-wallet/src/usecases/account"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type requestCreateAccount struct {
	CustomerID string `form:"customer_xid" binding:"required"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Missing data for required field."
	case "lte":
		return "Should be less than " + fe.Param()
	case "gte":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}

func (i *AccountController) CreateAccount(c *gin.Context) {
	account := accountusecase.New(&accountusecase.AccountUseCase{
		DB: i.DB,
	})

	var req requestCreateAccount
	if err := c.ShouldBindWith(&req, binding.Form); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	create := account.CreateAccount(req.CustomerID)

	if create != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":  create.Error(),
			"status": "fail",
		})
	}

	c.JSON(200, gin.H{
		"message": req.CustomerID,
	})
}
