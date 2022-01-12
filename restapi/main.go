package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	http.HandleFunc("/", welcome)
	//get all friends
	http.HandleFunc("/getallfriends", allFriends)
	http.ListenAndServe(":8888", nil)

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
