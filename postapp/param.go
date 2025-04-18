package postapp

import "github.com/MohammadBohluli/social-content-app/types"

type CreatePostReq struct {
	Content string   `json:"content" validate:"required,min=5,max=20"`
	Title   string   `json:"title" validate:"required,min=5,max=20"`
	Tags    []string `json:"tags"`
	UserID  types.ID `json:"user_id"`
}

type CreatePostResp struct {
	ID          types.ID          `json:"id"`
	Title       string            `json:"title"`
	UserID      types.ID          `json:"user_id"`
	Content     string            `json:"content"`
	Tags        []string          `json:"tags"`
	CreatedAt   string            `json:"created_at"`
	UpdatedAt   string            `json:"updated_at"`
	FieldErrors types.FieldErrors `json:"field_errors,omitempty"`
}
