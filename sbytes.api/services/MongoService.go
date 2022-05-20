package services

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

type MongoService struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoService() (*MongoService, error) {

	client, err := connectToCluster()

	if err != nil {
		return nil, err
	}

	return &MongoService{
		client:     client,
		collection: client.Database(os.Getenv("DB_DATABASE")).Collection(os.Getenv("DB_COLLECTION")),
	}, nil
}

func connectToCluster() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_HOST")))

	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (receiver *MongoService) InsertTicket(document bson.M) (interface{}, error) {
	documentInserted, err := receiver.collection.InsertOne(context.Background(), document)

	if err != nil {
		return nil, err
	}

	return documentInserted.InsertedID, nil
}

func (receiver *MongoService) UpdateTicket(ticketId uuid.UUID) {
	singleResult := receiver.collection.FindOne(context.Background(), bson.D{{"ticket", bson.D{{"id", ticketId}}}})

	print(singleResult)
}

func (receiver *MongoService) FindTicket(ticketId string) bson.E {
	filterCursor, err := receiver.collection.Find(context.Background(), bson.E{Key: "_id", Value: ticketId})

	if err != nil {
		return bson.E{
			Key: "error", Value: err.Error(),
		}
	}

	var document bson.E

	for filterCursor.Next(context.Background()) {
		err := filterCursor.Decode(&document)

		if err != nil {
			return bson.E{
				Key: "error", Value: err.Error(),
			}
		}
	}

	return document
}