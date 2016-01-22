package store

import (
	"domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	USERCOLLECTION_NAME = "users"
)


type UserDao struct {
}


func (dao *UserDao) Save(object domain.User) domain.User {
	return mongoTask(USERCOLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		existingUser := dao.Find(object.Username)
		if existingUser != nil {
			log.Printf("Can't insert user %v username already existing\n", object)
			panic("Can't insert user username already existing")
		}

		err := collection.Insert(object)
		if err != nil {
			log.Printf("Can't insert user %v due to error %v\n", object, err)
			panic(err)
		}
		return object
	}).(domain.User)
}

func (dao *UserDao) Update(object domain.User) domain.User {
	return mongoTask(USERCOLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		err := collection.Update(bson.M{"username": object.Username}, object)
		if err != nil {
			log.Printf("Can't update user %v due to error %v\n", object, err)
			panic(err)
		}
		return object
	}).(domain.User)
}

func (dao *UserDao) Delete(object domain.User) bool {
	return mongoTask(USERCOLLECTION_NAME, func(collection *mgo.Collection) interface{} {

		err := collection.Remove(bson.M{"username": object.Username})
		if err != nil {
			log.Printf("Can't delete user %v due to error %v\n", object, err)
			panic(err)
		}
		return true
	}).(bool)
}

func (dao *UserDao) Find(username string) *domain.User {
	value := mongoTask(USERCOLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		var search domain.User
		err := collection.Find(bson.M{"username": username}).One(&search)
		if err != nil {
			if err == mgo.ErrNotFound {
				return nil
			}
			log.Printf("Can find user with username %v due to error %v\n", username, err)
			panic(err)
		}

		return &search
	})
	if value == nil {
		return nil
	}
	return value.(*domain.User)
}

func (dao *UserDao) FindAll() []domain.User {
	return mongoTask(USERCOLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		var search []domain.User
		err := collection.Find(bson.M{}).All(&search)
		if err != nil {
			log.Printf("Can find all user due to error %v\n", err)
			panic(err)
		}

		return search
	}).([]domain.User)
}
