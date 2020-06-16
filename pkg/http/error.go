package http

import (
	"net/http"

	"mutant/api"
	"mutant/pkg/errors"
)

// Instantiates a new error returned by the API
func NewApiError(err error) api.Error {
	if customError, ok := err.(errors.Error); ok {
		var statusCode int

		switch customError.Type() {
		case errors.UserError:
			statusCode = http.StatusBadRequest
		case errors.NotFound:
			statusCode = http.StatusNotFound
		case errors.StatusUnsupportedMediaType:
			statusCode = http.StatusUnsupportedMediaType
		case errors.NotModified:
			statusCode = http.StatusNotModified
		case errors.Forbidden:
			statusCode = http.StatusForbidden
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
