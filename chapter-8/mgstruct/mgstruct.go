package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Category struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Description string
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	// Reads may not be entirely up-to-date, but they will always see the
	// history of changes moving foreward, the data read will be consistent
	// across sequential queries in the same session, and modification made
	// within the session will be observed in following queries (read-your-writes).
	session.SetMode(mgo.Monotonic, true)

	//get collection
	c := session.DB("taskdb").C("categories")

	doc := Category{
		bson.NewObjectId(),
		"Open Source",
		"Task for open source projects",
	}

	//insert a category object
	err = c.Insert(&doc)
	if err != nil {
		log.Fatal(err)
	}

	// insert two category objects
	err = c.Insert(&Category{bson.NewObjectId(), "R & D", "R & D Tasks"},
		&Category{bson.NewObjectId(), "Project", "Project Tasks"})
	var count int
	count, err = c.Count()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%d records inserted", count)
	}
}
