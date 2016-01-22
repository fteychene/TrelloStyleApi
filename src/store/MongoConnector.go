package store

import (
	"gopkg.in/mgo.v2"
	"log"
)

type MongoTask func(collection *mgo.Collection) (interface{})

func mongoTask(collectionName string, task MongoTask) interface{} {
	// TODO : Set mongodb url as env property
	uri := "mongodb://root:root@127.0.0.1"

	sess, err := mgo.Dial(uri)
	if err != nil {
		log.Printf("[MongoTak] Can't connect to mongo, go error %v\n", err)
		panic(err)
	}
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})
	
	return task(sess.DB("trello").C(collectionName))
	
}
