package postapp

import (
	"net/http"

	"github.com/MohammadBohluli/social-content-app/pkg/errormessage"
	jsonresponse "github.com/MohammadBohluli/social-content-app/pkg/json_response"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(s Service) Handler {
	return Handler{service: s}
}

func (h Handler) GetPost(c echo.Context) error {

	return nil
}

func (h Handler) GetAllPost(c echo.Context) error {
	return nil
}

func (h Handler) CreatePost(c echo.Context) error {
	var postReq CreatePostReq
	if err := c.Bind(&postReq); err != nil {
		return jsonresponse.ErrorResponse(c, http.StatusBadRequest, errormessage.ErrorMsgInvalidInput)
	}

	// TODO: remove fake user id
	post := CreatePostReq{
		Content: postReq.Content,
		Title:   postReq.Title,
		Tags:    postReq.Tags,
		UserID:  1,
	}

	resp, err := h.service.CreatePost(c.Request().Context(), post)
	if err != nil {
		return jsonresponse.ErrorResponseValidation(c, err, resp.FieldErrors)
	}

	return jsonresponse.Response(c, http.StatusCreated, resp)
}

func (h Handler) UpdatePost(c echo.Context) error {
	return nil
}

func (h Handler) DeletePost(c echo.Context) error {
	return nil
}
