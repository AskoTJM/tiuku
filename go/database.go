package swagger

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Global variable for database
var db *gorm.DB

//var usernamedb = "apiaccess"
//var userpassdb = "apipass"

/*
type Env struct {
	db *gorm.DB
}
*/

func initDB() {
	var err error

	//Maybe use global variables for database settings. But for now this is fine...
	//var dbconn = "\"" + usernamedb + ":" + userpassdb + "@tcp(db:3306)/tiukuDB?charset=utf8mb4"
	log.Printf("Trying to connect to database")
	db, err = gorm.Open("mysql", "apiaccess:apipass@tcp(db:3306)/tiukuDB?charset=utf8mb4")
	if err != nil {
		log.Panic(err)
	}
	//test := Testi{dii: 5, fff: "No totta kai"}
	//db.AutoMigrate(&User{})
	//fmt.Printf("%s", db.GetErrors())
}

/*
func connectToDB() {

}
*/
/*
func testCreate() {
	test := Testi{dii: 5, fff: "No totta kai"}
	db.Create(&test)
}

func testRead() {
	var testi []Testi
	db.Find(&testi)
}
*/
