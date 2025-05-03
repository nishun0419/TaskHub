package customer

import (
	domain "backend/domain/customer"
	usecase "backend/usecase/customer"
	utils "backend/utils"
	validators "backend/validators/customer"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomerController struct {
	Usecase *usecase.CustomerUsecase
}

func NewCustomerController(u *usecase.CustomerUsecase) *CustomerController {
	return &CustomerController{Usecase: u}
}

func (c *CustomerController) LoginHandler(ctx *gin.Context) {
	var input domain.LoginInput

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
	user, err := c.Usecase.Authenticate(input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": []string{"メールアドレスまたはパスワードが正しくありません。"},
		})
		return
	}

	// JWTトークンを生成
	token, err := c.Usecase.GenerateToken(user)
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

func (c *CustomerController) RegisterHandler(ctx *gin.Context) {
	var input domain.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		var errorMessages []string

		if ve, ok := err.(validator.ValidationErrors); ok {
			errorMessages = validators.CreateRegisterErrorMessage(ve)
		} else {
			errorMessages = append(errorMessages, "リクエストを正常に受け付けることができませんでした。")
		}

		utils.ErrorResponse(ctx, errorMessages)
		return
	}

	if err := c.Usecase.RegisterCustomer(input); err != nil {
		utils.ErrorResponse(ctx, err.Error())
		return
	}

	utils.SuccessResponse(ctx, "登録が完了しました。", nil)
}
