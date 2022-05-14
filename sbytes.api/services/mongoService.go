package services

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	client *mongo.Client
}

func NewMongoService() (*MongoService, error) {

	cluster, err := connectToCluster()

	if err != nil {
		return nil, err
	}

	return &MongoService{
		client: cluster,
	}, nil
}

func connectToCluster() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (receiver *MongoService) listDatabaseNames() {
	databaseNames, err := receiver.client.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		return
	}

	fmt.Println(databaseNames)
}

func (receiver *MongoService) dispose() {
	receiver.dispose()
}
