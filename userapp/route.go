package postapp

import (
	"github.com/MohammadBohluli/social-content-app/config"
	httpserver "github.com/MohammadBohluli/social-content-app/pkg/http_server"
)

func (h Handler) SetRoutes(s httpserver.Server) {
	prefix := config.ApiVersion + "/users"
	post := s.Router.Group(prefix)

	post.GET("/:id", h.GetUser)
	post.POST("/:id", h.CreateUser)
	post.PATCH("/:id", h.UpdateUser)
	post.DELETE("/:id", h.DeleteUser)
}
