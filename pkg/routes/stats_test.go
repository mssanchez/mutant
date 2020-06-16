package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mutant/api"
	"mutant/pkg/log"
	"mutant/pkg/routes/mocks"
	"mutant/pkg/stats"
	"mutant/pkg/test_fixture"
	"testing"
)

//go:generate mockgen -destination mocks/mock_Stats.go -package mocks -source ../stats/stats.go

type statsMocks struct {
	ctrl  *gomock.Controller
	stats *mocks.MockStats
}

func (builder *statsMocks) build() *gin.Engine {
	router := test_fixture.SetupRouter()
	addStatsHandler(router, log.NewLogger(true), builder.stats)
	return router
}

func statsRouteSetUp(t *testing.T) (*gin.Engine, *statsMocks) {
	ctrl := gomock.NewController(t)

	mocks := &statsMocks{
		ctrl:  ctrl,
		stats: mocks.NewMockStats(ctrl),
	}

	return mocks.build(), mocks
}

func TestStatsHandler(t *testing.T) {
	t.Run("Stats ok", func(t *testing.T) { statsOk(t) })
	t.Run("Stats error", func(t *testing.T) { statsError(t) })
}

func statsOk(t *testing.T) {
	router, mocks := statsRouteSetUp(t)
	defer mocks.ctrl.Finish()

	mutStats := stats.MutantStats{
		CountMutant: 10,
		CountHuman:  20,
		Ratio:       0.5,
	}

	apiMutStats := api.MutantStats{
		CountMutantDna: 10,
		CountHumanDna:  20,
		Ratio:          0.5,
	}

	mocks.stats.EXPECT().MutantStats().Return(&mutStats, nil)

	request, response := test_fixture.NewRequest("GET", "/stats", nil)

	router.ServeHTTP(response, request)

	actualResponseBody := api.MutantStats{}
	json.NewDecoder(response.Body).Decode(&actualResponseBody)

	assert.Equal(t, response.Code, 200)
	assert.Equal(t, apiMutStats, actualResponseBody)
}

func statsError(t *testing.T) {
	router, mocks := statsRouteSetUp(t)
	defer mocks.ctrl.Finish()

	err := fmt.Errorf("some error")

	mocks.stats.EXPECT().MutantStats().Return(nil, err)

	request, response := test_fixture.NewRequest("GET", "/stats", nil)

	router.ServeHTTP(response, request)

	assert.Equal(t, response.Code, 500)
}
