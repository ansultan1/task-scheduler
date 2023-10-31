package db

import (
	"log"
	"task-scheduler/models"
)

type DataStore interface {
	AddOrUpdateTask(task *models.Task) (string, error)
	GetTask(id string) (*models.Task, error)
	ListTasks() ([]*models.Task, error)
	DeleteTask(id string) error
	//UpdateTask(task *models.Task) error
}

// Option holds configuration for data store clients
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)

		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)

		return
	}
	datastoreFactories[name] = factory
}
