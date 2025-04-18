package richerror

import (
	"fmt"
	"net/http"

	"github.com/MohammadBohluli/social-content-app/pkg/errormessage"
)

func Error(err error) (message string, code int) {
	switch err := err.(type) {
	case RichError:

		msg := err.Message()

		// TODO: better logger
		fmt.Println("❌ errors: ", err.Error())
		code := mapKindToHTTPStatusCode(err.Kind())

		// we should not expose unexpected error messages
		if code >= 500 {
			msg = errormessage.ErrorMsgSomethingWentWrong
		}

		return msg, code
	default:
		return err.Error(), http.StatusBadRequest
	}
}

func mapKindToHTTPStatusCode(kind Kind) int {
	switch kind {
	case KindInvalid:
		return http.StatusUnprocessableEntity
	case KindNotFound:
		return http.StatusNotFound
	case KindForbidden:
		return http.StatusForbidden
	case KindUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}
