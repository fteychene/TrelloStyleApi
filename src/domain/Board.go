package domain

import (
	"gopkg.in/mgo.v2/bson"
)

type Board struct {
	Id bson.ObjectId
	Name string
	Columns []Column
}