package service

import (
	"task-scheduler/db"
)

// Service initializes our database instance
type Service struct {
	db db.DataStore
}

// NewService creates a connection to our database
func NewService(ds db.DataStore) *Service {
	return &Service{db: ds}
}
