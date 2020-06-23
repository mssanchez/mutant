package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mutant/pkg/test_fixture"
	"testing"
)

type healthCheckMocks struct {
	ctrl *gomock.Controller
}

func (builder *healthCheckMocks) build() *gin.Engine {
	router := test_fixture.SetupRouter()
	addHealthCheckRoutes(router)

	return router
}

func healthCheckSetUp(t *testing.T) (*gin.Engine, *healthCheckMocks) {
	ctrl := gomock.NewController(t)

	hcMocks := &healthCheckMocks{
		ctrl: ctrl,
	}

	return hcMocks.build(), hcMocks
}

func TestHealthCheckHandler(t *testing.T) {
	router, hcMocks := healthCheckSetUp(t)
	defer hcMocks.ctrl.Finish()

	request, response := test_fixture.NewRequest("GET", "/health-check", nil)

	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)
}
