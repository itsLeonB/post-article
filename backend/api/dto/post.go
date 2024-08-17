package dto

type PostResponse struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
	StatusID int64  `json:"status_id"`
}

type NewPostRequest struct {
	Title    string `json:"title" binding:"required,min=20,max=200"`
	Content  string `json:"content" binding:"required,min=200"`
	Category string `json:"category" binding:"required,min=3,max=100"`
	StatusID int64  `json:"status_id" binding:"required,numeric,min=1"`
}

type UpdatePostRequest struct {
	ID       int64  `json:"-"`
	Title    string `json:"title" binding:"required,min=20,max=200"`
	Content  string `json:"content" binding:"required,min=200"`
	Category string `json:"category" binding:"required,min=3,max=100"`
	StatusID int64  `json:"status_id" binding:"required,numeric,min=1"`
}

type PostStatusResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
