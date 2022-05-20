package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (receiver *MongoService) UpdateTicket(ticketId string, document bson.D) error {
	err := receiver.collection.FindOneAndUpdate(
		context.Background(),
		bson.M{"_id": ticketId},
		bson.M{"$set": document},
	).Decode(&document)

	if err != nil {
		return err
	}

	return nil
}

func (receiver *MongoService) FindTicket(ticketId string) (bson.M, error) {
	var document bson.M
	objectId, _ := primitive.ObjectIDFromHex(ticketId)

	err := receiver.collection.FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&document)

	if err != nil {
		return bson.M{}, err
	}

	return document, nil
}
