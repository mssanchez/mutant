package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mutant/api"
	"mutant/pkg/errors"
	"mutant/pkg/http"
	"mutant/pkg/mutant"
)

type mutantHandler struct {
	log           *logrus.Logger
	mutantService mutant.Mutant
}

func addMutantHandler(e *gin.Engine, log *logrus.Logger, m mutant.Mutant) {
	handler := &mutantHandler{
		log:           log,
		mutantService: m,
	}

	e.POST("/mutant", http.WithinContext(handler.Mutant, log))
}

func (s *mutantHandler) Mutant(ctx *gin.Context) (interface{}, error) {

	if err := http.CheckContentType(ctx.Request); err != nil {
		return nil, err
	}

	var apiMutant api.Mutant
	if err := http.DecodeBody(ctx.Request, &apiMutant); err != nil {
		return nil, err
	}

	if result, err := s.mutantService.IsMutant(apiMutant.Dna); err != nil {
		return nil, err
	} else if !*result {
		return nil, errors.Forbidden.Newf("not mutant")
	}

	return nil, nil
}
