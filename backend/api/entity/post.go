package entity

import (
	"post-api/dto"
	"time"
)

type Post struct {
	ID          int64
	Title       string
	Content     string
	Category    string
	CreatedDate time.Time
	UpdatedDate time.Time
	StatusID    int64
}

func (p *Post) ToResponse() *dto.PostResponse {
	return &dto.PostResponse{
		ID:       p.ID,
		Title:    p.Title,
		Content:  p.Content,
		Category: p.Category,
		StatusID: p.StatusID,
	}
}
