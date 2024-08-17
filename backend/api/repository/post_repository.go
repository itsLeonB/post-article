package repository

import (
	"context"
	"database/sql"
	"fmt"
	"post-api/appcontext"
	"post-api/apperror"
	"post-api/entity"
)

const postRepoFile = "post_repository.go"

type PostRepository interface {
	GetAll(context.Context) ([]*entity.Post, error)
	GetByID(context.Context, int64) (*entity.Post, error)
	Insert(context.Context, *entity.Post) error
	Update(context.Context, *entity.Post) error
	Delete(context.Context, int64) error
	GetStatus(context.Context) ([]*entity.PostStatus, error)
}

type postRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *postRepositoryImpl {
	return &postRepositoryImpl{db}
}

func (r *postRepositoryImpl) GetAll(ctx context.Context) ([]*entity.Post, error) {
	query := `
		SELECT
			id,
			title,
			content,
			category,
			created_date,
			updated_date,
			status_id
		FROM posts
		
	`

	limit := ctx.Value(appcontext.KeyLimit).(int64)
	offset := ctx.Value(appcontext.KeyOffset).(int64)
	statusID := ctx.Value(appcontext.KeyStatusID).(int64)
	if statusID != 0 {
		query += fmt.Sprintf(" WHERE status_id = %d", statusID)
	}
	query += " ORDER BY updated_date DESC"
	if limit != 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}
	if offset != 0 {
		query += fmt.Sprintf(" OFFSET %d", offset)
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, apperror.NewError(
			err,
			postRepoFile,
			"postRepositoryImpl.GetAll()",
			fmt.Sprintf("r.db.PrepareContext(%s)", query),
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
			id,
			title,
			content,
			category,
			created_date,
			updated_date,
			status_id
		FROM posts
		WHERE id = ?
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

	row, err := stmt.ExecContext(ctx,
		newPost.Title,
		newPost.Content,
		newPost.Category,
		newPost.StatusID,
	)
	if err != nil {
		return apperror.NewError(
			err,
			postRepoFile,
			fmt.Sprintf("postRepositoryImpl.Insert(%v)", *newPost),
			"dbtx.ExecContext()",
		)
	}

	id, err := row.LastInsertId()
	if err != nil {
		return apperror.NewError(
			err,
			postRepoFile,
			fmt.Sprintf("postRepositoryImpl.Insert(%v)", *newPost),
			"row.LastInsertId()",
		)
	}

	newPost.ID = id

	return nil
}

func (r *postRepositoryImpl) Update(ctx context.Context, editingPost *entity.Post) error {
	dbtx, err := r.getDBTX(ctx)
	if err != nil {
		return err
	}

	sql := `
		UPDATE posts
		SET title = ?, content = ?, category = ?, status_id = ?, updated_date = NOW()
		WHERE id = ?
	`

	stmt, err := dbtx.PrepareContext(ctx, sql)
	if err != nil {
		return apperror.NewError(
			err,
			postRepoFile,
			fmt.Sprintf("postRepositoryImpl.Update(%v)", *editingPost),
			"dbtx.PrepareContext()",
		)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		editingPost.Title,
		editingPost.Content,
		editingPost.Category,
		editingPost.StatusID,
		editingPost.ID,
	)
	if err != nil {
		return apperror.NewError(
			err,
			postRepoFile,
			fmt.Sprintf("postRepositoryImpl.Update(%v)", *editingPost),
			"dbtx.ExecContext()",
		)
	}

	return nil
}

func (r *postRepositoryImpl) Delete(ctx context.Context, id int64) error {
	dbtx, err := r.getDBTX(ctx)
	if err != nil {
		return err
	}

	sql := `
		DELETE FROM posts
		WHERE id = ?
	`

	stmt, err := dbtx.PrepareContext(ctx, sql)
	if err != nil {
		return apperror.NewError(
			err,
			postRepoFile,
			fmt.Sprintf("postRepositoryImpl.Delete(%d", id),
			"dbtx.PrepareContext()",
		)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return apperror.NewError(
			err,
			postRepoFile,
			fmt.Sprintf("postRepositoryImpl.Delete(%d)", id),
			"dbtx.ExecContext()",
		)
	}

	return nil
}

func (r *postRepositoryImpl) GetStatus(ctx context.Context) ([]*entity.PostStatus, error) {
	query := `
		SELECT id, name
		FROM post_statuses
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, apperror.NewError(
			err,
			postRepoFile,
			"postRepositoryImpl.GetStatus()",
			fmt.Sprintf("dbtx.PrepareContext(%s)", query),
		)
	}

	statuses := []*entity.PostStatus{}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, apperror.NewError(
			err,
			postRepoFile,
			"postRepositoryImpl.GetStatus()",
			"stmt.QueryRowContext()",
		)
	}

	for rows.Next() {
		status := new(entity.PostStatus)
		err = rows.Scan(&status.ID, &status.Name)
		if err != nil {
			return nil, apperror.NewError(
				err,
				postRepoFile,
				"postRepositoryImpl.GetStatus()",
				"rows.Scan()",
			)
		}
		statuses = append(statuses, status)
	}

	return statuses, nil
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
