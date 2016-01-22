package main 

import (
	"domain"
//	"log"
	"encoding/json"
//	"store"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)


func main() {
//	userDao := store.UserDao{}
//	
//	user := domain.User{Username:"fteychene", Password:"fteychene", FirstName:"Teychene", SecondName:"Francois"}
//	fmt.Printf("1 -%v\n", user)
//	
//	user2 := userDao.Save(user)
//	fmt.Printf("2 -%v\n", user2)
//	
//	user3 := userDao.Find(user2.Username)
//	fmt.Printf("3 -%v\n", user3)
//	
//	user3.Password = "fteychene2"
//	user3.FirstName = "Coucou"
//	user3.SecondName = "Test"
//	user4 := userDao.Update(*user3)
//	fmt.Printf("4 -%v\n", user4)
//	
//	user5 := userDao.Find(user4.Username)
//	fmt.Printf("5 -%v\n", user5)
//	
//	fmt.Printf("Users %v \n", userDao.FindAll())

//	boardDao := store.BoardDao{}
//	
//	board := domain.Board{
//		Name: "Mon Premier Board", 
//		OwnerId: "fteychene", 
//		Columns: []domain.Column {
//			domain.Column{Id: bson.NewObjectId(), Name:"BackLok", Order : 0},
//			domain.Column{Id: bson.NewObjectId(), Name:"Appro", Order : 1},
//			domain.Column{Id: bson.NewObjectId(), Name:"In Progress", Order : 2}},
//		AuthorizedId: []string{"cavalierch", "pontch"}}
//	
//	board2 := boardDao.Save(board)
//	fmt.Printf("1 - %v\n", board2)
//	
//	board3 := boardDao.FindAllBoardsForUser(domain.User{Username: "cavalierch"})
//	fmt.Printf("2 - %v\n", board3)

//	user := domain.User{}
//	str := `{"Username":"fteychene","Password":"fteychene2","FirstName":"Coucou","SecondName":"Test"}`
//	log.Printf("%v\n",str)
//	err1 := json.Unmarshal([]byte(str), &user)
//	if err1 != nil {
//		log.Println("error:", err1)
//	}
//	log.Printf("%v\n",user)
//	
//	var jsonBlob = []byte(`[
//		{"Name": "Platypus", "Order": "Monotremata"},
//		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
//	]`)
//	type Animal struct {
//		Name  string
//		Order string
//	}
//	var animals []Animal
//	err := json.Unmarshal(jsonBlob, &animals)
//	if err != nil {
//		log.Println("error:", err)
//	}
//	log.Printf("%+v", animals)

ticket := domain.Ticket{Id: bson.NewObjectIdWithTime(time.Now()), Name: "Tache 1", Description: "Description Tache 1", OpenDate: time.Now(), AffectedUserId : "fteychene"}

byteResult, _ := json.Marshal(ticket)
fmt.Printf("%v\n",string(byteResult))
}

