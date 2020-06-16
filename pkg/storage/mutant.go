package storage

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mutant/pkg/config"
	"os"
	"time"
)

type MutantDoc struct {
	Dna      []string `bson:"dna" json:"dna"`
	IsMutant bool     `bson:"is_mutant" json:"is_mutant"`
}

type MutantStorage interface {
	Save(val *MutantDoc) error
	Count(isMutant bool) (int64, error)
	Shutdown()
}

func NewMutantsStorage(config config.Configuration, log *logrus.Logger) MutantStorage {
	log.Info("init storage package...")

	var client *mongo.Client
	var err error

	log.Infof("storage client with %s environment", config.Environment)

	password := os.Getenv("MUTANT-DOCUMENTDB-PASS")
	connectionURI := config.Mongodb.Url
	if password != "" {
		connectionURI = fmt.Sprintf(config.Mongodb.Url, password)
	}

	client, err = mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Errorf("fatal error: %v", err)
		panic(err)
	}

	return newMongoClient(client, config, log)
}

func (mc *MongoClient) Count(isMutant bool) (int64, error) {
	opts := options.Count().SetMaxTime(30 * time.Second)
	filter := bson.M{"is_mutant": isMutant}

	return mc.collection().CountDocuments(context.TODO(), filter, opts)
}

func (mc *MongoClient) Save(doc *MutantDoc) error {
	mc.log.Infof("Storing mutant doc %v", doc)

	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"dna", bson.A{doc.Dna}}}
	update := bson.D{{"$set", bson.D{{"is_mutant", doc.IsMutant}}}}

	_, err := mc.collection().UpdateOne(ctx, filter, update, opts)
	return err
}

func (mc *MongoClient) Shutdown() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(mc.disconnectTimeoutInS)*time.Second)
	if err := mc.client.Disconnect(ctx); err != nil {
		mc.log.Error(err.Error())
	}

	cancel()
	mc.cancelFunc()
}

func (mc *MongoClient) collection() *mongo.Collection {
	return mc.client.Database(mc.DabaseName).Collection(mc.CollectionName)
}
