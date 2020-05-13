package domain

// Todo contains todo
type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ResponseData returns todos
type ResponseData struct {
	Todos []Todo `json:"todos"`
}
