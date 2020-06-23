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

func mutantSetUp(t *testing.T) (*gin.Engine, *mutantMocks) {
	ctrl := gomock.NewController(t)

	mutantMocks := &mutantMocks{
		ctrl:   ctrl,
		mutant: mocks.NewMockMutant(ctrl),
	}

	return mutantMocks.build(), mutantMocks
}

var mutantApi = api.Mutant{
	Dna: []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG"},
}

func TestMutantHandler(t *testing.T) {
	t.Run("Is mutant", func(t *testing.T) { isMutant(t) })
	t.Run("Is not mutant", func(t *testing.T) { isNotMutant(t) })
	t.Run("Mutant error", func(t *testing.T) { mutantError(t) })
	t.Run("Mutant panic error", func(t *testing.T) { mutantPanicError(t) })
	t.Run("Mutant invalid content-type", func(t *testing.T) { mutantInvalidContentType(t) })
	t.Run("Mutant decode body error", func(t *testing.T) { mutantDecodeBodyError(t) })
}

func isMutant(t *testing.T) {
	router, mutMocks := mutantSetUp(t)
	defer mutMocks.ctrl.Finish()

	true := true
	mutMocks.mutant.EXPECT().
		IsMutant(mutantApi.Dna).
		Return(&true, nil)

	bodyStr, _ := json.Marshal(mutantApi)
	request, response := test_fixture.NewRequest("POST", "/mutant", bytes.NewBuffer(bodyStr))

	router.ServeHTTP(response, request)

	assert.Equal(t, response.Code, 200)
}

func isNotMutant(t *testing.T) {
	router, mutMocks := mutantSetUp(t)
	defer mutMocks.ctrl.Finish()

	false := false
	mutMocks.mutant.EXPECT().
		IsMutant(mutantApi.Dna).
		Return(&false, nil)

	bodyStr, _ := json.Marshal(mutantApi)
	request, response := test_fixture.NewRequest("POST", "/mutant", bytes.NewBuffer(bodyStr))

	router.ServeHTTP(response, request)

	assert.Equal(t, 403, response.Code)
}

func mutantError(t *testing.T) {
	router, mutMocks := mutantSetUp(t)
	defer mutMocks.ctrl.Finish()

	err := fmt.Errorf("some error")

	mutMocks.mutant.EXPECT().
		IsMutant(mutantApi.Dna).
		Return(nil, err)

	bodyStr, _ := json.Marshal(mutantApi)
	request, response := test_fixture.NewRequest("POST", "/mutant", bytes.NewBuffer(bodyStr))

	router.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code)
}

type mutantPanicMock struct {
}

func (_m *mutantPanicMock) IsMutant(dna []string) (*bool, error) {
	panic("something happened")
}

func mutantPanicError(t *testing.T) {
	panicMock := mutantPanicMock{}
	router := test_fixture.SetupRouter()
	addMutantHandler(router, log.NewLogger(true), &panicMock)

	bodyStr, _ := json.Marshal(mutantApi)
	request, response := test_fixture.NewRequest("POST", "/mutant", bytes.NewBuffer(bodyStr))

	router.ServeHTTP(response, request)

	assert.Equal(t, 500, response.Code)
}

func mutantInvalidContentType(t *testing.T) {
	router, mutMocks := mutantSetUp(t)
	defer mutMocks.ctrl.Finish()

	bodyStr, _ := json.Marshal(mutantApi)
	request, response := test_fixture.NewRequestWithContentType("POST", "/mutant", bytes.NewBuffer(bodyStr), "invalid-ct")

	router.ServeHTTP(response, request)

	assert.Equal(t, 415, response.Code)
}

func mutantDecodeBodyError(t *testing.T) {
	router, mutMocks := mutantSetUp(t)
	defer mutMocks.ctrl.Finish()

	request, response := test_fixture.NewRequest("POST", "/mutant", bytes.NewBufferString("test"))

	router.ServeHTTP(response, request)

	assert.Equal(t, 400, response.Code)
}
