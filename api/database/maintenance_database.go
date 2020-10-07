package database

/*
// maintenance_database.go
// Description: code for maintenance of API
*/
import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// Contains scripts for maintenance of data

// Global variable for database
var Tiukudb *gorm.DB

// Global variables for School, etc
// Should be temporary solution, now just easier to change naming conventions
// Maybe at least replace with configuration file?
//var schoolShortName = "OAMK"
var CourseTableToEdit = "courses"
var SegmentTableToEdit = "segments"
var StudentsTableToEdit = "student_users"
var FacultyTableToEdit = "faculty_users"
var CategoriesTableToEdit = "segment_categories"
var EnrollmentSegmentList = "school_segments_sessions"

var DegreeTableToEdit = "degrees"
var ApartmentTableToEdit = "apartments"
var CampusTableToEdit = "campus"
var SchoolsTableToEdit = "schools"

// Debug mode for spamming your logs
var debugMode bool = true

// Establish connection to database
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

	Tiukudb, err = gorm.Open("mysql", "apiaccess:apipass@tcp(db:3306)/tiukuDB?charset=utf8mb4")
	if err != nil {
		log.Printf("Problem with connecting to database. <database/database.go->connectToDB>")
		log.Println(err)
	}

	//initDB()
	fmt.Printf("%s", Tiukudb.Error)
}

// Check if Student user exists student users table, returns ID if does.
// Status: Works, maybe with slight changes could be used for all row counting?
func CheckIfUserExists(StudentID string) int64 {
	if Tiukudb == nil {
		ConnectToDB()
	}

	//tableToEdit := schoolShortName + "_StudentUsers"
	var tempStudent StudentUser

	result := Tiukudb.Table(StudentsTableToEdit).Where("student_id = ?", StudentID).Find(&tempStudent)

	return result.RowsAffected
}

func CheckIfFacultyUserExists(FacultyID string) int64 {
	if Tiukudb == nil {
		ConnectToDB()
	}

	//tableToEdit := schoolShortName + "_StudentUsers"
	var tempFaculty FacultyUser

	result := Tiukudb.Table(FacultyTableToEdit).Where("faculty_id = ?", FacultyID).Find(&tempFaculty)

	return result.RowsAffected
}

/*
func CheckIfAnonIdExists(anonid string) int {

}
*/

// Count how many rows there are in the table. Can be used to count users, segments, course etc in table.
// status:
func CountTableRows(tableToEdit string) int {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var numberOfRows int
	Tiukudb.Table(tableToEdit).Count(&numberOfRows)
	return numberOfRows
}

// Toggle Archive status of course, it's segments and categories, true to archive, false to un-archive
// status:
func ArchiveCourse(courseToArchive Course, archive bool) {
	if Tiukudb == nil {
		ConnectToDB()
	}

	courseToArchive.Archived = archive
	Tiukudb.Save(&courseToArchive)
	// Set Courses Segment to Archived
	var tempSegment []Segment
	result := Tiukudb.Table(SegmentTableToEdit).Where("course_id = ?", courseToArchive.ID).Find(&tempSegment)
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
		Tiukudb.Save(&tempSegment2)
		// Change Categories for Segment to Archived
		var tempCat []SegmentCategory
		resultSeg := Tiukudb.Table(CategoriesTableToEdit).Where("segment_id = ?", tempSegment2.ID).Find(&tempCat)
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
			Tiukudb.Save(&tempCat2)
		}
	}
}

// Fix archive status of courses segments and categories
// Should work by working through courses and checking that their segments and categories have same archive status
//
func CheckArchiveConflicts() int {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var numberOfFixes int

	return numberOfFixes
}

// Test if required tables exist, should include all the necessary table or fail
func CheckIfRequiredTablesExist() bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	if !Tiukudb.HasTable(CourseTableToEdit) {
		return false
	}
	if !Tiukudb.HasTable(SegmentTableToEdit) {
		return false
	}
	if !Tiukudb.HasTable(StudentsTableToEdit) {
		return false
	}
	if !Tiukudb.HasTable(FacultyTableToEdit) {
		return false
	}
	if !Tiukudb.HasTable(CategoriesTableToEdit) {
		return false
	}
	return true
}

// Check if content of request is JSON
func CheckJSONContent(w http.ResponseWriter, r *http.Request) string {
	if r.Header.Get("Content-Type") == "" {
		// Removed to be able to use this code with empty values
		//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		//w.WriteHeader(http.StatusNoContent)
		response := "EMPTY"
		return response
	} else {
		//rbody, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		rbody := r.Header.Get("Content-Type")
		// Check if content type is correct one.
		if rbody != "application/json" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNotAcceptable)
			response := "TYPE_ERROR"
			return response
		}

	}
	return "PASS"
}
