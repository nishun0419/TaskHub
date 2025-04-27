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

			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": errorMessages,
			})
			return
		}

		// ユーザー認証を行う
		user, err := service.Authenticate(input.Email, input.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": []string{"メールアドレスまたはパスワードが正しくありません。"},
			})
			return
		}

		// JWTトークンを生成
		token, err := service.GenerateToken(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": []string{"認証トークンの生成に失敗しました。"},
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"token":   token,
			"user": gin.H{
				"customer_id": user.CustomerID,
				"username":    user.Username,
				"email":       user.Email,
			},
		})
	}
}
