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
	GetByID(context.Context, int64) (*entity.Post, error)
	Insert(context.Context, *entity.Post) error
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

func (r *postRepositoryImpl) GetByID(ctx context.Context, id int64) (*entity.Post, error) {
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
		WHERE posts.id = ?
	`

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return nil, apperror.NewError(
			err,
			postRepoFile,
			"postRepositoryImpl.GetByID()",
			fmt.Sprintf("r.db.PrepareContext(%s)", sql),
		)
	}
	defer stmt.Close()

	post := new(entity.Post)
	err = stmt.QueryRowContext(ctx, id).Scan(
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
			"postRepositoryImpl.GetByID()",
			fmt.Sprintf("stmt.QueryRowContext(%d).Scan()", id),
		)
	}

	return post, nil
}

func (r *postRepositoryImpl) Insert(ctx context.Context, newPost *entity.Post) error {
	dbtx, err := r.getDBTX(ctx)
	if err != nil {
		return err
	}

	sql := `
		INSERT INTO posts (title, content, category, status_id)
		VALUES (?, ?, ?, ?)
		RETURNING id, created_date, updated_date
	`

	stmt, err := dbtx.PrepareContext(ctx, sql)
	if err != nil {
		return apperror.NewError(
			err,
			postRepoFile,
			fmt.Sprintf("postRepositoryImpl.Insert(%v)", *newPost),
			"dbtx.PrepareContext()",
		)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx,
		newPost.Title,
		newPost.Content,
		newPost.Category,
		newPost.StatusID,
	).Scan(
		&newPost.ID,
		&newPost.CreatedDate,
		&newPost.UpdatedDate,
	)
	if err != nil {
		return apperror.NewError(
			err,
			postRepoFile,
			fmt.Sprintf("postRepositoryImpl.Insert(%v)", *newPost),
			"dbtx.QueryRowContext().Scan()",
		)
	}

	return nil
}

func (r *postRepositoryImpl) getDBTX(ctx context.Context) (DBTX, error) {
	tx, err := GetTxFromContext(ctx)
	if err != nil {
		return nil, err
	}
	if tx != nil {
		return tx, nil
	}

	return r.db, nil
}
