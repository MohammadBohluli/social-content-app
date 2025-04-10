package postapp

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) GetPost(c echo.Context) error {

	return nil
}

func (h Handler) GetAllPost(c echo.Context) error {
	return nil
}

func (h Handler) CreatePost(c echo.Context) error {
	return nil
}

func (h Handler) UpdatePost(c echo.Context) error {
	return nil
}

func (h Handler) DeletePost(c echo.Context) error {
	return nil
}
