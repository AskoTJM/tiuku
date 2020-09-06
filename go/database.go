package swagger

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Global variable for database
var db *gorm.DB
var usernamedb = "apiaccess"
var userpassdb = "apipass"

/*
type Env struct {
	db *gorm.DB
}
*/
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
	log.Printf("Trying to connect to database")
	db, err = gorm.Open("mysql", "apiaccess:apipass@tcp(db:3306)/tiukuDB?charset=utf8")
	if err != nil {
		log.Panic(err)
	}
	//test := Testi{dii: 5, fff: "No totta kai"}
	db.AutoMigrate(&User{})
	//fmt.Printf("%s", db.GetErrors())
}

/*
func connectToDB() {

}
*/

func testCreate() {
	test := Testi{dii: 5, fff: "No totta kai"}
	db.Create(&test)
}

func testRead() {
	var testi []Testi
	db.Find(&testi)
}
