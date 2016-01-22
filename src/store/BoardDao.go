package store

import (
	"domain"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

const (
	BOARD_COLLECTION_NAME = "boards"
)

type BoardDao struct {
}

func (dao *BoardDao) Save(object domain.Board) domain.Board {
	return mongoTask(BOARD_COLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		object.Id = bson.NewObjectIdWithTime(time.Now())
		err := collection.Insert(object)
		if err != nil {
			log.Printf("[BoardDao][Save] Error during query execution : %v\n", err)
			panic(err)
		}
		return object
	}).(domain.Board)
}

func (dao *BoardDao) Update(object domain.Board) bool {
	return mongoTask(BOARD_COLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		err := collection.Update(bson.M{"id": object.Id}, object)
		if err != nil {
			if err == mgo.ErrNotFound {
				return false
			}
			log.Printf("[BoardDao][Update] Error during query execution : %v\n", err)
			panic(err)
		}
		return true
	}).(bool)
}

func (dao *BoardDao) Delete(object domain.Board) bool {
	return mongoTask(BOARD_COLLECTION_NAME, func(collection *mgo.Collection) interface{} {

		err := collection.Remove(bson.M{"id": object.Id})
		if err != nil {
			if err == mgo.ErrNotFound {
				return false
			}
			log.Printf("[BoardDao][Delete] Error during query execution : %v\n", err)
			panic(err)
		}
		return true
	}).(bool)
}

func (dao *BoardDao) Find(id bson.ObjectId) *domain.Board {
	value := mongoTask(BOARD_COLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		var search domain.Board
		err := collection.Find(bson.M{"id": id}).One(&search)
		if err != nil {
			if err == mgo.ErrNotFound {
				return nil
			}
			log.Printf("[BoardDao][Find] Error during query execution : %v\n", err)
			panic(err)
		}

		return &search
	})
	if value == nil {
		return nil
	}
	return value.(*domain.Board)
}

func (dao *BoardDao) FindAll() []domain.Board {
	return mongoTask(BOARD_COLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		var search []domain.Board
		err := collection.Find(bson.M{}).All(&search)
		if err != nil {
			log.Printf("[BoardDao][FindAll] Error during query execution : %v\n", err)
			panic(err)
		}

		return search
	}).([]domain.Board)
}

func (dao *BoardDao) FindAllBoardsForUser(user domain.User) []domain.Board {
	return mongoTask(BOARD_COLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		var search []domain.Board
		err := collection.Find(bson.M{"$or": []interface{}{
			bson.M{"ownerid": user.Username},
			bson.M{"authorizedid": user.Username},
		}}).All(&search)
		if err != nil {
			log.Printf("[BoardDao][FindAllBoardsForUser] Error during query execution : %v\n", err)
			panic(err)
		}

		return search
	}).([]domain.Board)
}

// ------------------------------------------------
// Column 
// ------------------------------------------------

func (dao *BoardDao) FindColumnForBoard(boardId bson.ObjectId, columnId bson.ObjectId) *domain.Column {
	value :=  mongoTask(BOARD_COLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		var search domain.Column
		aggregate := collection.Pipe([]bson.M{
				{"$match": bson.M{"id": boardId}},
				{"$unwind": "$columns"},
				{"$match": bson.M{"columns.id": columnId}},
				{"$project" : bson.M{"id":"$columns.id", "name":"$columns.name", "order":"$columns.order", "tickets":"$columns.tickets"}}})
		err := aggregate.One(&search)
		if err != nil {
			if err == mgo.ErrNotFound {
				return nil
			}
			log.Printf("[BoardDao][FindColumnForBoard] Error during query execution : %v\n", err)
			panic(err)
		}

		return &search
	})
	if value == nil {
		return nil
	}
	return value.(*domain.Column)
}


func (dao *BoardDao) UpdateColumn(boardId bson.ObjectId, object domain.Column) bool {
	return mongoTask(BOARD_COLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		err := collection.Update(
				bson.M{"id": boardId, "columns.id": object.Id}, 
				bson.M{"$set": bson.M{"columns.$" : object}})
		if err != nil {
			if err == mgo.ErrNotFound {
				return false
			}
			log.Printf("[BoardDao][UpdateColumn] Error during query execution : %v\n", err)
			panic(err)
		}
		return true
	}).(bool)
}

func (dao *BoardDao) DeleteColumn(boardId bson.ObjectId, object domain.Column) bool {
	return mongoTask(BOARD_COLLECTION_NAME, func(collection *mgo.Collection) interface{} {

		err := collection.Update(
			bson.M{"id": boardId, "columns.id": object.Id}, 
			bson.M{"$pull": bson.M{"columns": bson.M{"id": object.Id}}})
		if err != nil {
			if err == mgo.ErrNotFound {
				return false
			}
			log.Printf("[BoardDao][DeleteColumn] Error during query execution : %v\n", err)
			panic(err)
		}
		return true
	}).(bool)
}

// ------------------------------------------------
// Ticket 
// ------------------------------------------------

func (dao *BoardDao) FindTicketForBoardAndColumn(boardId bson.ObjectId, columnId bson.ObjectId, ticketId bson.ObjectId) *domain.Ticket {
	value :=  mongoTask(BOARD_COLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		var search domain.Ticket
		aggregate := collection.Pipe([]bson.M{
				{"$match": bson.M{"id": boardId, "columns.id": columnId}},
				{"$unwind": "$columns"},
				{"$unwind": "$columns.tickets"},
				{"$match": bson.M{"columns.tickets.id": ticketId}},
				{"$project" : bson.M{"id":"$columns.tickets.id", "name":"$columns.tickets.name", "description":"$columns.tickets.description", "opendate":"$columns.tickets.opendate", "affecteduserid":"$columns.tickets.affecteduserid" }}})
		err := aggregate.One(&search)
		if err != nil {
			if err == mgo.ErrNotFound {
				return nil
			}
			log.Printf("[BoardDao][FindTicketForBoardAndColumn] Error during query execution : %v\n", err)
			panic(err)
		}

		return &search
	})
	if value == nil {
		return nil
	}
	return value.(*domain.Ticket)
}

// TODO : Change solution to use direct update, in mermory solution to be quick
func (dao *BoardDao) UpdateTicket(boardId bson.ObjectId, columnId bson.ObjectId, object domain.Ticket) bool {
	return mongoTask(BOARD_COLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		column := dao.FindColumnForBoard(boardId, columnId)
		found := false
		for i:=0; i<len(column.Tickets); i++ {
			if column.Tickets[i].Id == object.Id {
				column.Tickets[i] = object
				found = true
				break
			}
		}
		if !found {
			return false
		}
		dao.UpdateColumn(boardId, *column)
		return true
	}).(bool)
}

// TODO : Change solution to use direct update, in mermory solution to be quick
func (dao *BoardDao) DeleteTicket(boardId bson.ObjectId, column domain.Column, object domain.Ticket) bool {
	return mongoTask(BOARD_COLLECTION_NAME, func(collection *mgo.Collection) interface{} {
		foundIndex := -1
		for i:=0; i<len(column.Tickets); i++ {
			if column.Tickets[i].Id == object.Id {
				foundIndex = i
				break
			}
		}
		if foundIndex < 0 {
			return false
		}
		column.Tickets = append(column.Tickets[:foundIndex], column.Tickets[foundIndex+1:]...)
		dao.UpdateColumn(boardId, column)
		return true
	}).(bool)
}