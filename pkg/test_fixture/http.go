package test_fixture

import (
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	return r
}

func NewRequest(method, url string, body io.Reader) (*http.Request, *httptest.ResponseRecorder) {
	request, _ := http.NewRequest(method, url, body)
	request.Header.Add("Content-Type", "application/json")
	return request, httptest.NewRecorder()
}
