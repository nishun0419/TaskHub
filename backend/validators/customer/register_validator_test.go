package validators

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type RegisterRequest struct {
	Username string `validate:"required,min=3,max=20"`
	Password string `validate:"required,min=8"`
	Email    string `validate:"required,email"`
}

func TestCreateRegisterErrorMessage(t *testing.T) {
	v := validator.New()

	tests := []struct {
		name     string
		input    RegisterRequest
		expected []string
	}{
		{
			name: "全てのフィールドが空の場合",
			input: RegisterRequest{
				Username: "",
				Password: "",
				Email:    "",
			},
			expected: []string{
				"ユーザー名は必須項目です。",
				"パスワードは必須項目です。",
				"メールアドレスは必須項目です。",
			},
		},
		{
			name: "無効なメールアドレスの場合",
			input: RegisterRequest{
				Username: "test",
				Password: "password123",
				Email:    "invalid-email",
			},
			expected: []string{
				"メールアドレスは有効なメールアドレスではありません。",
			},
		},
		{
			name: "ユーザー名が短すぎる場合",
			input: RegisterRequest{
				Username: "te",
				Password: "password123",
				Email:    "test@example.com",
			},
			expected: []string{
				"ユーザー名は3文字以上で入力してください。",
			},
		},
		{
			name: "ユーザー名が長すぎる場合",
			input: RegisterRequest{
				Username: "thisusernameistoolongforthevalidation",
				Password: "password123",
				Email:    "test@example.com",
			},
			expected: []string{
				"ユーザー名は20文字以下で入力してください。",
			},
		},
		{
			name: "パスワードが短すぎる場合",
			input: RegisterRequest{
				Username: "testuser",
				Password: "pass",
				Email:    "test@example.com",
			},
			expected: []string{
				"パスワードは8文字以上で入力してください。",
			},
		},
		{
			name: "全てのバリデーションをパスする場合",
			input: RegisterRequest{
				Username: "testuser",
				Password: "password123",
				Email:    "test@example.com",
			},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Struct(tt.input)
			if err != nil {
				validationErrors := err.(validator.ValidationErrors)
				messages := CreateRegisterErrorMessage(validationErrors)
				assert.ElementsMatch(t, tt.expected, messages)
			} else {
				assert.Empty(t, tt.expected, "エラーが発生しないはずのケースでエラーが発生しました")
			}
		})
	}
}
