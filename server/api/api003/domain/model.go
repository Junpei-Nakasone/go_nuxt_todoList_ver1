package domain

// RequestParam contains RequestParam
type RequestParam struct {
	ID      int    `json:"id"`
	NewTodo string `json:"new_todo"`
}

// UpdateData is for update
type UpdateData struct {
	Name string `json:"name"`
}
