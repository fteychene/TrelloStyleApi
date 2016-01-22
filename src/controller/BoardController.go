package controller

import (
	"domain"
	"store"
	"log"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type BoardController struct {
	BoardDao *store.BoardDao
}

func (controller *BoardController) GetAll(parameters map[string][]string) (int, interface{}) {
	log.Printf("[UserController][GetAll] Begin %v\n", parameters)
	boards := controller.BoardDao.FindAll()
	log.Printf("[UserController][GetAll] End\n")
	if boards == nil {
		return 200, []domain.Board{}
	}
	return 200, boards
}

func (controller *BoardController) Get(parameters map[string][]string) (int, interface{}) {
	log.Printf("[UserController][Get] Begin %v\n", parameters)
	board := controller.BoardDao.Find(bson.ObjectIdHex(parameters["boardId"][0]))
	log.Printf("[UserController][Get] End\n")
	return 200, board
}

func (controller *BoardController) Create(parameters map[string][]string) (int, interface{}) {
	log.Printf("[UserController][Create] Begin %v\n", parameters)
	validArguments, invalidMessage := checkParameters(parameters, []string{"board"})
	if !validArguments {
		return 412, invalidMessage
	}
	
	board := domain.Board{}
	err := json.Unmarshal([]byte(parameters["board"][0]), &board)
	board.Id = bson.NewObjectIdWithTime(time.Now())
	if err != nil {
		return 412, "Error parsing : board"
	}
	createdBoard := controller.BoardDao.Save(board)
	log.Printf("[UserController][Create] End\n")
	return 200, createdBoard
}

func (controller *BoardController) Update(parameters map[string][]string) (int, interface{}) {
	log.Printf("[UserController][Update] Begin %v\n", parameters)
	validArguments, invalidMessage := checkParameters(parameters, []string{"boardId", "board"})
	if !validArguments {
		return 412, invalidMessage
	}
	
	board := domain.Board{}
	err := json.Unmarshal([]byte(parameters["board"][0]), &board)
	board.Id = bson.ObjectIdHex(parameters["boardId"][0])
	if err != nil {
		return 412, "Error parsing : board"
	}
	updatedBoard := controller.BoardDao.Update(board)
	log.Printf("[UserController][Update] End\n")
	return 200, updatedBoard
}

func (controller *BoardController) Delete(parameters map[string][]string) (int, interface{}) {
	log.Printf("[UserController][Delete] Begin %v\n", parameters)
	
	succes := controller.BoardDao.Delete(domain.Board{Id : bson.ObjectIdHex(parameters["boardId"][0])})
	log.Printf("[UserController][Delete] End\n")
	return 200, succes
}
