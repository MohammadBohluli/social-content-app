package postapp

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) GetUser(c echo.Context) error {

	return nil
}

func (h Handler) CreateUser(c echo.Context) error {
	return nil
}

func (h Handler) UpdateUser(c echo.Context) error {
	return nil
}

func (h Handler) DeleteUser(c echo.Context) error {
	return nil
}
