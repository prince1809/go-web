package data

import (
	"github.com/prince1809/go-web/TaskManager/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type NoteRepository struct {
	C *mgo.Collection
}

func (r *NoteRepository) Create(note *models.TaskNote) error {
	objId := bson.NewObjectId()
	note.Id = objId
	note.CreatedOn = time.Now()
	err := r.C.Insert(&note)
	return err
}

func (r *NoteRepository) Update(note *models.TaskNote) error {
	err := r.C.Update(bson.M{"_id": note.Id},
		bson.M{"$SET": bson.M{
			"description": note.Description,
		}})
	return err
}

func (r *NoteRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (r *NoteRepository) GetByTask(id string) []models.TaskNote {
	var notes []models.TaskNote
	taskid := bson.ObjectIdHex(id)
	iter := r.C.Find(bson.M{"taskid": taskid}).Iter()
	result := models.TaskNote{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}

func (r *NoteRepository) GetAll() []models.TaskNote {
	var notes []models.TaskNote
	iter := r.C.Find(nil).Iter()
	result := models.TaskNote{}
	for iter.Next(&result) {
		notes = append(notes, result)
	}
	return notes
}

func (r *NoteRepository) GetById(id string) (note models.TaskNote, err error) {
	err = r.C.FindId(bson.ObjectIdHex(id)).One(&note)
	return note, err
}
