package service

import (
	"context"
	"post-api/dto"
	"post-api/repository"
)

type PostService interface {
	GetAll(context.Context) ([]*dto.PostResponse, error)
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
