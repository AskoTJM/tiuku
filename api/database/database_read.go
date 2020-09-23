package database

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tidwall/gjson"
	//"github.com/tidwall/gjson"
)

// Global variable for database
var db *gorm.DB

// Global variable for School,
// Temporary solution needs to be replaced by smarter solution
// after getting at least basic functionality inplace.
var schoolShortName = "OAMK"

// Establish connection to database
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

// GetAnonId
// Input: StudentID
// Output: AnonID
// HOX! AnonID SHOULD NOT LEAVE OUTSIDE OF THE API
// Status: Done
func GetAnonId(StudentID string) (tempstring string) {
	if db == nil {
		ConnectToDB()
	}
	tableToEdit := schoolShortName + "_StudentUsers"
	var tempStudent StudentUser

	result := db.Table(tableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)
	//if result.Error == nil {
	//	log.Panic(result)
	//}
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	tempJSON := gjson.Get(s, "Value.AnonID")
	return tempJSON.String()
}

// GetStudent , get Students data
// Input: StudentID as string
// Output *gorm.DB (not sure about this, probably should transform to JSON)
// Status: Works, but needs more. Return value and obfuscing of AnonID if used outside
func GetStudent(StudentID string) *gorm.DB {
	if db == nil {
		ConnectToDB()
	}
	tableToEdit := schoolShortName + "_StudentUsers"
	var tempStudent StudentUser

	result := db.Table(tableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)
	//if result.Error != nil {
	//	log.Panic(result)
	//}

	return result
}

// Get Courses
// Status: Almost done, still needs switching for showing all and/or archived courses

func GetCourses() (tempstring string) {
	if db == nil {
		ConnectToDB()
	}
	tableToEdit := schoolShortName + "_Courses"
	var tempCourses []Course
	result := db.Table(tableToEdit).Where("archived = ?", false).Find(&tempCourses)
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	tempJSON := gjson.Get(s, "Value")
	return tempJSON.String()
}
