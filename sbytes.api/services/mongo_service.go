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
	var client *mongo.Client
	var err error

	if client, err = connectToCluster(); err != nil {
		return nil, err
	}

	return &MongoService{
		client:     client,
		collection: client.Database(os.Getenv("DB_DATABASE")).Collection(os.Getenv("DB_COLLECTION")),
	}, nil
}

func connectToCluster() (*mongo.Client, error) {
	var client *mongo.Client
	var err error

	if client, err = mongo.NewClient(options.Client().ApplyURI(os.Getenv("DB_HOST"))); err != nil {
		return nil, err
	}

	if err = client.Connect(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}

func (receiver *MongoService) InsertTicket(document bson.M) (interface{}, error) {
	var documentInserted *mongo.InsertOneResult
	var err error

	if documentInserted, err = receiver.collection.InsertOne(context.Background(), document); err != nil {
		return nil, err
	}

	return documentInserted.InsertedID, nil
}

func (receiver *MongoService) UpdateTicket(ticketId string, document bson.D) error {

	filter := bson.M{"_id": ticketId}
	update := bson.M{"$set": document}

	if err := receiver.collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&document); err != nil {
		return err
	}

	return nil
}

func (receiver *MongoService) FindTicket(ticketId string) (bson.M, error) {
	var document bson.M
	objectId, _ := primitive.ObjectIDFromHex(ticketId)

	filter := bson.M{"_id": objectId}

	if err := receiver.collection.FindOne(context.Background(), filter).Decode(&document); err != nil {
		return bson.M{}, err
	}

	return document, nil
}
