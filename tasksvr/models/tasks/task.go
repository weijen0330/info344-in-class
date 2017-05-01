package tasks

import (
	"fmt"
	"time"
)

//NewTask represents a new task posted to the server
// We make the function and field name uppercase so that Go
// automatically exports them
type NewTask struct {
	//TODO: fill out fields
	Title string   `json: "title"`
	Tags  []string `json: "tags"`
}

//Task represents a task stored in the database
type Task struct {
	//TODO: fill out fields
	ID         interface{} `json:"id" bson: "_id"`
	Title      string      `json:"title"`
	Tags       []string    `json:"tags"`
	CreatedAt  time.Time   `json:"createdAt"`
	ModifiedAt time.Time   `json:"modifiedAt"`
	COmplete   bool        `json:"complete"`
}

//Validate will validate the NewTask
// Think this way: the struct upon which the validate function is called
func (nt *NewTask) Validate() error {
	//Title field must be non-zero length
	if len(nt.Title) == 0 {
		return fmt.Errorf("title must be something")
	}
	return nil
}

//ToTask converts a NewTask to a Task
func (nt *NewTask) ToTask() *Task {
	t := &Task{
		Title:     nt.Title,
		Tags:      nt.Tags,
		CreatedAt: time.Now(),
		// Must have a comma at the end. Easier to copy and paste for objects.
		ModifiedAt: time.Now(),
	}

	return t //pointer to a task
}
