package storage

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mutant/pkg/config"
	"time"
)

// MutantDoc represents a request made by the user and the result (if it's a mutant or not)
type MutantDoc struct {
	Dna      []string `bson:"dna" json:"dna"`
	IsMutant bool     `bson:"is_mutant" json:"is_mutant"`
}

// MutantStorage handles all operations to the database related to saving and counting mutants documents
type MutantStorage interface {
	Save(val *MutantDoc) error
	Count(isMutant bool) (int64, error)
	Shutdown()
}

type mongoClient struct {
	DabaseName           string
	CollectionName       string
	client               *mongo.Client
	log                  *logrus.Logger
	disconnectTimeoutInS int
	cancelFunc           context.CancelFunc
}

// NewMutantsStorage builds a MutantStorage
func NewMutantsStorage(config config.Configuration, log *logrus.Logger) MutantStorage {
	log.Info("init storage package...")

	var client *mongo.Client
	var err error

	log.Infof("storage client with %s environment", config.Environment)

	connectionURI := config.Mongodb.Url
	if config.Mongodb.Password != "" {
		connectionURI = fmt.Sprintf(config.Mongodb.Url, config.Mongodb.Password)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Errorf("fatal error: %v", err)
		panic(err)
	}

	pingCtx, cancelPingFunc := context.WithTimeout(context.Background(), 20*time.Second)
	if err = client.Ping(pingCtx, nil); err != nil {
		panic("could not ping to mongo db service: " + err.Error())
	}
	cancelPingFunc()

	return &mongoClient{
		DabaseName:           config.Mongodb.DatabaseName,
		CollectionName:       config.Mongodb.CollectionName,
		client:               client,
		cancelFunc:           cancel,
		disconnectTimeoutInS: config.Mongodb.DisconnectTimeoutInSeconds,
		log:                  log,
	}
}

func (mc *mongoClient) Count(isMutant bool) (int64, error) {
	opts := options.Count().SetMaxTime(30 * time.Second)
	filter := bson.M{"is_mutant": isMutant}

	return mc.collection().CountDocuments(context.TODO(), filter, opts)
}

func (mc *mongoClient) Save(doc *MutantDoc) error {
	mc.log.Infof("Storing mutant doc %v", doc)

	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"dna", bson.A{doc.Dna}}}
	update := bson.D{{"$set", bson.D{{"is_mutant", doc.IsMutant}}}}

	_, err := mc.collection().UpdateOne(ctx, filter, update, opts)
	return err
}

func (mc *mongoClient) Shutdown() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(mc.disconnectTimeoutInS)*time.Second)
	if err := mc.client.Disconnect(ctx); err != nil {
		mc.log.Error(err.Error())
	}

	cancel()
	mc.cancelFunc()
}

func (mc *mongoClient) collection() *mongo.Collection {
	return mc.client.Database(mc.DabaseName).Collection(mc.CollectionName)
}
