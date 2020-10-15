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

// Validate that incoming Session has required data
// W0rks might need fine tuning.
func ValidateNewSessionStruct(newSession StudentSegmentSession) (bool, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseBool bool = true
	var responseString string
	var tempCategory SegmentCategory

	// Minimum needed to pass
	// Check for SegmentID
	if newSession.SegmentID == (StudentSegmentSession{}.SegmentID) {
		log.Printf("Segment check.")
		responseBool = false
		responseString = "Error: SegmentID required. "
	}
	// Check for Category
	var test bool
	var tempString string
	// First check if the Category and Segment match
	test, tempString, tempCategory = CheckIfSessionMatchesCategory(newSession)
	if !test {
		responseBool = false
		responseString = responseString + tempString
	} else {
		// If Category is mandatory to comment
		if tempCategory.MandatoryToComment {

			if newSession.Comment == (StudentSegmentSession{}.Comment) {
				if responseBool {
					responseBool = false
				}
				responseString = responseString + "Error: Comment required. "
			}
		}
		// Student name has to be visible.
		if tempCategory.MandatoryToTrack {
			if newSession.Privacy {
				if responseBool {
					responseBool = false
				}
				responseString = responseString + "Error: This category requires name to be visible. "
			}
		}
	}
	if responseString == "" {
		responseString = "Course Valid."
	}
	return responseBool, responseString
}

// Check if New Category has required minimum of data
// W0rks
func ValidateNewCategory(newCategory SegmentCategory) (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int = http.StatusOK
	var responseString string
	// Check if segment exists

	// Mandatory data checks
	if newCategory.SegmentID == (SegmentCategory{}.SegmentID) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = "Error: Missing SegmentID. \n"
	}
	if newCategory.MainCategory == (SegmentCategory{}.MainCategory) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing Maincategory. \n"
	}
	if newCategory.SubCategory == "" {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing Subcategory name. \n"
	}
	if responseString == "" {
		responseString = "New Category Valid."
	}
	return responseCode, responseString
}

// Check if New Course has required minimum of data CourseCode, CourseName, Degree,
// W0rks
func ValidateNewCourse(newCourse Course) (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int = http.StatusOK
	var responseString string
	// Check if segment exists

	// Data checks
	if newCourse.CourseCode == (Course{}.CourseCode) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = "Error: Missing CourseCode. \n"
	}
	if newCourse.Degree == (Course{}.Degree) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing Degree. \n"
	}
	if newCourse.CourseName == "" {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing Course name. \n"
	}
	if responseString == "" {
		responseString = "Course Valid."
	}
	return responseCode, responseString
}

// Check if New Segment has required minimum of data
// W0rks
func ValidateNewSegment(newSegment Segment) (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int = http.StatusOK
	var responseString string
	// Check if segment exists

	// Data checks
	if newSegment.CourseID == (Segment{}.CourseID) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = "Error: Missing CourseID. \n"
	}
	if newSegment.TeacherID == (Segment{}.TeacherID) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing TeachedID. \n"
	}
	if newSegment.Scope == (Segment{}.Scope) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing Scope. \n"
	}
	if newSegment.ExpectedAttendance == (Segment{}.ExpectedAttendance) {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing ExpectedAttendance. \n"
	}

	if newSegment.SegmentName == "" {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = responseString + "Error: Missing Segment name. \n"
	}
	if responseString == "" {
		responseString = "Segment Valid."
	}
	return responseCode, responseString
}

// Check if New {StudentUser} has valid data
// T35T
func ValidateNewStudentUser(newStudent StudentUser) (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int
	var responseString string
	// Check if already exists
	res := CheckIfUserExists(newStudent.StudentID)
	if res != 0 {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = "Error: StudentID already exists. \n"
	} else {
		if newStudent.StudentID == (StudentUser{}.StudentID) {
			if responseCode != http.StatusBadRequest {
				responseCode = http.StatusBadRequest
			}
			responseString = "Error: Missing StudentID. \n"
		}
		if newStudent.StudentEmail == "" {
			if responseCode != http.StatusBadRequest {
				responseCode = http.StatusBadRequest
			}
			responseString = responseString + "Error: Missing StudentEmail. \n"
		}

		if newStudent.StudentName == "" {
			if responseCode != http.StatusBadRequest {
				responseCode = http.StatusBadRequest
			}
			responseString = responseString + "Error: Missing Student name. \n"
		}
		if responseString == "" {
			responseString = "Segment Valid."
		}
	}
	return responseCode, responseString
}

// Check if New {FacultyUser} has valid data
// T35T
func ValidateNewFacultyUser(newFaculty FacultyUser) (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int
	var responseString string
	// Check if already exists
	res := CheckIfFacultyUserExists(newFaculty.FacultyID)
	if res != 0 {
		if responseCode != http.StatusBadRequest {
			responseCode = http.StatusBadRequest
		}
		responseString = "Error: StudentID already exists. \n"
	} else {
		if newFaculty.FacultyID == (FacultyUser{}.FacultyID) {
			if responseCode != http.StatusBadRequest {
				responseCode = http.StatusBadRequest
			}
			responseString = "Error: Missing StudentID. \n"
		}
		if newFaculty.FacultyEmail == "" {
			if responseCode != http.StatusBadRequest {
				responseCode = http.StatusBadRequest
			}
			responseString = responseString + "Error: Missing StudentEmail. \n"
		}

		if newFaculty.FacultyName == "" {
			if responseCode != http.StatusBadRequest {
				responseCode = http.StatusBadRequest
			}
			responseString = responseString + "Error: Missing Student name. \n"
		}
		if responseString == "" {
			responseString = "Segment Valid."
		}
	}
	return responseCode, responseString
}
