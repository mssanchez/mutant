package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"mutant/pkg/routes"
)

// Initialize a Gin router and register routes
func NewGinHandler(routes *routes.HttpRoutes) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	e := gin.New()

	routes.AddAllHttpRoutes(e)

	return e
}

func StartServer(s *http.Server, log *logrus.Logger) {
	log.Infof("Starting Server on %s\n", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Warnf("Server shutted down [%s]", err)
		panic(err)
	}
}

//This function will block until a SIGINT occur
func ListenShutdownSignal(s *http.Server, log *logrus.Logger, shutdown func()) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info("Shutting down server")

	ctx, cancelFunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelFunc()

	if err := s.Shutdown(ctx); err != nil {
		log.Infof("Server shutdown error [%s]", err)
	} else {
		log.Info("Server exited")
	}

	c := make(chan struct{})
	go func() {
		defer close(c)
		shutdown()
	}()

	select {
	case <-ctx.Done():
	case <-c:
	}
}
