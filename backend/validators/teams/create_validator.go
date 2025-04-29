package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

const (
	CreateErrorRequired = "%sは必須項目です。"
	CreateErrorMin      = "%sは%d文字以上で入力してください。"
	CreateErrorMax      = "%sは%d文字以下で入力してください。"
)

func CreateTeamErrorMessage(err validator.ValidationErrors) []string {
	var messages []string
	for _, err := range err {
		switch err.Field() {
		case "Name":
			messages = append(messages, fmt.Sprintf(CreateErrorRequired, err.Field()))
		case "Description":
			messages = append(messages, fmt.Sprintf(CreateErrorRequired, err.Field()))
		}
	}
	return messages
}
