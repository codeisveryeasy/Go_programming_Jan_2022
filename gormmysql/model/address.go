package model

type PersonAddress struct {
	Id      int `gorm:"primary_key;autoIncrement:true"`
	Area    string
	Pincode string
}
