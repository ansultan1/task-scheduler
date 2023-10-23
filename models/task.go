package models

import (
	"github.com/fatih/structs"
	"time"
)

type Task struct {
	ID            string    `json:"id" bson:"_id" structs:"id" db:"id"`
	Name          string    `json:"name" bson:"name" structs:"name" db:"name"`
	Command       string    `json:"command" bson:"command" structs:"command" db:"command"`
	ScheduledTime time.Time `json:"scheduledTime" bson:"scheduled_time" structs:"scheduled_time" db:"scheduled_time"`
	Recurring     bool      `json:"recurring" bson:"recurring" structs:"recurring" db:"recurring"`
	TimeZone      string    `json:"timeZone" bson:"time_zone" structs:"time_zone" db:"time_zone"`
}

func (t *Task) Map() map[string]interface{} {
	return structs.Map(t)
}

// Names returns the field names of Task model
func (t *Task) Names() []string {
	fields := structs.Fields(t)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}
	return names
}
