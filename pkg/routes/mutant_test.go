package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mutant/api"
	"mutant/pkg/log"
	"mutant/pkg/routes/mocks"
	"mutant/pkg/test_fixture"
	"testing"
)

//go:generate mockgen -destination mocks/mock_Mutant.go -package mocks -source ../mutant/mutant.go

type mutantMocks struct {
	ctrl   *gomock.Controller
	mutant *mocks.MockMutant
}

func (builder *mutantMocks) build() *gin.Engine {
	router := test_fixture.SetupRouter()
	addMutantHandler(router, log.NewLogger(true), builder.mutant)
	return router
}

func routeSetUp(t *testing.T) (*gin.Engine, *mutantMocks) {
	ctrl := gomock.NewController(t)

	mocks := &mutantMocks{
		ctrl:   ctrl,
		mutant: mocks.NewMockMutant(ctrl),
	}

	return mocks.build(), mocks
}

var mutantApi = api.Mutant{
	Dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG"},
}

func TestMutantHandler(t *testing.T) {
	t.Run("Is mutant", func(t *testing.T) { isMutant(t) })
	t.Run("Is not mutant", func(t *testing.T) { isNotMutant(t) })
	t.Run("Mutant error", func(t *testing.T) { mutantError(t) })
}

func isMutant(t *testing.T) {
	router, mocks := routeSetUp(t)
	defer mocks.ctrl.Finish()

	true := true
	mocks.mutant.EXPECT().
		IsMutant(mutantApi.Dna).
		Return(&true, nil)

	bodyStr, _ := json.Marshal(mutantApi)
	request, response := test_fixture.NewRequest("POST", "/mutant", bytes.NewBuffer(bodyStr))

	router.ServeHTTP(response, request)

	assert.Equal(t, response.Code, 200)
}

func isNotMutant(t *testing.T) {
	router, mocks := routeSetUp(t)
	defer mocks.ctrl.Finish()

	false := false
	mocks.mutant.EXPECT().
		IsMutant(mutantApi.Dna).
		Return(&false, nil)

	bodyStr, _ := json.Marshal(mutantApi)
	request, response := test_fixture.NewRequest("POST", "/mutant", bytes.NewBuffer(bodyStr))

	router.ServeHTTP(response, request)

	assert.Equal(t, response.Code, 403)
}

func mutantError(t *testing.T) {
	router, mocks := routeSetUp(t)
	defer mocks.ctrl.Finish()

	err := fmt.Errorf("some error")

	mocks.mutant.EXPECT().
		IsMutant(mutantApi.Dna).
		Return(nil, err)

	bodyStr, _ := json.Marshal(mutantApi)
	request, response := test_fixture.NewRequest("POST", "/mutant", bytes.NewBuffer(bodyStr))

	router.ServeHTTP(response, request)

	assert.Equal(t, response.Code, 500)
}
