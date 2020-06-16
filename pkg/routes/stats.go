package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mutant/api"
	"mutant/pkg/http"
	"mutant/pkg/stats"
)

type statsHandler struct {
	log          *logrus.Logger
	statsService stats.Stats
}

func addStatsHandler(e *gin.Engine, log *logrus.Logger, m stats.Stats) {
	handler := &statsHandler{
		log:          log,
		statsService: m,
	}

	e.GET("/stats", http.WithinContext(handler.stats, log))
}

func (s *statsHandler) stats(ctx *gin.Context) (interface{}, error) {
	if result, err := s.statsService.MutantStats(); err != nil {
		return nil, err
	} else {
		return api.MutantStats{
			CountMutantDna: result.CountMutant,
			CountHumanDna:  result.CountHuman,
			Ratio:          result.Ratio,
		}, nil
	}
}
