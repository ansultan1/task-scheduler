package models

import (
	"reflect"
	"testing"
	"time"
)

func TestTask_Map(t *testing.T) {
	type fields struct {
		ID            string
		Name          string
		Command       string
		ScheduledTime time.Time
		Recurring     bool
		TimeZone      string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success - task struct to map",
			fields: fields{
				ID:            "1",
				Name:          "Sample Task",
				Command:       "echo 'Hello, World'",
				ScheduledTime: time.Date(2023, 10, 23, 21, 51, 22, 0, time.UTC),
				Recurring:     false,
				TimeZone:      "UTC",
			},
			want: map[string]interface{}{
				"id":             "1",
				"name":           "Sample Task",
				"command":        "echo 'Hello, World'",
				"scheduled_time": time.Date(2023, 10, 23, 21, 51, 22, 0, time.UTC),
				"recurring":      false,
				"time_zone":      "UTC",
			},
		},
		{
			name: "success - task struct to map with fewer fields",
			fields: fields{
				Name:      "Sample Task",
				Recurring: true,
				TimeZone:  "UTC",
			},
			want: map[string]interface{}{
				"id":             "",
				"name":           "Sample Task",
				"command":        "",
				"scheduled_time": time.Time{},
				"recurring":      true,
				"time_zone":      "UTC",
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &Task{
				ID:            tt.fields.ID,
				Name:          tt.fields.Name,
				Command:       tt.fields.Command,
				ScheduledTime: tt.fields.ScheduledTime,
				Recurring:     tt.fields.Recurring,
				TimeZone:      tt.fields.TimeZone,
			}
			if got := task.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_Names(t *testing.T) {
	type fields struct {
		ID            string
		Name          string
		Command       string
		ScheduledTime time.Time
		Recurring     bool
		TimeZone      string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success - names of task struct fields",
			fields: fields{
				ID:            "1",
				Name:          "Sample Task",
				Command:       "echo 'Hello, World'",
				ScheduledTime: time.Now(),
				Recurring:     false,
				TimeZone:      "UTC",
			},
			want: []string{"id", "name", "command", "scheduled_time", "recurring", "time_zone"},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &Task{
				ID:            tt.fields.ID,
				Name:          tt.fields.Name,
				Command:       tt.fields.Command,
				ScheduledTime: tt.fields.ScheduledTime,
				Recurring:     tt.fields.Recurring,
				TimeZone:      tt.fields.TimeZone,
			}
			if got := task.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
