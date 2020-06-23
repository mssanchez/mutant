package stats

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"mutant/pkg/mutant/mocks"
	"testing"
)

//go:generate mockgen -destination mocks/mock_MutantStorage.go -package mocks -source ../storage/mutant.go

type statsMocks struct {
	ctrl          *gomock.Controller
	mutantStorage *mocks.MockMutantStorage
}

func (builder *statsMocks) build() Stats {
	return NewStats(builder.mutantStorage)
}

func setUp(t *testing.T) (Stats, *statsMocks) {
	ctrl := gomock.NewController(t)

	mutMocks := &statsMocks{
		ctrl:          ctrl,
		mutantStorage: mocks.NewMockMutantStorage(ctrl),
	}

	return mutMocks.build(), mutMocks
}

var mutantStats = MutantStats{10, 10, 1}

var mutantStatsZeroHumans = MutantStats{10, 0, 0}

func TestMutantStatsSuccessful(t *testing.T) {
	stats, mutMocks := setUp(t)
	defer mutMocks.ctrl.Finish()

	mutMocks.mutantStorage.EXPECT().
		Count(true).
		Return(int64(10), nil)

	mutMocks.mutantStorage.EXPECT().
		Count(false).
		Return(int64(10), nil)

	response, err := stats.MutantStats()

	assert.NoError(t, err)
	assert.Equal(t, &mutantStats, response)
}

func TestMutantStatsError(t *testing.T) {
	stats, mutMocks := setUp(t)
	defer mutMocks.ctrl.Finish()

	someError := fmt.Errorf("some error")

	mutMocks.mutantStorage.EXPECT().
		Count(true).
		Return(int64(0), someError)

	_, err := stats.MutantStats()

	assert.EqualError(t, err, someError.Error())
}

func TestMutantStatsZeroHumans(t *testing.T) {
	stats, mutMocks := setUp(t)
	defer mutMocks.ctrl.Finish()

	mutMocks.mutantStorage.EXPECT().
		Count(true).
		Return(int64(10), nil)

	mutMocks.mutantStorage.EXPECT().
		Count(false).
		Return(int64(0), nil)

	response, err := stats.MutantStats()

	assert.NoError(t, err)
	assert.Equal(t, &mutantStatsZeroHumans, response)
}
