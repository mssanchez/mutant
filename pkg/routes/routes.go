package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mutant/pkg/mutant"
	"mutant/pkg/stats"
)

type HttpRoutes struct {
	log    *logrus.Logger
	mutant mutant.Mutant
	stats  stats.Stats
}

func NewHttpRoutes(log *logrus.Logger, mutant mutant.Mutant, stats stats.Stats) HttpRoutes {
	log.Info("init routes package...")

	return HttpRoutes{
		log:    log,
		mutant: mutant,
		stats:  stats,
	}
}

func (r *HttpRoutes) AddAllHttpRoutes(e *gin.Engine) {
	addHealthCheckRoutes(e)
	addMutantHandler(e, r.log, r.mutant)
	addStatsHandler(e, r.log, r.stats)
}
