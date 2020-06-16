package main

import (
	"mutant/pkg/app"
	"mutant/pkg/config"
)

func main() {
	configuration := config.NewConfiguration()

	application := app.NewApplication(configuration)

	application.RunServer()
}
