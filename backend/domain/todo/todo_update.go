package todo

type TodoUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CustomerID  int    `json:"customer_id"`
}
