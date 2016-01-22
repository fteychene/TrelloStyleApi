package controller

import (
	"domain"
	"store"
	"log"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ColumnController struct {
	BoardDao *store.BoardDao
}

func (controller *ColumnController) GetAllColumnsForBoard(parameters map[string][]string) (int, interface{}) {
	log.Printf("[ColumnController][GetColumnsForBoard] Begin %v\n", parameters)
	board := controller.BoardDao.Find(bson.ObjectIdHex(parameters["boardId"][0]))
	log.Printf("[ColumnController][GetColumnsForBoard] End\n") 
	if board == nil {
		return 200, []domain.Column{}
	}
	return 200, board.Columns
}

func (controller *ColumnController) Get(parameters map[string][]string) (int, interface{}) {
	log.Printf("[ColumnController][Get] Begin %v\n", parameters)
	column := controller.BoardDao.FindColumnForBoard(bson.ObjectIdHex(parameters["boardId"][0]), bson.ObjectIdHex(parameters["columnId"][0]))
	log.Printf("[ColumnController][Get] End\n") 
	return 200, column
}

func (controller *ColumnController) Create(parameters map[string][]string) (int, interface{}) {
	log.Printf("[ColumnController][Create] Begin %v\n", parameters)
	validArguments, invalidMessage := checkParameters(parameters, []string{"boardId", "column"})
	if !validArguments {
		return 412, invalidMessage
	}
	
	board := controller.BoardDao.Find(bson.ObjectIdHex(parameters["boardId"][0]))
	if board == nil {
		return 412, "Unexisting board "+parameters["boardId"][0]
	}
	column := domain.Column{}
	err := json.Unmarshal([]byte(parameters["column"][0]), &column)
	column.Id = bson.NewObjectIdWithTime(time.Now())
	if err != nil {
		return 412, "Error parsin : column"
	}
	board.Columns = append(board.Columns, column)
	controller.BoardDao.Update(*board)
	log.Printf("[ColumnController][Create] End\n")
	return 200, column
}

func (controller *ColumnController) Update(parameters map[string][]string) (int, interface{}) {
	log.Printf("[ColumnController][Update] Begin %v\n", parameters)
	validArguments, invalidMessage := checkParameters(parameters, []string{"boardId", "columnId", "column"})
	if !validArguments {
		return 412, invalidMessage
	}
	
	board := controller.BoardDao.Find(bson.ObjectIdHex(parameters["boardId"][0]))
	if board == nil {
		return 412, "Unexisting board "+parameters["boardId"][0]
	}
	
	column := domain.Column{}
	err := json.Unmarshal([]byte(parameters["column"][0]), &column)
	column.Id = bson.ObjectIdHex(parameters["columnId"][0])
	if err != nil {
		return 412, "Cannot parse column received"
	}
	succes := controller.BoardDao.UpdateColumn(board.Id, column)
	log.Printf("[ColumnController][Update] End\n")
	return 200, succes
}

func (controller *ColumnController) Delete(parameters map[string][]string) (int, interface{}) {
	log.Printf("[ColumnController][Delete] Begin %v\n", parameters)

	board := controller.BoardDao.Find(bson.ObjectIdHex(parameters["boardId"][0]))
	if board == nil {
		return 412, "Unexisting board "+parameters["boardId"][0]
	}
	
	succes := controller.BoardDao.DeleteColumn(board.Id, domain.Column{Id : bson.ObjectIdHex(parameters["columnId"][0])})
	log.Printf("[ColumnController][Delete] End\n")
	return 200, succes
}
