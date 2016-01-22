package main 

import (
    "rest"
    "store"
    "controller"
    "log"
)

func main() {
	log.Printf("Starting the application ...\n");
	var userDao = store.UserDao{}
	var boardDao = store.BoardDao{}
	var userController = controller.UserController{UserDao: &userDao}
	var boardController = controller.BoardController{BoardDao: &boardDao}
	var columnController = controller.ColumnController{BoardDao: &boardDao}
	var ticketController = controller.TicketController{BoardDao: &boardDao}
	
	var routeMatcher = rest.NewRouteMatcher()
	
	// User 
	routeMatcher.AddRoute(rest.POST, "/user/:username", userController.Authenticate)
	routeMatcher.AddRoute(rest.POST, "/user", userController.Create)
	routeMatcher.AddRoute(rest.GET, "/user/:username", userController.Get)
	routeMatcher.AddRoute(rest.PUT, "/user/:username", userController.Update)
	routeMatcher.AddRoute(rest.DELETE, "/user/:username", userController.Delete)
	
	// Board
	routeMatcher.AddRoute(rest.GET, "/board", boardController.GetAll)
	routeMatcher.AddRoute(rest.POST, "/board", boardController.Create)
	routeMatcher.AddRoute(rest.GET, "/board/:boardId", boardController.Get)
	routeMatcher.AddRoute(rest.PUT, "/board/:boardId", boardController.Update)
	routeMatcher.AddRoute(rest.DELETE, "/board/:boardId", boardController.Delete)
	
	// Column
	routeMatcher.AddRoute(rest.GET, "/board/:boardId/column", columnController.GetAllColumnsForBoard)
	routeMatcher.AddRoute(rest.POST, "/board/:boardId/column", columnController.Create)
	routeMatcher.AddRoute(rest.GET, "/board/:boardId/column/:columnId", columnController.Get)
	routeMatcher.AddRoute(rest.PUT, "/board/:boardId/column/:columnId", columnController.Update)
	routeMatcher.AddRoute(rest.DELETE, "/board/:boardId/column/:columnId", columnController.Delete)
	
	// Ticket
	routeMatcher.AddRoute(rest.GET, "/board/:boardId/column/:columnId/ticket", ticketController.GetTicketsForBoardAndColumn)
	routeMatcher.AddRoute(rest.POST, "/board/:boardId/column/:columnId/ticket", ticketController.Create)
	routeMatcher.AddRoute(rest.GET, "/board/:boardId/column/:columnId/ticket/:ticketId", ticketController.Get)
	routeMatcher.AddRoute(rest.PUT, "/board/:boardId/column/:columnId/ticket/:ticketId", ticketController.Update)
	routeMatcher.AddRoute(rest.DELETE, "/board/:boardId/column/:columnId/ticket/:ticketId", ticketController.Delete)
	
	
	new(rest.RestServer).Start(8080, routeMatcher)

}

