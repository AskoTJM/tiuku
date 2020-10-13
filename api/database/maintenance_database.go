package database

/*
// maintenance_database.go
// Description: code for maintenance of API
*/
import (
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
var SchoolParticipationList = "school_segments_sessions"

var DegreeTableToEdit = "degrees"
var ApartmentTableToEdit = "apartments"
var CampusTableToEdit = "campus"
var SchoolsTableToEdit = "schools"

// Variable for empty field in mySQL, because GORM
var StringForEmpy = "N0TS3T"

// Debug mode for spamming your logs
var DebugMode bool = true

// Establish connection to database
// W0rks
func ConnectToDB() {
	var err error

	//Maybe use global variables for database settings. But for now this is fine...
	//var dbconn = "\"" + usernamedb + ":" + userpassdb + "@tcp(db:3306)/tiukuDB?charset=utf8mb4"
	if DebugMode {
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
	//fmt.Printf("%s", Tiukudb.Error)
}

// Check if Student user exists student users table, returns ID if does.
// W0rks , maybe with slight changes could be used for all row counting?
func CheckIfUserExists(StudentID string) int64 {
	if Tiukudb == nil {
		ConnectToDB()
	}

	//tableToEdit := schoolShortName + "_StudentUsers"
	var tempStudent StudentUser

	result := Tiukudb.Table(StudentsTableToEdit).Where("student_id = ?", StudentID).Find(&tempStudent)

	return result.RowsAffected
}

// Check if Faculty User is in DB
// ??
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
// W0rks
func CountTableRows(tableToEdit string) int {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var numberOfRows int
	Tiukudb.Table(tableToEdit).Count(&numberOfRows)
	return numberOfRows
}

// Toggle Archive status of course, it's segments and categories, true to archive, false to un-archive
// ??
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
// T0D0
func CheckArchiveConflicts() int {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var numberOfFixes int

	return numberOfFixes
}

// Test if required tables exist, should include all the necessary table or fail
// W0rks
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
// W0rks
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

// Check if Category matches Segment, returns True if match False if not
// W0rks
func CheckIfCategoryMatchSegment(testCategory uint, testSegment uint) (bool, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var responseBool bool
	var responseString string
	var tempCategory SegmentCategory
	if testCategory == 0 {
		responseBool = false
		responseString = "Category not provided or incorrect one."
	} else if testCategory > 3 {
		res := Tiukudb.Table(CategoriesTableToEdit).Where("id = ?", testCategory).Where("segment_id = ?", testSegment).Find(&tempCategory).RowsAffected
		if res == 0 {
			responseBool = false
			responseString = "Error. Incorrect Category for Segment."
		}
		if res == 1 {
			responseBool = true
			responseString = "Category matches the Segment."
		}
	} else {
		responseBool = true
		responseString = "Category is default one."
	}
	return responseBool, responseString
}

// Check if Session matches Category, returns True if match False if not
// T0D0
func CheckIfSessionMatchesCategory(tempSession StudentSegmentSession) (bool, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}

	var responseBool bool
	var responseString string
	var tempCategory SegmentCategory
	if tempSession.Category == 0 {
		responseBool = false
		responseString = "Category not provided or incorrect one."
	} else if tempSession.Category > 3 {
		result := Tiukudb.Table(CategoriesTableToEdit).Where("id = ?", tempSession.Category).Where("segment_id = ?", tempSession.SegmentID).Find(&tempCategory)
		res := result.RowsAffected
		if res == 0 {
			responseBool = false
			responseString = "Error. Incorrect Category for Segment."
		}
		if res == 1 {
			// If the Category matches
			//tempCategory.
			responseBool = true
			responseString = "Category matches the Segment."
		}
	} else {
		responseBool = true
		responseString = "Category is default one."
	}
	return responseBool, responseString
}
