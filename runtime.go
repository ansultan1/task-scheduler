package runtime

import (
	"github.com/ansultan1/task-scheduler/db"
	"github.com/ansultan1/task-scheduler/db/mongo"
	"github.com/ansultan1/task-scheduler/service"
)

type Runtime struct {
	svc *service.Service
}

// NewRuntime creates a new database service
func NewRuntime() (*Runtime, error) {
	store, err := mongo.NewClient(db.Option{})
	if err != nil {
		return nil, err
	}
	return &Runtime{svc: service.NewService(store)}, err
}

// Service returns pointer to our service variable
func (r Runtime) Service() *service.Service {
	return r.svc
}
