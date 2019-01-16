package data

import (
	"github.com/prince1809/go-web/TaskManager/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type userRepository struct {
	c *mgo.Collection
}


func (r *userRepository) createUser(user *models.User) error {
	obj_id := bson.NewObjectId()
	user.Id = obj_id
	hpass, err := bcrypt
}
