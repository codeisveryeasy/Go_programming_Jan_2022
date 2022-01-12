package model

type PersonDetail struct {
	Id        int `gorm:"primary_key;autoIncrement:true"`
	Name      string
	Location  string
	Address   PersonAddress
	AddressId int `gorm:"ForeignKey:id"`
}
