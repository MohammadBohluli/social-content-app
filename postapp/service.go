package postapp

import (
	"context"

	"github.com/MohammadBohluli/social-content-app/pkg/richerror"
	vld "github.com/MohammadBohluli/social-content-app/pkg/validator"
	"github.com/go-playground/validator/v10"
)

type Repository interface {
	GetPost()
	GetAllPost()
	CreatePost(ctx context.Context, p Post) (Post, error)
	UpdatePost()
	DeletePost()
}

type Service struct {
	repo      Repository
	validator *validator.Validate
}

func NewService(r Repository, v *validator.Validate) Service {
	return Service{
		repo:      r,
		validator: v,
	}
}

func (s Service) GetPost() {}

func (s Service) GetAllPost() {}

func (s Service) CreatePost(ctx context.Context, reqPost CreatePostReq) (CreatePostResp, error) {
	const op = "service.CreatePost"

	if err := s.validator.Struct(reqPost); err != nil {
		validationErrors := vld.ParseValidationErrors(err)
		return CreatePostResp{FieldErrors: validationErrors}, err
	}

	p := Post{
		Content: reqPost.Content,
		Title:   reqPost.Title,
		Tags:    reqPost.Tags,
		UserID:  reqPost.UserID,
	}
	post, err := s.repo.CreatePost(ctx, p)
	if err != nil {
		return CreatePostResp{}, richerror.New(op).WithErr(err)
	}

	return CreatePostResp{
		ID:        post.ID,
		Title:     post.Title,
		UserID:    post.UserID,
		Content:   post.Content,
		Tags:      post.Tags,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}, nil
}

func (s Service) UpdatePost() {}

func (s Service) DeletePost() {}
