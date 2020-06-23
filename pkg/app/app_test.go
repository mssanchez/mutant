package app

import (
	"github.com/stretchr/testify/assert"
	"mutant/pkg/config"
	"testing"
)

func TestNewAppInvalidStorage(t *testing.T) {
	configuration := config.Configuration{
		App: config.App{
			Mongodb: config.Mongodb{
				DatabaseName:   "test",
				CollectionName: "test",
				Url:            "test",
				Password:       "test",
			},
		},
		Environment: "test",
	}

	assert.Panics(t, func() { NewApplication(configuration) })
}
