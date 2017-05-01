package tasks

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongoStore struct {
	Session        *mgo.Session
	DatabaseName   string
	CollectionName string
}

func (ms *MongoStore) Insert(newtask *NewTask) (*Task, error) {
	t := newtask.ToTask()
	t.ID = bson.NewObjectId()
	// assuming that the user is connected to the database
	// collection
	err := ms.Session.DB(ms.DatabaseName).C(ms.CollectionName).Insert(t)
	return t, err // returning it to let the caller handle it
}

func (ms *MongoStore) Get(ID interface{}) (*Task, error) {
	// looks at ID parameter, if ID is a string, then ok gets set to true
	// and sID is set to that ID
	if sID, ok := ID.(string); ok {
		// bson : binary id that identifies the object
		ID = bson.ObjectIdHex(sID)
	}
	task := &Task{}
	// Passing a struct pointer for mongo to fill out the fields
	// Because mongo is NoSQL so it doesn't know what is stored
	err := ms.Session.DB(ms.DatabaseName).C(ms.CollectionName).FindId(ID).One(task)
	return t, err
}

func (ms *MongoStore) GetAll() ([]*Task, error) {
	tasks := []*Task{}
	// using nil for no filtering, just return all documents
	ms.Session.DB(ms.DatabaseName).C(ms.CollectionName).Find(nil).All(&tasks)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ms *MongoStore) Update(task *Task) error {
	task.ModifiedAt = time.Now()

	// reference to the collection
	col := ms.Session.DB(ms.DatabaseName).C(ms.CollectionName)
	// weird syntax from Mongo Update
	updates := bson.M{"$set": bson.M{"completed": task.Complete, "modifiedAt": task.ModifiedAt}}
	col.UpdateId(task.ID, updated)
}
