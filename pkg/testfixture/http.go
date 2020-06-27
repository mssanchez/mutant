package testfixture

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// SetupRouter instantiates a new Gin Engine for testing
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return r
}

// NewRequest builds an http Request given a method, url and body, and application/json content type
func NewRequest(method, url string, body io.Reader) (*http.Request, *httptest.ResponseRecorder) {
	return NewRequestWithContentType(method, url, body, "application/json")
}

// NewRequestWithContentType builds an http Request given a method, url, body and content type
func NewRequestWithContentType(method, url string, body io.Reader, contentType string) (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest(method, url, body)
	request.Header.Add("Content-Type", contentType)
	return request, httptest.NewRecorder()
}
