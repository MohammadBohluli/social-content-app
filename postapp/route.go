package postapp

import (
	"github.com/MohammadBohluli/social-content-app/config"
	httpserver "github.com/MohammadBohluli/social-content-app/pkg/http_server"
)

func (h Handler) SetRoutes(s httpserver.Server) {
	prefix := config.ApiVersion + "/posts"
	post := s.Router.Group(prefix)

	post.GET("", h.GetAllPost)
	post.GET("/:id", h.GetPost)
	post.POST("/:id", h.CreatePost)
	post.PATCH("/:id", h.UpdatePost)
	post.DELETE("/:id", h.DeletePost)
}
