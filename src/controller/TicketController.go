package controller

import (
	"domain"
	"store"
	"log"
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TicketController struct {
	BoardDao *store.BoardDao
	UserDao *store.UserDao
}

func (controller *TicketController) GetTicketsForBoardAndColumn(parameters map[string][]string) (int, interface{}) {
	log.Printf("[TicketController][GetTicketsForBoardAndColumn] Begin %v\n", parameters)
	board := controller.BoardDao.Find(bson.ObjectIdHex(parameters["boardId"][0]))
	if board == nil {
		return 412, "Unexisting board "+parameters["boardId"][0]
	}
	
	column := controller.BoardDao.FindColumnForBoard(board.Id, bson.ObjectIdHex(parameters["columnId"][0]))
	log.Printf("[TicketController][GetTicketsForBoardAndColumn] End \n") 
	if column == nil {
		return 200, []domain.Ticket{}
	}
	return 200, column.Tickets
}

func (controller *TicketController) Get(parameters map[string][]string) (int, interface{}) {
	log.Printf("[TicketController][Get] Begin %v\n", parameters)
	column := controller.BoardDao.FindTicketForBoardAndColumn(bson.ObjectIdHex(parameters["boardId"][0]), bson.ObjectIdHex(parameters["columnId"][0]), bson.ObjectIdHex(parameters["ticketId"][0]))
	log.Printf("[TicketController][Get] End\n") 
	return 200, column
}

func (controller *TicketController) Create(parameters map[string][]string) (int, interface{}) {
	log.Printf("[TicketController][Create] Begin %v\n", parameters)
	validArguments, invalidMessage := checkParameters(parameters, []string{"boardId", "columnId", "ticket"})
	if !validArguments {
		return 412, invalidMessage
	}
	
	board := controller.BoardDao.Find(bson.ObjectIdHex(parameters["boardId"][0]))
	if board == nil {
		return 412, "Unexisting board "+parameters["boardId"][0]
	}
	column := controller.BoardDao.FindColumnForBoard(board.Id, bson.ObjectIdHex(parameters["columnId"][0]))
	if board == nil {
		return 412, "Unexisting column "+parameters["columnId"][0]
	}
	
	ticket := domain.Ticket{}
	err := json.Unmarshal([]byte(parameters["ticket"][0]), &ticket)
	ticket.Id = bson.NewObjectIdWithTime(time.Now())
	
	if err != nil {
		return 412, "Cannot parse ticket received"
	}
	column.Tickets = append(column.Tickets, ticket)
	controller.BoardDao.UpdateColumn(board.Id, *column)
	log.Printf("[TicketController][Create] End\n")
	return 200, ticket
}

func (controller *TicketController) Update(parameters map[string][]string) (int, interface{}) {
	log.Printf("[TicketController][Update] Begin %v\n", parameters)
	validArguments, invalidMessage := checkParameters(parameters, []string{"boardId", "columnId", "ticketId", "ticket"})
	if !validArguments {
		return 412, invalidMessage
	}
	
	board := controller.BoardDao.Find(bson.ObjectIdHex(parameters["boardId"][0]))
	if board == nil {
		return 412, "Unexisting board "+parameters["boardId"][0]
	}
	column := controller.BoardDao.FindColumnForBoard(board.Id, bson.ObjectIdHex(parameters["columnId"][0]))
	if column == nil {
		return 412, "Unexisting column "+parameters["columnId"][0]
	}
	
	ticket := domain.Ticket{}
	err := json.Unmarshal([]byte(parameters["ticket"][0]), &ticket)
	ticket.Id = bson.ObjectIdHex(parameters["ticketId"][0])
	if err != nil {
		return 412, "Cannot parse ticket received"
	}
	updatedTicket := controller.BoardDao.UpdateTicket(board.Id, column.Id, ticket)
	log.Printf("[TicketController][Update] End\n")
	return 200, updatedTicket
}

func (controller *TicketController) Delete(parameters map[string][]string) (int, interface{}) {
	log.Printf("[TicketController][Delete] Begin %v\n", parameters)

	board := controller.BoardDao.Find(bson.ObjectIdHex(parameters["boardId"][0]))
	if board == nil {
		return 412, "Unexisting board "+parameters["boardId"][0]
	}
	
	column := controller.BoardDao.FindColumnForBoard(board.Id, bson.ObjectIdHex(parameters["columnId"][0]))
	if column == nil {
		return 412, "Unexisting column "+parameters["columnId"][0]
	}
	
	succes := controller.BoardDao.DeleteTicket(board.Id, *column, domain.Ticket{Id :bson.ObjectIdHex(parameters["ticketId"][0])})
	log.Printf("[TicketController][Delete] End\n")
	return 200, succes
}
