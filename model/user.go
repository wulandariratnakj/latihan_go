package model

type User struct {
	Id      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Gender  string `json:"gender"`
	Address string `json:"address"`
}
