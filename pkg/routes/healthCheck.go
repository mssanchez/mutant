package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addHealthCheckRoutes(e *gin.Engine) {
	e.GET("/health-check", healthCheck)
	e.GET("/", healthCheck)
}

func healthCheck(c *gin.Context) {
	c.Header("cache-control", "no-cache")
	c.String(http.StatusOK, "OK")
}
