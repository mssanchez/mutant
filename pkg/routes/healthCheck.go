package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addHealthCheckRoutes(e *gin.Engine) {
	e.GET("/health", healthCheck)
}

func healthCheck(c *gin.Context) {
	c.Header("cache-control", "no-cache")
	c.String(http.StatusOK, "OK")
}
