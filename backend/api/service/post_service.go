package service

import (
	"context"
	"post-api/dto"
	"post-api/entity"
	"post-api/repository"
)

// const postSvcFile = "post_service.go"

type PostService interface {
	GetAll(context.Context) ([]*dto.PostResponse, error)
	GetByID(context.Context, int64) (*dto.PostResponse, error)
	Insert(context.Context, *dto.NewPostRequest) (*dto.PostResponse, error)
	Update(context.Context, *dto.UpdatePostRequest) (*dto.PostResponse, error)
	Delete(context.Context, int64) error
}

type postServiceImpl struct {
	trx      repository.Transactor
	postRepo repository.PostRepository
}

func NewPostService(t repository.Transactor, pr repository.PostRepository) *postServiceImpl {
	return &postServiceImpl{t, pr}
}

func (s *postServiceImpl) GetAll(ctx context.Context) ([]*dto.PostResponse, error) {
	posts, err := s.postRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	postResponses := []*dto.PostResponse{}
	for _, post := range posts {
		postResponse := post.ToResponse()
		postResponses = append(postResponses, postResponse)
	}

	return postResponses, nil
}

func (s *postServiceImpl) GetByID(ctx context.Context, id int64) (*dto.PostResponse, error) {
	post, err := s.postRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return post.ToResponse(), nil
}

func (s *postServiceImpl) Insert(ctx context.Context, newPost *dto.NewPostRequest) (*dto.PostResponse, error) {
	ctx, err := s.trx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer s.trx.Rollback(ctx)

	creatingPost := entity.Post{
		Title:    newPost.Title,
		Content:  newPost.Content,
		Category: newPost.Category,
		StatusID: newPost.StatusID,
	}

	err = s.postRepo.Insert(ctx, &creatingPost)
	if err != nil {
		return nil, err
	}

	err = s.trx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return creatingPost.ToResponse(), nil
}

func (s *postServiceImpl) Update(ctx context.Context, updatePost *dto.UpdatePostRequest) (*dto.PostResponse, error) {
	ctx, err := s.trx.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer s.trx.Rollback(ctx)

	updatingPost := entity.Post{
		ID:       updatePost.ID,
		Title:    updatePost.Title,
		Content:  updatePost.Content,
		Category: updatePost.Category,
		StatusID: updatePost.StatusID,
	}

	_, err = s.postRepo.GetByID(ctx, updatingPost.ID)
	if err != nil {
		return nil, err
	}

	err = s.postRepo.Update(ctx, &updatingPost)
	if err != nil {
		return nil, err
	}

	err = s.trx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return updatingPost.ToResponse(), nil
}

func (s *postServiceImpl) Delete(ctx context.Context, id int64) error {
	ctx, err := s.trx.Begin(ctx)
	if err != nil {
		return err
	}
	defer s.trx.Rollback(ctx)

	_, err = s.postRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	err = s.postRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	err = s.trx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
