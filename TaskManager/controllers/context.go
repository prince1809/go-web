package controllers

import (
	"github.com/prince1809/go-web/TaskManager/common"
	"gopkg.in/mgo.v2"
)

type Context struct {
	MongoSession *mgo.Session
	User         string
}

// Close mgo.session
func (c *Context) close() {
	c.MongoSession.Close()
}

//DbCollection returns mgo.collection for the given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

// NewContext creates a new context object for each HTTP request
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}
