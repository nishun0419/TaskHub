package customer

import (
	models "backend/models/customer"
	"backend/service/customer"
	validators "backend/validators/customer"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RegisterHandler(service *customer.CustomerService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.RegisterInput

		if err := c.ShouldBindJSON(&input); err != nil {
			var errorMessages []string

			if ve, ok := err.(validator.ValidationErrors); ok {
				errorMessages = validators.CreateRegisterErrorMessage(ve)
			} else {
				errorMessages = append(errorMessages, "リクエストを正常に受け付けることができませんでした。")
			}

			c.JSON(http.StatusBadRequest, gin.H{"message": errorMessages})
			return
		}
		if _, err := service.RegisterCustomer(input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": []string{err.Error()}})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": []string{"登録OK"},
		})
	}
}
