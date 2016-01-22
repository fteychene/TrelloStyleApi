package controller

import (
	"domain"
	"store"
	"log"
	"encoding/json"
)

type UserController struct {
	UserDao *store.UserDao
}

func (controller *UserController) Authenticate(parameters map[string][]string) (int, interface{}) {
	log.Printf("[UserController][Authenticate] Begin %v\n", parameters)
	validArguments, invalidMessage := checkParameters(parameters, []string{"username", "password"})
	if !validArguments {
		return 412, invalidMessage
	}
	user := controller.UserDao.Find(parameters["username"][0])
	if user == nil || user.Password != parameters["password"][0] {
		return 409 , "Wrong username or password"
	}
	log.Printf("[UserController][Authenticate] End\n")
	return 200, user
}

func (controller *UserController) Get(parameters map[string][]string) (int, interface{}) {
	log.Printf("[UserController][Get] Begin %v\n", parameters)
	user := controller.UserDao.Find(parameters["username"][0])
	log.Printf("[UserController][Get] End\n")
	return 200, user
}

func (controller *UserController) Create(parameters map[string][]string) (int, interface{}) {
	log.Printf("[UserController][Create] Begin %v\n", parameters)
	validArguments, invalidMessage := checkParameters(parameters, []string{"username", "user"})
	if !validArguments {
		return 412, invalidMessage
	}
	
	user := domain.User{}
	err := json.Unmarshal([]byte(parameters["user"][0]), &user)
	if err != nil {
		return 412, "Error parsing : user"
	}
	createdUser := controller.UserDao.Update(user)
	log.Printf("[UserController][Create] End\n")
	return 200, createdUser
}

func (controller *UserController) Update(parameters map[string][]string) (int, interface{}) {
	log.Printf("[UserController][Update] Begin %v\n", parameters)
	validArguments, invalidMessage := checkParameters(parameters, []string{"username", "user"})
	if !validArguments {
		return 412, invalidMessage
	}
	
	user := domain.User{}
	err := json.Unmarshal([]byte(parameters["user"][0]), &user)
	user.Username = parameters["username"][0]
	if err != nil {
		return 412, "Error parsing : user"
	}
	succes := controller.UserDao.Update(user)
	log.Printf("[UserController][Update] End\n")
	return 200, succes
}

func (controller *UserController) Delete(parameters map[string][]string) (int, interface{}) {
	log.Printf("[UserController][Delete] Begin %v\n", parameters)
	
	succes := controller.UserDao.Delete(domain.User{Username : parameters["username"][0]})
	log.Printf("[UserController][Delete] End\n")
	return 200, succes
}
