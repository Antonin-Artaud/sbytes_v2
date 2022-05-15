package services

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	database   = "sbytes"
	collection = "tickets"
	mongoDbUri = "mongodb://localhost:27017"
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
		collection: client.Database(database).Collection(collection),
	}, nil
}

func connectToCluster() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDbUri))

	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (receiver *MongoService) InsertTicket(document bson.D) (bson.D, error) {
	_, err := receiver.collection.InsertOne(context.Background(), document)

	if err != nil {
		return bson.D{{
			"error", err.Error(),
		}}, err
	}

	return bson.D{{
		"status", "document successfully inserted",
	}}, nil
}

func (receiver *MongoService) UpdateTicket(ticketId uuid.UUID) {
	singleResult := receiver.collection.FindOne(context.Background(), bson.D{{"ticket", bson.D{{"id", ticketId}}}})

	print(singleResult)
}

func (receiver *MongoService) FindTicket(ticketId string) bson.D {
	filterCursor, err := receiver.collection.Find(context.Background(), bson.D{{"ticket.guid", ticketId}})

	if err != nil {
		return bson.D{{
			"error", err.Error(),
		}}
	}

	var document bson.D

	for filterCursor.Next(context.Background()) {
		err := filterCursor.Decode(&document)

		if err != nil {
			return bson.D{{
				"error", err.Error(),
			}}
		}
	}
	return document
}

func (receiver *MongoService) RemoveTicket(ticketId uuid.UUID) {

}
