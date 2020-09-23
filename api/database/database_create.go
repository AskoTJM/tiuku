package database

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang/gddo/httputil/header"
)

// CreateStudentSegmentTable
// Description: For creating Segment table for new Student users
// Input: StudentID
// Status: Works
func CreateStudentSegmentTable(myAnonID string) string {
	if db == nil {
		ConnectToDB()
	}
	tableToEdit := myAnonID + "_SegmentsTable"
	result := db.HasTable(tableToEdit)
	if result {
		log.Println("Error: Table already exists. <database/database_maintenance.go->CreateStudentSegmentTable>")
		return "Error: Table already exists."
	} else {
		if err := db.Table(tableToEdit).AutoMigrate(&StudentSegment{
			ID:                     0,
			ResourceID:             "",
			StudentID:              myAnonID,
			Course:                 Course{},
			SegmentNumber:          0,
			StudentSegmentSessions: StudentSegmentSession{},
			SegmentCategory:        SegmentCategory{},
			Archived:               false,
		}).Error; err != nil {
			log.Panic("Problems creating Segment table of StudentUsers. <database/database_create->CreateStudentSegmentTable>")
		}
		return tableToEdit
	}
}

// CreateFacultySegmentTable
// Description: For creating Segment table for new Faculty users
// Input: FacultyID
// Status: No clue, just copy+pasted and edited from CreateStudentSegmentTable
func CreateFacultySegmentTable(myFacultyID string) string {
	if db == nil {
		ConnectToDB()
	}
	tableToEdit := myFacultyID + "_SegmentsTable"
	result := db.HasTable(tableToEdit)
	if result {
		log.Println("Error: Table already exists. <database/database_maintenance.go->CreateFacultySegmentTable>")
		return "Error: Table already exists."
	} else {
		if err := db.Table(tableToEdit).AutoMigrate(&FacultySegment{
			ID:                    0,
			ResourceID:            "",
			Course:                Course{},
			SegmentNumber:         0,
			SchoolSegmentsSession: SchoolSegmentsSession{},
			SegmentCategories:     SegmentCategory{},
			Archived:              false,
		}).Error; err != nil {
			log.Panic("Problems creating Segment table of FacultyUsers. <database/database_create->CreateFacultySegmentTable>")
		}
		return tableToEdit
	}
}

func CreateCourse(r *http.Request) string {
	// Check if there is connection to database if not connect to it
	if db == nil {
		ConnectToDB()
	}

	// Check if there is table for courses.
	tableToEdit := schoolShortName + "_Courses"
	result := db.HasTable(tableToEdit)
	if !result {
		log.Panic("Problems creating new Course, table doesn't exist. <database/database_create->CreateCourse>")
	} else {
		// Check if content type is set.
		if r.Header.Get("Content-Type") == "" {
			//log.Panic("Problems creating new Course, no body in request information. <database/database_create->CreateCourse>")
			return "Error: No body information available."
		} else {
			rbody, _ := header.ParseValueAndParams(r.Header, "Content-Type")
			// Check if content type is correct one.
			if rbody != "application/json" {
				return "Error: Content-Type is not application/json."
			}

		}
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var newCourse Course
		err := dec.Decode(&newCourse)
		if err != nil {
			log.Panic("Problem with json decoding <database/database_create->CreateCourse")
		}
		log.Println(newCourse)
		return "What is this?"
	}
	return "What is this2?"
}
