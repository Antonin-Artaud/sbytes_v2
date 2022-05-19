package services

import (
	"log"
	"sync"
)

var (
	lock     = &sync.Mutex{}
	instance *singleton
)

type singleton struct {
	MongoDb *MongoService
}

func GetInstance() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &singleton{
				MongoDb: nil,
			}
		}
	}

	return instance
}

func (receiver *singleton) InitiateDbConnection() {
	mongoClient, err := NewMongoService()

	if err != nil {
		log.Fatalln(err)
	}

	instance = &singleton{
		MongoDb: mongoClient,
	}
}
