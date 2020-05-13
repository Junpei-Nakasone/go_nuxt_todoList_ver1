package domain

// Todo contains todo
type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RequestParam struct {
	Name string `json:"name"`
}
