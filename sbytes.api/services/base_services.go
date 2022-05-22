package services

import (
	"sync"
)

var (
	lock     = &sync.Mutex{}
	instance *baseService
)

type baseService struct {
	MongoDb *MongoService
}

func GetService() *baseService {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &baseService{
				MongoDb: nil,
			}
		}
	}

	return instance
}

func (receiver *baseService) InitiateDbConnection() error {
	var mongoClient *MongoService
	var err error

	if mongoClient, err = NewMongoService(); err != nil {
		return err
	}

	instance = &baseService{
		MongoDb: mongoClient,
	}

	return nil
}
