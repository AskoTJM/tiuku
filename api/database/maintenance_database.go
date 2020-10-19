package database

/*
// maintenance_database.go
// Description: code for maintenance of API
*/
import (
	"log"

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
var ArchiveTableToEdit = "archived_sessions_table"

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
// W1P , Archiving Course and it's Segments + Categories works already.
func ArchiveCourse(courseToArchive Course, archive bool) {
	if Tiukudb == nil {
		ConnectToDB()
	}

	// Archiving the Course
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
		// Archiving Segment
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
			// Archiving Segments Categories
			tempCat2.Archived = archive
			Tiukudb.Save(&tempCat2)
		}
	}
}

// Get ArchivedSessions template for the segment
// W1P
func CreateArchivedSessionTemplate(segmentId uint) ArchivedSessionsTable {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var returnArchive ArchivedSessionsTable

	// Get and set Segments data
	var tempSegment Segment
	Tiukudb.Table(SegmentTableToEdit).Where("segment_id = ?", segmentId).Find(&tempSegment)
	returnArchive.SegmentName = tempSegment.SegmentName
	returnArchive.TeacherID = tempSegment.TeacherID
	returnArchive.Scope = tempSegment.Scope
	returnArchive.ExpectedAttendance = tempSegment.ExpectedAttendance
	// Get and set Courses data
	var tempCourse Course
	Tiukudb.Table(CourseTableToEdit).Where("course_id = ? ", tempSegment.CourseID).Find(&tempCourse)
	returnArchive.CourseCode = tempCourse.CourseCode
	returnArchive.CourseName = tempCourse.CourseName
	returnArchive.CourseStartDate = tempCourse.CourseStartDate
	returnArchive.CourseEndDate = tempCourse.CourseEndDate
	returnArchive.DegreeID = tempCourse.Degree
	// Get and set Degree data
	var tempDegree Degree
	Tiukudb.Table(DegreeTableToEdit).Where("id = ?", tempCourse.Degree).Find(&tempDegree)
	returnArchive.ApartmentID = tempDegree.ApartmentID
	// Get and set Apartment data
	var tempApartment Apartment
	Tiukudb.Table(ApartmentTableToEdit).Where("id = ?", tempDegree.ApartmentID).Find(&tempApartment)
	returnArchive.CampusID = tempApartment.CampusID
	// Get and set Campus data
	var tempCampus Campus
	Tiukudb.Table(SchoolsTableToEdit).Where("id = ?", tempApartment.CampusID).Find(&tempCampus)
	returnArchive.SchoolID = tempCampus.SchoolID

	return returnArchive
}

// Archive Sessions to school Archive_Table
// W1P
func ArchiveToSchoolTable(user string, segmentId uint) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseBool bool = true
	var tempStudent StudentUser
	tempArchive := CreateArchivedSessionTemplate(segmentId)
	var tempStudentSessions []StudentSegmentSession

	tempStudent = GetStudentUserWithStudentID(user)
	tempArchive.AnonID = tempStudent.AnonID
	tableToCopyFrom := tempStudent.AnonID + "_sessions"
	tableToCopyTo := tempStudent.AnonID + "_sessions_archived"
	resFrom := Tiukudb.Table(tableToCopyFrom).Where("segment_id = ?", segmentId).Find(&tempStudentSessions)
	if resFrom.Error != nil {
		responseBool = false
		log.Printf("Error reading user table in <database/maintenance_database.go->ArchiveToSchoolTable> %v", resFrom.Error)
	} else {
		var tempCategories []SegmentCategory
		Tiukudb.Table(CategoriesTableToEdit).Where("segment_id = ?", segmentId).Find(&tempCategories)

		resTo, _ := resFrom.Rows()
		var tempRes StudentSegmentSession
		for resTo.Next() {
			if err2 := resFrom.ScanRows(resTo, &tempRes); err2 != nil {
				log.Printf("Error in <database/maintenance_database.go->ArchiveToSchoolTable> %v", err2)
				responseBool = false
			}

			// Copy to Student Users Archive Table
			Tiukudb.Table(tableToCopyTo).Create(tempRes)
			// Copy to Schools ArchiveTable
			tempArchive.StartTime = tempRes.StartTime
			tempArchive.EndTime = tempRes.EndTime
			tempArchive.Created = tempRes.Created
			tempArchive.Updated = tempRes.Updated
			for i := range tempCategories {
				if tempCategories[i].ID == tempRes.Category {
					tempArchive.MainCategory = tempCategories[i].MainCategory
					tempArchive.SubCategory = tempCategories[i].SubCategory
					tempArchive.MandatoryToComment = tempCategories[i].MandatoryToComment
					tempArchive.MandatoryToTrack = tempCategories[i].MandatoryToTrack
					tempArchive.Tickable = tempCategories[i].Tickable
					break
				}
			}
			// Check if requirement for category
			if tempArchive.MandatoryToComment {
				tempArchive.Comment = tempRes.Comment
			} else {
				if tempRes.Comment == "" {
					tempArchive.Comment = "NotCommented"
				} else {
					tempArchive.Comment = "Commented"
				}
			}
		}
	}
	return responseBool
}
