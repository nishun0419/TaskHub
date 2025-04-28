package customer

import (
	models "backend/models/customer"
	"backend/service/customer"
	"backend/utils"
	validators "backend/validators/customer"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// RegisterHandler handles customer registration requests.
// It validates the input, processes the registration, and returns appropriate responses.
func RegisterHandler(service customer.CustomerServiceInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.RegisterInput

		if err := c.ShouldBindJSON(&input); err != nil {
			var errorMessages []string

			if ve, ok := err.(validator.ValidationErrors); ok {
				errorMessages = validators.CreateRegisterErrorMessage(ve)
			} else {
				errorMessages = append(errorMessages, "リクエストを正常に受け付けることができませんでした。")
			}

			c.JSON(http.StatusBadRequest, utils.ErrorResponse(errorMessages))
			return
		}

		if _, err := service.RegisterCustomer(input); err != nil {
			c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error()))
			return
		}

		c.JSON(http.StatusOK, utils.SuccessResponse("登録が完了しました。", nil))
	}
}
