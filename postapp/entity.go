package postapp

import "github.com/MohammadBohluli/social-content-app/types"

type Post struct {
	ID      types.ID `json:"id"`
	Content string   `json:"content"`
	Title   string   `json:"title"`
	UserID  types.ID `json:"user_id"`
	Tags    []string `json:"tags"`
	// Comments  []Comment `json:"comments"`
	// User      User      `json:"user"`
	Version   int    `json:"version"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
