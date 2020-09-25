package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

// Contains scripts for maintenance of data

// Global variable for database
var db *gorm.DB

// Global variable for School,
// Temporary solution needs to be replaced by smarter one
// after getting at least basic functionality inplace.
var schoolShortName = "OAMK"

// Desc: Establish connection to database
// Status: Done
func ConnectToDB() {
	var err error

	//Maybe use global variables for database settings. But for now this is fine...
	//var dbconn = "\"" + usernamedb + ":" + userpassdb + "@tcp(db:3306)/tiukuDB?charset=utf8mb4"
	log.Printf("Trying to connect to database. <go/database.go->connectToDB>")

	//For GORM v2 following should be used, but doesn't seem to work.
	//dsn := "apiaccess:apipass@tcp(db:3306)/tiukuDB?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err = gorm.Open("mysql", "apiaccess:apipass@tcp(db:3306)/tiukuDB?charset=utf8mb4")
	if err != nil {
		log.Printf("Problem with connecting to database. <go/database.go->connectToDB>")
		log.Panic(err)
	}

	//initDB()
	fmt.Printf("%s", db.Error)
}

// Desc: Check if Student user exists
// Status: Works, maybe with slight changes could be used for all row counting?
func CheckIfUserExists(StudentID string) int64 {
	if db == nil {
		ConnectToDB()
	}

	tableToEdit := schoolShortName + "_StudentUsers"
	var tempStudent StudentUser

	result := db.Table(tableToEdit).Where("student_id = ?", StudentID).Find(&tempStudent)

	return result.RowsAffected
}

/*
func CheckIfAnonIdExists(anonid string) int {

}
*/
