package teams

type CreateInput struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}
