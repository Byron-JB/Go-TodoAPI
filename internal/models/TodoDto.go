package models

type TodoDto struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Priority    int    `json:"priority"`
	CreatedAt   string `json:"created_at"`
	CompletedAt string `json:"completed_at,omitempty"`
	DueDate     string `json:"due_date"`
}

func SeedTodoData() ([]TodoDto, error) {
	var seedTodoData = []TodoDto{
		{ID: 1, Title: "Learn Go", Description: "Study the Go programming language", Priority: 1, CreatedAt: "2023-10-01", DueDate: "2023-10-01"},
		{ID: 2, Title: "Build a REST API", Description: "Create a RESTful API using Go", Priority: 2, CreatedAt: "2023-10-02", DueDate: "2023-10-01"},
		{ID: 3, Title: "Write Tests", Description: "Write unit tests for the API", Priority: 3, CreatedAt: "2023-10-03", DueDate: "2023-10-01"},
		{ID: 4, Title: "Deploy Application", Description: "Deploy the Go application to a server", Priority: 4, CreatedAt: "2023-10-04", DueDate: "2023-10-01"},
	}
	return seedTodoData, nil
}
