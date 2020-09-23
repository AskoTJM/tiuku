package database

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang/gddo/httputil/header"
)

// Desc: For creating Segment table for new Student users
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

// Desc: Create Segment table for new Faculty users
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

// Desc: Create new course in Courses table.
// Status: Partially working.
func CreateCourse(r *http.Request) string {
	// Check if there is connection to database if not connect to it
	if db == nil {
		ConnectToDB()
	}

	// Check if there is table for courses.
	tableToEdit := schoolShortName + "_Courses"
	result := db.HasTable(tableToEdit)

	if !result {
		log.Panic("Problems creating new Course, table for courses doesn't exist. <database/database_create->CreateCourse>")
	} else {
		// Check if content type is set.
		if r.Header.Get("Content-Type") == "" {
			log.Panic("Problems creating new Course, no body in request information. <database/database_create->CreateCourse>")
			log.Panic("Error: No body information available.")
		} else {
			rbody, _ := header.ParseValueAndParams(r.Header, "Content-Type")
			// Check if content type is correct one.
			if rbody != "application/json" {
				log.Panic("Error: Content-Type is not application/json.")
			}

		}

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var newCourse Course
		err := dec.Decode(&newCourse)
		if err != nil {
			log.Panic("Problem with json decoding <database/database_create->CreateCourse")
		}
		db.Table(tableToEdit).Create(&newCourse)
		// Need to fix error checking.
		/*
			err2 := db.Table(tableToEdit).AutoMigrate(&newCourse)
			if err2 != nil {
				log.Panic("Problems creating new course on course table. <database/database_create->CreateCourse>")
				log.Panic(err2)
			}
		*/
		return "Done adding"
		//log.Println(newCourse)

	}
	return "no"
}
