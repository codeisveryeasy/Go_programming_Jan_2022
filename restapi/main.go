package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Friend struct {
	Name     string `json:"name"`
	Location string `json:"loc"`
	Id       int    `json:"id"`
}

var Friends []Friend

func main() {
	Friends = []Friend{
		Friend{Name: "Obb", Location: "Chennai", Id: 1},
		Friend{Name: "Pki", Location: "Paris", Id: 2},
		Friend{Name: "Moo", Location: "Moon", Id: 3},
		Friend{Name: "Omg", Location: "Sky", Id: 4},
	}

	handleAllRequests()

}

func handleAllRequests() {

	//create instance of type mux router
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/", welcome).Methods("GET")
	muxRouter.HandleFunc("/getallfriends", allFriends).Methods("GET")
	muxRouter.HandleFunc("/getfriend/{id}", singleFriend).Methods("GET")
	muxRouter.HandleFunc("/addfriend", addNewFriend).Methods("POST")
	//http.HandleFunc("/", welcome)
	//get all friends
	// http.HandleFunc("/getallfriends", allFriends)
	// http.HandleFunc("/getfriend/id", allFriends)
	http.ListenAndServe(":8888", muxRouter)

}

func welcome(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Root endpoint called!")
	fmt.Fprint(response, "Welcome to REST with Go!")
}

func allFriends(response http.ResponseWriter, request *http.Request) {
	fmt.Println("/getallfriends called!")
	//marshalling
	json.NewEncoder(response).Encode(Friends)
}

func singleFriend(response http.ResponseWriter, request *http.Request) {
	//fmt.Println(request)
	fmt.Println(mux.Vars(request))
	fmt.Println(mux.Vars(request)["id"])
	receivedMap := mux.Vars(request)
	receivedId := receivedMap["id"]
	id, _ := strconv.Atoi(receivedId)

	for _, friend := range Friends {
		if friend.Id == id {
			json.NewEncoder(response).Encode(friend)
		}
	}

}

func addNewFriend(response http.ResponseWriter, request *http.Request) {
	//extract request body from request
	requestBody, _ := ioutil.ReadAll(request.Body)
	fmt.Println(string(requestBody))
	//create the instance of type Friend which can be inserted in Friends array
	var friend Friend
	//unmarshal
	json.Unmarshal(requestBody, &friend)
	//append to Friends
	Friends = append(Friends, friend)

	json.NewEncoder(response).Encode(Friends)

}
