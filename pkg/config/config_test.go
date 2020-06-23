package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMutantHandler(t *testing.T) {
	t.Run("Parse config successful", func(t *testing.T) { parseConfigOk(t) })
	t.Run("Parse config error", func(t *testing.T) { parseConfigError(t) })
}

func parseConfigOk(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")
	config := NewConfiguration()

	// Environment
	assert.Equal(t, "test", config.Environment)

	//MongoDB config
	assert.NotEmpty(t, config.Mongodb.CollectionName)
	assert.NotEmpty(t, config.Mongodb.DatabaseName)
	assert.NotEmpty(t, config.Mongodb.DisconnectTimeoutInSeconds)
	assert.NotEmpty(t, config.Mongodb.Url)

	// Server config
	assert.NotEmpty(t, config.Server.Port)
}

func parseConfigError(t *testing.T) {
	os.Setenv("ENVIRONMENT", "")
	// If ENVIRONMENT env var is not set, local configuration is taken and file is not found
	assert.Panics(t, func() { NewConfiguration() })
}
