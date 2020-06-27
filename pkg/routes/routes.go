package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mutant/pkg/mutant"
	"mutant/pkg/stats"
)

// HTTPRoutes is in charge of adding all http routes
type HTTPRoutes struct {
	log    *logrus.Logger
	mutant mutant.Mutant
	stats  stats.Stats
}

// NewHTTPRoutes builds an HTTPRoutes
func NewHTTPRoutes(log *logrus.Logger, mutant mutant.Mutant, stats stats.Stats) HTTPRoutes {
	log.Info("init routes package...")

	return HTTPRoutes{
		log:    log,
		mutant: mutant,
		stats:  stats,
	}
}

// AddAllHTTPRoutes adds health check, mutant and stats routes
func (r *HTTPRoutes) AddAllHTTPRoutes(e *gin.Engine) {
	addHealthCheckRoutes(e)
	addMutantHandler(e, r.log, r.mutant)
	addStatsHandler(e, r.log, r.stats)
}
