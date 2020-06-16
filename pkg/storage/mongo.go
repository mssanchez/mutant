package storage

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"mutant/pkg/config"
	"time"
)

type MongoClient struct {
	DabaseName           string
	CollectionName       string
	client               *mongo.Client
	log                  *logrus.Logger
	disconnectTimeoutInS int
	cancelFunc           context.CancelFunc
}

func newMongoClient(client *mongo.Client, config config.Configuration, log *logrus.Logger) *MongoClient {
	// Connect the mongo client to the MongoDB server
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	// Ping MongoDB
	pingCtx, cancelPingFunc := context.WithCancel(context.Background())
	if err = client.Ping(pingCtx, nil); err != nil {
		panic("could not ping to mongo db service: " + err.Error())
	}
	cancelPingFunc()

	return &MongoClient{
		DabaseName:           config.Mongodb.DatabaseName,
		CollectionName:       config.Mongodb.CollectionName,
		client:               client,
		cancelFunc:           cancelFunc,
		disconnectTimeoutInS: config.Mongodb.DisconnectTimeoutInSeconds,
		log:                  log,
	}
}
