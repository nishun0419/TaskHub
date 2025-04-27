package utils

// ErrorMessages contains all error messages used in the application
var ErrorMessages = struct {
	// Authentication errors
	InvalidCredentials    string
	TokenGenerationFailed string
	InvalidRequest        string

	// Validation errors
	RequiredField    string
	MinLength        string
	MaxLength        string
	InvalidEmail     string
	ValidationFailed string

	// Field names
	FieldNames struct {
		Username string
		Email    string
		Password string
	}
}{
	// Authentication errors
	InvalidCredentials:    "メールアドレスまたはパスワードが正しくありません。",
	TokenGenerationFailed: "認証トークンの生成に失敗しました。",
	InvalidRequest:        "リクエストを正常に受け付けることができませんでした。",

	// Validation errors
	RequiredField:    "%sは必須項目です。",
	MinLength:        "%sは%s文字以上で入力してください。",
	MaxLength:        "%sは%s文字以下で入力してください。",
	InvalidEmail:     "%sは有効なメールアドレスではありません。",
	ValidationFailed: "%sフィールドのバリデーションに失敗しました。",

	// Field names
	FieldNames: struct {
		Username string
		Email    string
		Password string
	}{
		Username: "ユーザー名",
		Email:    "メールアドレス",
		Password: "パスワード",
	},
}
