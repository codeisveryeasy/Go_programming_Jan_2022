package main

import (
	"fmt"
	"gormexample/model"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

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
	//db.DropTableIfExists(&model.PersonDetail{})
	//create table Person
	//db.CreateTable(&model.PersonDetail{})

	//check if table is created!
	if db.HasTable(&model.PersonDetail{}) {
		fmt.Println("Table present")
	} else {
		fmt.Println("Table not present")
	}

	//insert in table
	// pd := &model.PersonDetail{
	// 	Name:     "Obb",
	// 	Location: "Chennai",
	// }
	// db.Create(pd)

	//bulk insert
	// var persons []model.PersonDetail = []model.PersonDetail{
	// 	model.PersonDetail{Name: "Omi", Location: "Japan"},
	// 	model.PersonDetail{Name: "Lpi", Location: "Indonesia"},
	// 	model.PersonDetail{Name: "Pki", Location: "Singapore"},
	// 	model.PersonDetail{Name: "Nii", Location: "Malaysia"},
	// }

	// for _, v := range persons {
	// 	db.Create(&v)
	// }

	// pd1 := &model.PersonDetail{
	// 	Name:     "Lpi",
	// 	Location: "Indonesia",
	// }

	//find and update
	// db.Find(&pd1)
	// fmt.Println(pd1)
	//update name column for all rows to "Check"
	//db.Model(&pd1).Update("Name", "Check")

	//interact with table directly to update specifc column
	//db.Table("person_detail").Where("location=?", "Chennai").Update("name", "OBBO")

	//db.Where("location=?", "Indonesia").Delete(&model.PersonDetail{})

	//auto migrate
	//db.AutoMigrate(&model.PersonAddress{}, &model.PersonDetail{})
	//provide information about foreign keys
	//db.Model(&model.PersonDetail{}).AddForeignKey("address_id", "person_address(id)", "CASCADE", "CASCADE")

	// a1 := model.PersonAddress{
	// 	Area:    "Crater1",
	// 	Pincode: "0008",
	// }
	// p1 := model.PersonDetail{
	// 	Name:     "one",
	// 	Location: "Moon",
	// 	Address:  a1,
	// }

	// db.Save(&p1)

	// fmt.Println("Person: ", p1)
	// fmt.Println("Address: ", a1)

	// a2 := model.PersonAddress{
	// 	Area:    "Crater2",
	// 	Pincode: "0007",
	// }
	// p2 := model.PersonDetail{
	// 	Name:     "four",
	// 	Location: "Mars",
	// 	Address:  a2,
	// }

	// a3 := model.PersonAddress{
	// 	Area:    "Mountain 1",
	// 	Pincode: "0017",
	// }
	// p3 := model.PersonDetail{
	// 	Name:     "alien",
	// 	Location: "Venus",
	// 	Address:  a3,
	// }

	a4 := model.PersonAddress{
		Area:    "Volcano",
		Pincode: "0099",
	}
	// p4 := model.PersonDetail{
	// 	Name:     "Heat",
	// 	Location: "Mercury",
	// 	Address:  a4,
	// }

	// db.Save(&p2)
	// db.Save(&p3)
	// db.Save(&p4)

	// fmt.Println("Person: ", p2)
	// fmt.Println("Person: ", p3)
	// fmt.Println("Person: ", p4)

	p5 := model.PersonDetail{
		Name:     "Close",
		Location: "Mercury Moon",
		Address:  a4,
	}

	p6 := model.PersonDetail{
		Name:     "Hidden",
		Location: "Mercury Satellite",
		Address:  a4,
	}

	db.Save(&p5)
	db.Save(&p6)

	//it is creating duplicate values for a4 in person_address table
	fmt.Println("Person: ", p5)
	fmt.Println("Person: ", p6)
}
