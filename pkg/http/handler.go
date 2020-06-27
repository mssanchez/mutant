package http

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"mutant/pkg/errors"
	"net/http"
	"runtime/debug"
	"strings"
)

// APIHandlerFunc is a function that given a gin Contexts returns a struct or an error
type APIHandlerFunc func(ctx *gin.Context) (interface{}, error)

// WithinContext handles a request (logging the request, handles errors, status codes, etc)
func WithinContext(handler APIHandlerFunc, log *logrus.Logger) gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				stack := string(debug.Stack())
				onError(ginContext, log, errors.Newf("%v.\n%s", err, stack))
			}
		}()

		logRequest(ginContext, log)

		if payload, err := handler(ginContext); err != nil {
			onError(ginContext, log, err)
		} else {
			ginContext.JSON(http.StatusOK, payload)
		}
	}
}

func onError(ctx *gin.Context, log *logrus.Logger, err error) {
	log.Errorf("%s", err)
	apiError := NewAPIError(err)
	ctx.AbortWithStatusJSON(apiError.StatusCode, apiError)
}

func logRequest(ctx *gin.Context, log *logrus.Logger) {
	var body []byte
	if ctx.Request.Body != nil {
		body, _ = ioutil.ReadAll(ctx.Request.Body)
	} else {
		body = make([]byte, 0)
	}

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	log.Infof("[REQUEST] %s %s %s", ctx.Request.Method, ctx.Request.URL, string(body))
}

// DecodeBody transforms a request body into a struct
func DecodeBody(request *http.Request, target interface{}) error {
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(target); err != nil {
		return errors.UserError.Wrapf(err, "Wrong JSON format: %v", request.Body)
	}

	return nil
}

// CheckContentType verifies that the given content-type is valid
func CheckContentType(request *http.Request) error {
	contentType := request.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return errors.StatusUnsupportedMediaType.Newf("invalid Content-Type, expect `application/json`, got `%s`", contentType)
	}
	return nil
}
