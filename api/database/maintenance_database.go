package database

/*
// maintenance_database.go
// Description: code for maintenance of API
*/
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

// Global variables for School, etc
// Should be temporary solution, now just easier to change naming conventions
// Maybe replace with configuration file?
//var schoolShortName = "OAMK"
var courseTableToEdit = "courses"
var segmentTableToEdit = "segments"
var studentsTableToEdit = "student_users"
var facultyTableToEdit = "faculty_users"
var categoriesTableToEdit = "segment_categories"

// Debug mode for spamming your logs
var debugMode bool = true

// Desc: Establish connection to database
// Status: Done
func ConnectToDB() {
	var err error

	//Maybe use global variables for database settings. But for now this is fine...
	//var dbconn = "\"" + usernamedb + ":" + userpassdb + "@tcp(db:3306)/tiukuDB?charset=utf8mb4"
	if debugMode {
		log.Printf("Trying to connect to database. <database/database.go->connectToDB>")
	}
	//For GORM v2 following should be used, but doesn't seem to work.
	//dsn := "apiaccess:apipass@tcp(db:3306)/tiukuDB?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err = gorm.Open("mysql", "apiaccess:apipass@tcp(db:3306)/tiukuDB?charset=utf8mb4")
	if err != nil {
		log.Printf("Problem with connecting to database. <database/database.go->connectToDB>")
		log.Println(err)
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

	//tableToEdit := schoolShortName + "_StudentUsers"
	var tempStudent StudentUser

	result := db.Table(studentsTableToEdit).Where("student_id = ?", StudentID).Find(&tempStudent)

	return result.RowsAffected
}

func CheckIfFacultyUserExists(FacultyID string) int64 {
	if db == nil {
		ConnectToDB()
	}

	//tableToEdit := schoolShortName + "_StudentUsers"
	var tempFaculty FacultyUser

	result := db.Table(facultyTableToEdit).Where("faculty_id = ?", FacultyID).Find(&tempFaculty)

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

// desc: Count how many rows there are in the table.
// status:
func CountTableRows(tableToEdit string) int {
	if db == nil {
		ConnectToDB()
	}
	var numberOfRows int
	db.Table(tableToEdit).Count(&numberOfRows)
	return numberOfRows
}

// desc: Toggle Archive status of course, it's segments and categories, true to archive, false to un-archive
// status:
func ArchiveCourse(courseToArchive Course, archive bool) {
	if db == nil {
		ConnectToDB()
	}

	courseToArchive.Archived = archive
	db.Save(&courseToArchive)
	// Set Courses Segment to Archived
	var tempSegment []Segment
	result := db.Table(segmentTableToEdit).Where("course_id = ?", courseToArchive.ID).Find(&tempSegment)
	if result != nil {
		log.Println(result)
	}
	result2, _ := result.Rows()
	var tempSegment2 Segment
	for result2.Next() {
		if err3 := result.ScanRows(result2, &tempSegment2); err3 != nil {
			log.Println(err3)
		}
		tempSegment2.Archived = archive
		db.Save(&tempSegment2)
		// Change Categories for Segment to Archived
		var tempCat []SegmentCategory
		resultSeg := db.Table(categoriesTableToEdit).Where("segment_id = ?", tempSegment2.ID).Find(&tempCat)
		if resultSeg != nil {
			log.Println(resultSeg)
		}

		resultSeg2, _ := resultSeg.Rows()
		var tempCat2 SegmentCategory
		for resultSeg2.Next() {
			if err4 := result.ScanRows(resultSeg2, &tempCat2); err4 != nil {
				log.Println(err4)
			}
			tempCat2.Archived = archive
			db.Save(&tempCat2)
		}
	}
}
