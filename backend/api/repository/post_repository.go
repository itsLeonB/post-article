package repository

import (
	"context"
	"database/sql"
	"fmt"
	"post-api/apperror"
	"post-api/entity"
)

const postRepoFile = "post_repository.go"

type PostRepository interface {
	GetAll(context.Context) ([]*entity.Post, error)
}

type postRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *postRepositoryImpl {
	return &postRepositoryImpl{db}
}

func (r *postRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Post, error) {
	sql := `
		SELECT
			posts.id,
			posts.title,
			posts.content,
			posts.category,
			posts.created_date,
			posts.updated_date,
			posts.status_id,
			post_statuses.name
		FROM posts JOIN post_statuses ON posts.status_id = post_statuses.id
	`

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return nil, apperror.NewError(
			err,
			postRepoFile,
			"postRepositoryImpl.GetAll()",
			fmt.Sprintf("r.db.PrepareContext(%s)", sql),
		)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, apperror.NewError(
			err,
			postRepoFile,
			"postRepositoryImpl.GetAll()",
			"stmt.QueryContext()",
		)
	}

	posts := []*entity.Post{}
	for rows.Next() {
		post := new(entity.Post)
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Category,
			&post.CreatedDate,
			&post.UpdatedDate,
			&post.StatusID,
			&post.Status,
		)
		if err != nil {
			return nil, apperror.NewError(
				err,
				postRepoFile,
				"postRepositoryImpl.GetAll()",
				"rows.Scan()",
			)
		}

		posts = append(posts, post)
	}

	return posts, nil
}
