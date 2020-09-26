package database

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// Contains scripts for maintenance of data

// Global variable for database
var db *gorm.DB

// Global variable for School,
// Temporary solution needs to be replaced by smarter one
// after getting at least basic functionality inplace.
var schoolShortName = "OAMK"
var courseTableToEdit = "courses"
var segmentTableToEdit = "segments"

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

// DOESN'T WORK! DON'T USE!
func CheckAssociation(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		ConnectToDB()
	}

	var checkSchool School
	var checkCampus Campus
	result := db.Model(&checkSchool).Association("Campus").Find(&checkCampus)
	if result.Error != nil {
		log.Println(result)
	} else {
		anon, _ := json.Marshal(result)
		log.Println(anon)
		n := len(anon)
		s := string(anon[:n])
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		log.Println(s)
		fmt.Fprintf(w, "%s", s)
	}
}
