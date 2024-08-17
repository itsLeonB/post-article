package service

import (
	"context"
	"post-api/dto"
	"post-api/repository"
)

type PostService interface {
	GetAll(context.Context) ([]*dto.PostResponse, error)
	GetByID(context.Context, int64) (*dto.PostResponse, error)
}

type postServiceImpl struct {
	postRepo repository.PostRepository
}

func NewPostService(pr repository.PostRepository) *postServiceImpl {
	return &postServiceImpl{pr}
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
