package mysql

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/ansultan1/task-scheduler/db"
	"github.com/ansultan1/task-scheduler/models"
)

func Test_client_AddOrUpdateTask(t *testing.T) {
	setDBENV()
	type args struct {
		task *models.Task
		id   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// test cases
		{
			name:    "success - add task in db",
			args:    args{task: &models.Task{ID: "9", Name: "echo task", Command: "echo", ScheduledTime: time.Date(2023, 10, 23, 21, 51, 22, 0, time.UTC), Recurring: false, TimeZone: "UTC"}},
			wantErr: false,
		},
		{
			name:    "fail - add invalid task in db",
			args:    args{task: &models.Task{ID: "4", Name: "echo task", Command: "echo khan world", ScheduledTime: time.Date(2023, 10, 23, 21, 51, 22, 0, time.UTC), Recurring: false, TimeZone: "UTC"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(db.Option{})
			_, err := c.AddOrUpdateTask(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_client_GetTask(t *testing.T) {
	setDBENV()
	c, _ := NewClient(db.Option{})
	task := &models.Task{

		Name:          "AddTask",
		Command:       "echo hello world",
		ScheduledTime: time.Now(),
		Recurring:     false,
		TimeZone:      "UTC",
	}
	_, err := c.AddOrUpdateTask(task)
	if err != nil {
		t.Errorf("Failed to save task: %v", err)
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - get task from db",
			args: args{id: task.ID},

			wantErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := c.GetTask(testCase.args.id)
			if (err != nil) != testCase.wantErr {
				t.Errorf("GetTask() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}

			fmt.Printf("got: %#v\n", got)

		})
	}
}

func Test_client_ListTask(t *testing.T) {
	setDBENV()
	type args struct {
		task *models.Task
	}
	c, _ := NewClient(db.Option{})

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - List Tasks from db",

			wantErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {

			got, err := c.ListTasks()
			if (err != nil) != testCase.wantErr {
				t.Errorf("ListTasks() error = %v, wantErr %v", err, testCase.wantErr)
			}

			for _, task := range got {
				fmt.Printf("Task ID: %s\n", task.ID)
				fmt.Printf("Task Name: %s\n", task.Name)
			}
		})
	}

}

func Test_client_DeleteTask(t *testing.T) {
	setDBENV()
	c, _ := NewClient(db.Option{})
	task := &models.Task{Name: "AddTask", Command: "echo hello world", ScheduledTime: time.Date(2023, 10, 23, 21, 51, 22, 0, time.UTC), Recurring: false, TimeZone: "UTC"}
	taskID, _ := c.AddOrUpdateTask(task)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// test cases
		{
			name:    "success - delete task from db",
			args:    args{id: taskID},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteTask(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// setDBENV has connection for DB
func setDBENV() {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "test")
	os.Setenv("DB_NAME", "task_schedular")
}
