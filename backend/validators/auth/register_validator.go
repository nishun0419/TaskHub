package validators

import (
    "fmt"
    "github.com/go-playground/validator/v10"
)

// フィールド名をユーザー向けに変換する配列
var returnFieldNames = map[string]string{
	"Username": "ユーザー名",
	"Password": "パスワード",
}

func CreateRegisterErrorMessage(err validator.ValidationErrors) []string {
	// エラーメッセージを格納する変数
  var errorMessages []string

  for _, fe := range err {
      fieldName := fe.Field()
      switch fe.Tag() {
      case "required":
          errorMessages = append(errorMessages, fmt.Sprintf("%sは必須項目です。", returnFieldNames[fieldName]))
      case "min":
          errorMessages = append(errorMessages, fmt.Sprintf("%sは%s文字以上で入力してください。", returnFieldNames[fieldName], fe.Param()))
      case "max":
          errorMessages = append(errorMessages, fmt.Sprintf("%sは%s文字以下で入力してください。", returnFieldNames[fieldName], fe.Param()))
      case "email":
          errorMessages = append(errorMessages, fmt.Sprintf("%sは有効なメールアドレスではありません。", returnFieldNames[fieldName]))
      default:
          errorMessages = append(errorMessages, fmt.Sprintf("%sフィールドのバリデーションに失敗しました。", returnFieldNames[fieldName]))
      }
  }
  return errorMessages
}
