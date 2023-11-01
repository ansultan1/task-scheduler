package service

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
	"task-scheduler/models"
)

var cronTasks = make(map[string]*cron.Cron)

func (s *Service) executeTask(task *models.Task) error {
	//  execute the task's
	cmd := exec.Command("sh", "-c", task.Command)
	return cmd.Run()
}

// AddOrUpdateTask adds or update task into database
func (s *Service) AddOrUpdateTask(task *models.Task) (string, error) {
	if task == nil {

		return "", errors.New("task is empty")
	}

	// loc extract the timezone enter by user
	loc, err := time.LoadLocation(task.TimeZone)
	if err != nil {

		return "", errors.New("time zone is incorrect")
	}

	// convert the scheduled time in utc to work with globally
	task.ScheduledTime = task.ScheduledTime.In(loc).UTC()

	if _, err := s.db.AddOrUpdateTask(task); err != nil {

		return "", errors.Wrap(err, "failed to add task")
	}

	// Generate a unique task name to identify by cron
	taskName := fmt.Sprintf("task_%s", task.ID)

	c := cron.New()

	cronTasks[taskName] = c

	// Schedule the task to run at the specified time
	schedule := fmt.Sprintf("%d %d %d %d %d %d", task.ScheduledTime.Second(), task.ScheduledTime.Minute(), task.ScheduledTime.Hour(), task.ScheduledTime.Weekday(), task.ScheduledTime.Month(), task.ScheduledTime.Year())
	c.AddFunc(schedule, func() {
		// Execute the task concurrently using a goroutine
		go func() {
			if err := s.executeTask(task); err != nil {
				log.Printf("Failed to execute task: %v", err)
			}
		}()
	})

	c.Start()

	return task.ID, nil
}

// GetTask gets task from database using the id
func (s *Service) GetTask(id string) (*models.Task, error) {
	task, err := s.db.GetTask(id)
	if err != nil {

		return nil, err
	}

	return task, nil
}

// ListTasks list all the tasks to the user
func (s *Service) ListTasks() ([]*models.Task, error) {
	tasks, err := s.db.ListTasks()
	if err != nil {

		return nil, err
	}

	return tasks, nil
}

// DeleteTask deletes task from database
func (s *Service) DeleteTask(id string) error {
	_, err := s.db.GetTask(id)
	if err != nil {

		return err
	}

	return s.db.DeleteTask(id)
}
