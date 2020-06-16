package app

import (
	"github.com/sirupsen/logrus"
	"mutant/pkg/config"
	"mutant/pkg/log"
	"mutant/pkg/mutant"
	"mutant/pkg/routes"
	"mutant/pkg/server"
	"mutant/pkg/stats"
	"mutant/pkg/storage"
	"net/http"
)

// Used to initialize the application
type Application struct {
	log           *logrus.Logger
	httpRoutes    *routes.HttpRoutes
	configuration config.Configuration
	mutantStorage storage.MutantStorage
}

// Instantiates a new application
func NewApplication(configuration config.Configuration) *Application {
	logger := log.NewLogger(true)

	mutantStorage := storage.NewMutantsStorage(configuration, logger)

	mutantSvc := mutant.NewMutant(mutantStorage)

	mutantStats := stats.NewStats(mutantStorage)

	httpRoutes := routes.NewHttpRoutes(logger, mutantSvc, mutantStats)

	return &Application{
		log:           logger,
		httpRoutes:    &httpRoutes,
		configuration: configuration,
		mutantStorage: mutantStorage,
	}
}

// Starts the server in the configured port
func (app *Application) RunServer() {

	s := &http.Server{
		Addr:    app.configuration.Server.Port,
		Handler: server.NewGinHandler(app.httpRoutes),
	}

	go server.StartServer(s, app.log)

	server.ListenShutdownSignal(s, app.log, app.ShutDown)
}

func (app *Application) ShutDown() {
	app.mutantStorage.Shutdown()
}
