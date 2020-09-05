package swagger

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Global variable for database
var db *gorm.DB

type User struct {
	ID       int
	Username string
}

type UserModel struct {
	ID      int `gorm:"primary_key"`
	Name    string
	Address string
}

type Testi struct {
	dii int
	fff string
}

func initDB() {
	var err error
	db, err := gorm.Open("mysql", "testiuser:testipassword@tcp(localhost:3306)/tiukuDB?charset=utf8")
	if err != nil {
		log.Panic(err)
	}

}

func connectToDB() {

}

func testCreate() {
	test := Testi{dii: 5, fff: "No totta kai"}
	db.Create(&test)
}

func testRead() {
	var testi []Testi
	db.Find(&testi)
}
