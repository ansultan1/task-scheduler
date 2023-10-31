package mysql

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"task-scheduler/config"
	"task-scheduler/db"
	domainErr "task-scheduler/errors"
	"task-scheduler/models"
)

const (
	taskTableName = "task"
)

func init() {
	db.Register("mysql", NewClient)
}

// The first implementation.
type client struct {
	db *sqlx.DB
}

// FormatDSN generates a Data Source Name (DSN) for MySQL from configuration
func FormatDSN() string {

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		viper.GetString(config.DbUser),
		viper.GetString(config.DbPass),
		viper.GetString(config.DbHost),
		viper.GetString(config.DbPort),
		viper.GetString(config.DbName),
	)
}

// NewClient initializes a mysql database connection
func NewClient(conf db.Option) (db.DataStore, error) {
	log().Info("initializing mysql connection: " + FormatDSN())

	cli, err := sqlx.Connect("mysql", FormatDSN())
	if err != nil {

		return nil, errors.Wrap(err, "failed to connect to db")
	}

	return &client{db: cli}, nil
}

// AddOrUpdateTask allows the user to store a task in the database
func (c *client) AddOrUpdateTask(task *models.Task) (string, error) {
	if task == nil {

		return "", errors.New("task  is empty")
	}

	names := task.Names()

	query := fmt.Sprintf(`REPLACE INTO %s (%s) VALUES(%s)`, taskTableName, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ","))
	if _, err := c.db.NamedExec(query, task); err != nil {

		return "", errors.Wrap(err, "failed to add task")
	}

	return task.ID, nil
}

// GetTask gets the task from database based on id
func (c *client) GetTask(id string) (*models.Task, error) {
	var task models.Task
	if err := c.db.Get(&task, fmt.Sprintf(`SELECT * FROM %s WHERE id = '%s'`, taskTableName, id)); err != nil {
		if err == sql.ErrNoRows {

			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found", id))
		}

		return nil, err
	}

	return &task, nil
}

// ListTasks lists the tasks from mysql db
func (c *client) ListTasks() ([]*models.Task, error) {
	// Prepare the SQL query to retrieve tasks
	query := "SELECT id, name, command, scheduled_time, recurring, time_zone FROM task"

	rows, err := c.db.Query(query)
	if err != nil {

		return nil, err
	}

	defer rows.Close()

	tasks := []*models.Task{}
	for rows.Next() {
		task := &models.Task{}
		err := rows.Scan(&task.ID, &task.Name, &task.Command, &task.ScheduledTime, &task.Recurring, &task.TimeZone)
		if err != nil {

			return nil, err
		}

		tasks = append(tasks, task)
	}

	if len(tasks) == 0 {

		return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: not found"))
	}

	if err = rows.Err(); err != nil {

		return nil, err
	}

	return tasks, nil
}

// DeleteTask removes a task from db
func (c *client) DeleteTask(id string) error {
	if _, err := c.db.Query(fmt.Sprintf(`DELETE FROM %s WHERE id= '%s'`, taskTableName, id)); err != nil {

		return errors.Wrap(err, "failed to delete task")
	}

	return nil
}

func mkPlaceHolder(names []string, prefix string, formatName func(name, prefix string) string) []string {
	ph := make([]string, len(names))
	for i, name := range names {
		ph[i] = formatName(name, prefix)
	}

	return ph
}
