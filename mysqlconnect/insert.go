package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Friend struct {
	ID             int64
	FriendName     string
	FriendLocation string
	FriendSince    int64
}

func main() {

	config := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "friends",
		AllowNativePasswords: true,
	}
	fmt.Println(config.FormatDSN())
	//get database handle
	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Println("Error: ", err)

	}

	pingError := db.Ping()
	if pingError != nil {
		log.Println("Ping Error:", pingError)
	}
	fmt.Println("Connected!")

	f := Friend{
		FriendName:     "Fool",
		FriendLocation: "Chennai",
		FriendSince:    20,
	}
	fmt.Println("Friend instance passing!")
	//add a new friend
	id, err := addFriend(&f)
	if err != nil {
		log.Println("Error: ", err)

	}

	fmt.Println("New friend inserted with id:", id)

	//retrieve all friends
	printFriends()
}

func addFriend(f *Friend) (int64, error) {

	result, err := db.Exec("insert into detail (friendname, friendlocation, friendsince) values (?,?,?)", f.FriendName, f.FriendLocation, f.FriendSince)
	if err != nil {
		return 0, fmt.Errorf("addFriend: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addFriend: %v", err)
	}

	return id, nil

}

func printFriends() {

	var friends []Friend
	rows, err := db.Query("select * from detail")
	if err != nil {
		log.Println("Select Error:", err)
	}

	defer rows.Close()
	//iterate on rows
	for rows.Next() {
		var f Friend
		error := rows.Scan(&f.ID, &f.FriendName, &f.FriendLocation, &f.FriendSince)
		if error != nil {
			log.Println("Error in scan", error)
			return
		}
		friends = append(friends, f)
	}
	fmt.Println(friends)

}
