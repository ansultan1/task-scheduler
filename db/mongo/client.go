package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"task-scheduler/config"
	"task-scheduler/db"
	domainErr "task-scheduler/errors"
	"task-scheduler/models"
)

const (
	taskCollection = "task"
)

func init() {
	db.Register("mongo", NewClient)
}

type client struct {
	conn *mongo.Client
}

// NewClient initializes a mysql database connection
func NewClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)

	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {

		return nil, errors.Wrap(err, "failed to connect to db")
	}

	return &client{conn: cli}, nil
}

// AddOrUpdateTask create or update task in the mongodb
func (c *client) AddOrUpdateTask(task *models.Task) (string, error) {
	if task == nil {

		return "", errors.New("task cannot be empty")
	}

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(taskCollection)
	if _, err := collection.UpdateOne(context.TODO(), bson.D{{"_id", task.ID}}, bson.D{{"$set", task}}, options.Update().SetUpsert(true)); err != nil {

		return "", errors.Wrap(err, "failed to add task")
	}

	return task.ID, nil
}

// GetTask get the task from db based on ID
func (c *client) GetTask(id string) (*models.Task, error) {
	var tsk *models.Task

	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(taskCollection)
	if err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&tsk); err != nil {
		if err == mongo.ErrNoDocuments {

			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found", id))
		}

		return nil, err
	}

	return tsk, nil
}

// ListTasks Fetching the data of all teachers present in the database
func (c *client) ListTasks() ([]*models.Task, error) {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(taskCollection)
	var tasks []*models.Task

	tasksCursor, err := collection.Find(context.TODO(), bson.D{}, options.Find().SetSort(bson.D{{"_id", 1}}))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {

			return nil, errors.Wrap(err, "No Teachers' Data Found!")
		}

		return nil, err
	}

	if err := tasksCursor.All(context.TODO(), &tasks); err != nil {

		return nil, err
	}

	return tasks, nil
}

// DeleteTask remover task from the database
func (c *client) DeleteTask(id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(taskCollection)
	if _, err := collection.DeleteOne(context.TODO(), bson.M{"_id": id}); err != nil {

		return errors.Wrap(err, "failed to delete task")
	}

	return nil
}
