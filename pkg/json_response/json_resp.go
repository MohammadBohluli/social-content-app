package jsonresponse

import (
	"github.com/MohammadBohluli/social-content-app/pkg/errormessage"
	"github.com/MohammadBohluli/social-content-app/pkg/richerror"
	"github.com/MohammadBohluli/social-content-app/types"
	"github.com/labstack/echo/v4"
)

type response struct {
	Data any `json:"data"`
}

func Response(c echo.Context, statusCode int, data any) error {
	d := response{
		Data: data,
	}
	return c.JSON(statusCode, d)
}

func errorResponseValidation(c echo.Context, status int, message string, fieldErrors types.FieldErrors) error {
	resp := echo.Map{
		"message": message,
	}

	if len(fieldErrors) > 0 {
		resp["errors"] = fieldErrors
	}

	return c.JSON(status, resp)
}

func ErrorResponseValidation(c echo.Context, err error, fieldErrors types.FieldErrors) error {
	if err == nil {
		return nil
	}

	msg, code := richerror.Error(err)

	if len(fieldErrors) > 0 {
		return errorResponseValidation(c, code, errormessage.ErrorMsgValidation, fieldErrors)
	}

	return errorResponseValidation(c, code, msg, nil)
}

func ErrorResponse(c echo.Context, status int, msg string) error {
	return c.JSON(status, map[string]string{"message": msg})
}
