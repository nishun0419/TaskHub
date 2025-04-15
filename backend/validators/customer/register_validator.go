// Package validators provides validation utilities for the application.
// It includes functions for validating user input and generating appropriate error messages.
package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

const (
	errorRequired    = "%sは必須項目です。"
	errorMinLength   = "%sは%s文字以上で入力してください。"
	errorMaxLength   = "%sは%s文字以下で入力してください。"
	errorEmailFormat = "%sは有効なメールアドレスではありません。"
	errorDefault     = "%sフィールドのバリデーションに失敗しました。"
)

// fieldNameMap maps field names to their display names in Japanese.
var fieldNameMap = map[string]string{
	"Username": "ユーザー名",
	"Email":    "メールアドレス",
	"Password": "パスワード",
}

// CreateRegisterErrorMessage converts validation errors into user-friendly error messages.
// It takes a validator.ValidationErrors and returns a slice of formatted error messages.
// The messages are in Japanese and are formatted according to the validation rules that failed.
func CreateRegisterErrorMessage(err validator.ValidationErrors) []string {
	var errorMessages []string

	for _, fe := range err {
		fieldName := fe.Field()
		displayName := fieldNameMap[fieldName]

		var message string
		switch fe.Tag() {
		case "required":
			message = fmt.Sprintf(errorRequired, displayName)
		case "min":
			message = fmt.Sprintf(errorMinLength, displayName, fe.Param())
		case "max":
			message = fmt.Sprintf(errorMaxLength, displayName, fe.Param())
		case "email":
			message = fmt.Sprintf(errorEmailFormat, displayName)
		default:
			message = fmt.Sprintf(errorDefault, displayName)
		}
		errorMessages = append(errorMessages, message)
	}
	return errorMessages
}
