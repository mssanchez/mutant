package http

import (
	"net/http"

	"mutant/api"
	"mutant/pkg/errors"
)

// NewAPIError instantiates a new error returned by the API
func NewAPIError(err error) api.Error {
	if customError, ok := err.(errors.Error); ok {
		var statusCode int

		switch customError.Type() {
		case errors.StatusUnsupportedMediaType:
			statusCode = http.StatusUnsupportedMediaType
		case errors.Forbidden:
			statusCode = http.StatusForbidden
		case errors.UserError:
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusInternalServerError
		}

		return api.Error{
			Cause:      customError.Error(),
			Context:    customError.Context(),
			StatusCode: statusCode,
		}
	}

	return api.Error{
		Cause:      err.Error(),
		StatusCode: http.StatusInternalServerError,
	}
}
