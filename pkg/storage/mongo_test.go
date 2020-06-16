package storage
/*
import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"sam/pkg/config"
	"testing"
	"time"
)

func TestSaveMutant(t *testing.T) {

	conf := config.TestConfig()
	containerId, client, err := startContainer(conf)
	assert.NoError(t, err)
	defer stopContainer(containerId)

	db := newMongoClient(client, conf)

	id := uuid.New()
	err = db.Save(&DeploymentDoc{
		DeploymentId: id.String(),
		Name:         "test-deployment",
		Service:      "test",
		Environment:  "local",
		Version:      "0.0.1-local-12345",
		Tag:          "0.0.1",
		Status:       "NEW",
		LastChange:   time.Time{},
	})
	assert.NoError(t, err)

	byId := db.FindDeploymentById(context.Background(), id)
	assert.True(t, "NEW" == byId.Status)
}

func TestUpdateDocument(t *testing.T) {

	conf := config.TestConfig()
	containerId, client, err := startContainer(conf)
	assert.NoError(t, err)
	defer stopContainer(containerId)

	db := newMongoClient(client, conf)

	id := uuid.New()
	err = db.Save(&DeploymentDoc{
		DeploymentId: id.String(),
		Name:         "test-deployment",
		Service:      "test",
		Environment:  "local",
		Version:      "0.0.1-local-12345",
		Tag:          "0.0.1",
		Status:       "NEW",
		LastChange:   time.Time{},
	})
	assert.NoError(t, err)

	doc := db.FindDeploymentById(context.Background(), id)
	assert.True(t, "NEW" == doc.Status)

	err = db.Update(&DeploymentDoc{
		DeploymentId: id.String(),
		Name:         "test-deployment",
		Service:      "test",
		Version:      "0.0.1-local-12345",
		Status:       "STARTING",
		LastChange:   time.Time{},
	})
	assert.NoError(t, err)

	updatedDoc := db.FindDeploymentById(context.Background(), id)
	assert.True(t, "STARTING" == updatedDoc.Status)
	assert.True(t, "local" == updatedDoc.Environment)
	assert.True(t, "0.0.1" == updatedDoc.Tag)

}
*/




