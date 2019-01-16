package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

type Task struct {
	Description string
	Due         time.Time
}

type Category struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Description string
	Tasks       []Task
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	//get collection
	c := session.DB("taskdb").C("categories")
	// Embedding child collection
	doc := Category{
		bson.NewObjectId(),
		"Open-Source",
		"Task for open source projects",
		[]Task{
			{"Create project in mgo", time.Date(2018, time.January, 16, 15, 0, 0, 0, time.UTC)},
			{"Create REST API", time.Date(2019, time.January, 10, 10, 1, 0, 0, time.UTC)},
		},
	}

	// insert a  category object with embedded tasks
	err = c.Insert(&doc)
	if err != nil {
		log.Fatal(err)
	}
}
