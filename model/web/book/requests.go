package book

type CreateRequest struct {
	Title string `json:"title"`
	Status string `json:"status"`
	Author string `json:"author"`
	Year string `json:"year"`
	UserID  string `json:"user_id"`
}

type FindByIdRequest struct {
	Title string `json:"title"`
	Status string `json:"status"`
	Author string `json:"author"`
	Year string `json:"year"`
	UserID  string `json:"user_id"`
}

