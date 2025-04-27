package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

const (
	LoginErrorRequired    = "%sは必須項目です。"
	LoginErrorMinLength   = "%sは%s文字以上で入力してください。"
	LoginErrorMaxLength   = "%sは%s文字以下で入力してください。"
	LoginErrorEmailFormat = "%sは有効なメールアドレスではありません。"
	LoginErrorDefault     = "%sフィールドのバリデーションに失敗しました。"
)

// fieldNameMap maps field names to their display names in Japanese.
var LoginFieldNameMap = map[string]string{
	"Email":    "メールアドレス",
	"Password": "パスワード",
}

// CreateLoginErrorMessage converts validation errors into user-friendly error messages.
func CreateLoginErrorMessage(err validator.ValidationErrors) []string {
	var errorMessages []string

	for _, fe := range err {
		fieldName := fe.Field()
		displayName := LoginFieldNameMap[fieldName]

		var message string
		switch fe.Tag() {
		case "required":
			message = fmt.Sprintf(LoginErrorRequired, displayName)
		case "min":
			message = fmt.Sprintf(LoginErrorMinLength, displayName, fe.Param())
		case "max":
			message = fmt.Sprintf(LoginErrorMaxLength, displayName, fe.Param())
		case "email":
			message = fmt.Sprintf(LoginErrorEmailFormat, displayName)
		default:
			message = fmt.Sprintf(LoginErrorDefault, displayName)
		}
		errorMessages = append(errorMessages, message)
	}
	return errorMessages
}
