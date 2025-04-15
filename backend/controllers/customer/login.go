package customer

import (
	models "backend/models/customer"
	"backend/service/customer"
	validators "backend/validators/customer"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func LoginHandler(service *customer.CustomerService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.LoginInput

		if err := ctx.ShouldBindJSON(&input); err != nil {
			var errorMessages []string

			if ve, ok := err.(validator.ValidationErrors); ok {
				errorMessages = validators.CreateLoginErrorMessage(ve)
			} else {
				errorMessages = append(errorMessages, "リクエストを正常に受け付けることができませんでした。")
			}

			ctx.JSON(http.StatusBadRequest, gin.H{"message": errorMessages})
			return
		}
	}
}
