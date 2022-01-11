package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type PersonDetail struct {
	Id       int `gorm:"primary_key;autoIncrement:true"`
	Name     string
	Location string
}

func main() {
	db, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/friendsorm?checkConnLiveness=false&maxAllowedPacket=0")
	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection success")
	}
	defer db.Close()

	//singular name
	db.SingularTable(true)
	//drop if table is already present
	//db.DropTableIfExists(&PersonDetail{})
	//create table Person
	//db.CreateTable(&PersonDetail{})

	//check if table is created!
	if db.HasTable(&PersonDetail{}) {
		fmt.Println("Table present")
	} else {
		fmt.Println("Table not present")
	}

	//insert in table
	// pd := &PersonDetail{
	// 	Name:     "Obb",
	// 	Location: "Chennai",
	// }
	// db.Create(pd)

	//bulk insert
	var persons []PersonDetail = []PersonDetail{
		PersonDetail{Name: "Omi", Location: "Japan"},
		PersonDetail{Name: "Lpi", Location: "Indonesia"},
		PersonDetail{Name: "Pki", Location: "Singapore"},
		PersonDetail{Name: "Nii", Location: "Malaysia"},
	}

	for _, v := range persons {
		db.Create(&v)
	}

}
