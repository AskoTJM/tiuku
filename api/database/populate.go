package database

import (
	"log"
	"strconv"
	//"github.com/AskoTJM/tiuku/api/database"
)

// Place for scripts to initalization and for populating database with test data
// Stuff that should not be needed when in use.

// initDB() for creating needed tables for database
func InitDB() {

	if db == nil {
		ConnectToDB()
	}

	log.Printf("Trying to AutoMigrate Course table to database. <go/database.go->initDB>")
	if err := db.Table(schoolShortName + "_Courses").AutoMigrate(&Course{
		ID:              0,
		ResourceID:      0,
		Schools:         School{},
		CourseCode:      "",
		CourseName:      "",
		CourseStartDate: "",
		CourseEndDate:   "",
		Archived:        false,
		Segment:         []Segment{},
	}).Error; err != nil {
		log.Panic("Problems creating table for Course. <go/database.go->initDB>")
	}

	// This should be created and named when creating Segment
	/*
		log.Printf("Trying to AutoMigrate Segment table to database. <go/database.go->initDB>")
		if err := db.AutoMigrate(&Segment{
			ID:                    0,
			SegmentName:           "",
			TeacherID:             "",
			Scope:                 0,
			SegmentCategories:     SegmentCategory{},
			ExpectedAttendance:    0,
			SchoolSegmentsSession: SchoolSegmentsSession{},
		}).Error; err != nil {
			log.Panic("Problems creating table for Segment. <go/database.go->initDB>")
		}
	*/

	// This should be created and named when creating new Faculty User
	/*
		log.Printf("Trying to AutoMigrate Faculty table to database. <go/database.go->initDB>")
		if err := db.AutoMigrate(&FacultySegment{
			ID:                    0,
			ResourceID:            "",
			Course:                Course{},
			SegmentNumber:         0,
			SchoolSegmentsSession: SchoolSegmentsSession{},
			SegmentCategories:     SegmentCategory{},
			Archived:              false,
		}).Error; err != nil {
			log.Panic("Problems creating table for FacultySegments. <go/database.go->initDB>")
		}
	*/

	// This should be created and named when new Student Users is created
	/*
		log.Printf("Trying to AutoMigrate StudentSegments table to database. <go/database.go->initDB>")
		if err := db.AutoMigrate(&StudentSegment{
			ID:                     0,
			ResourceID:             "",
			Course:                 Course{},
			SegmentNumber:          0,
			StudentSegmentSessions: StudentSegmentSession{},
			SegmentCategory:        SegmentCategory{},
			Archived:               false,
		}).Error; err != nil {
			log.Panic("Problems creating table for StudentSegments. <go/database.go->initDB>")
		}
	*/
	log.Printf("Trying to AutoMigrate StudentUsers table to database. <go/database.go->initDB>")
	if err := db.Table(schoolShortName + "_StudentUsers").AutoMigrate(&StudentUser{
		ID:              0,
		StudentID:       "",
		AnonID:          "",
		StudentName:     "",
		StudentSegments: StudentSegment{},
		StudentEmail:    "",
		StudentClass:    "",
	}).Error; err != nil {
		log.Panic("Problems creating table for StudentUsers. <go/database.go->initDB>")
	}

	log.Printf("Trying to AutoMigrate FacultyUsers table to database. <go/database.go->initDB>")
	if err := db.Table(schoolShortName + "_FacultyUsers").AutoMigrate(&FacultyUser{
		ID:             0,
		FacultyID:      "",
		FacultyName:    "",
		FacultyEmail:   "",
		FacultySegment: FacultySegment{},
	}).Error; err != nil {
		log.Panic("Problems creating table for FacultyUsers. <go/database.go->initDB>")
	}
	// Tables for School data

	log.Printf("Trying to AutoMigrate Schools table to database. <go/database.go->initDB>")
	if err := db.AutoMigrate(&School{
		ID:       0,
		Finnish:  "",
		English:  "",
		Campuses: []Campus{},
	}).Error; err != nil {
		log.Panic("Problems creating table for School. <go/database.go->initDB>")
	}

	if err := db.Table(schoolShortName + "_Campuses").AutoMigrate(&Campus{
		ID:         0,
		Finnish:    "",
		English:    "",
		Apartments: []Apartment{},
	}).Error; err != nil {
		log.Panic("Problems creating table for Campuses. <go/database.go->initDB>")
	}

	if err := db.Table(schoolShortName + "_Apartments").AutoMigrate(&Apartment{
		ID:      0,
		Finnish: "",
		English: "",
		Degrees: []Degree{},
	}).Error; err != nil {
		log.Panic("Problems creating table for Apartments. <go/database.go->initDB>")
	}

	if err := db.Table(schoolShortName + "_Degrees").AutoMigrate(&Degree{
		ID:      0,
		Finnish: "",
		English: "",
	}).Error; err != nil {
		log.Panic("Problems creating table for Degrees. <go/database.go->initDB>")
	}

}

// Alternative test for InitDB initalization
func InitDBv2() {

	if db == nil {
		ConnectToDB()
	}

	log.Printf("Trying to AutoMigrate Course table to database. <go/database.go->initDB>")
	if err := db.AutoMigrate(&Course{
		ID:              0,
		ResourceID:      0,
		Schools:         School{},
		CourseCode:      "",
		CourseName:      "",
		CourseStartDate: "",
		CourseEndDate:   "",
		Archived:        false,
		Segment:         []Segment{},
	}, &Segment{
		ID:                    0,
		SegmentName:           "",
		TeacherID:             "",
		Scope:                 0,
		SegmentCategories:     SegmentCategory{},
		ExpectedAttendance:    0,
		SchoolSegmentsSession: SchoolSegmentsSession{},
	}, &FacultySegment{
		ID:                    0,
		Course:                Course{},
		SegmentNumber:         0,
		SchoolSegmentsSession: SchoolSegmentsSession{},
		SegmentCategories:     SegmentCategory{},
		Archived:              false,
	}, &StudentSegment{
		ID:                     0,
		Course:                 Course{},
		SegmentNumber:          0,
		StudentSegmentSessions: StudentSegmentSession{},
		SegmentCategory:        SegmentCategory{},
		Archived:               false,
	}, &StudentUser{
		ID:              0,
		StudentID:       "",
		AnonID:          "",
		StudentName:     "",
		StudentSegments: StudentSegment{},
		StudentEmail:    "",
		StudentClass:    "",
	}, &FacultyUser{
		ID:             0,
		FacultyID:      "",
		FacultyName:    "",
		FacultyEmail:   "",
		FacultySegment: FacultySegment{},
	}, &School{
		ID:       0,
		Finnish:  "",
		English:  "",
		Campuses: []Campus{},
	}, &Campus{
		ID:         0,
		Finnish:    "",
		English:    "",
		Apartments: []Apartment{},
	}, &Apartment{
		ID:      0,
		Finnish: "",
		English: "",
		Degrees: []Degree{},
	}, &Degree{
		ID:      0,
		Finnish: "",
		English: "",
	}).Error; err != nil {
		log.Panic("Problems creating initial tables. <go/database.go->initDBv2>")
	}
	//db.Model(&School.Campuses{}).AddForeignKey("")

	//if err := db.CreateTable("")
}

func PopulateSchool() {
	if db == nil {
		ConnectToDB()
	}

	if err := db.Table(schoolShortName + "_Degrees").Create(&Degree{
		ID:        0,
		Shorthand: "bEng",
		Finnish:   "Insinööri (AMK), tieto- ja viestintätekniikka",
		English:   "Bachelor of Engineering, Information Technology",
	}).Error; err != nil {
		log.Panic("Problems populating table of Degrees. <go/populate.go->populateSchool>")
	}

	if err := db.Table(schoolShortName + "_Apartments").Create(&Apartment{
		ID:        0,
		Shorthand: "ICT",
		Finnish:   "Informaatioteknologia",
		English:   "Information Technology",
		Degrees:   []Degree{},
	}).Error; err != nil {
		log.Panic("Problems populating table of Apartments. <go/populate.go->populateSchool>")
	}

	if err := db.Table(schoolShortName + "_Campuses").Create(&Campus{
		ID:         0,
		Shorthand:  "Linna",
		Finnish:    "Linnanmaan Kampus",
		English:    "Campus Linnanmaa",
		Apartments: []Apartment{},
	}).Error; err != nil {
		log.Panic("Problems populating table of Campuses. <go/populate.go->populateSchool>")
	}

	if err := db.Create(&School{
		ID:        0,
		Shorthand: "OAMK",
		Finnish:   "Oulun Ammattikorkeakoulu",
		English:   "Oulu University of Applied Sciences",
		Campuses:  []Campus{},
	}).Error; err != nil {
		log.Panic("Problems populating table of Schools. <go/populate.go->populateSchool>")
	}

}

// desc: Auto-generating student users for testing purposes
func PopulateStudents(p int) {
	if db == nil {
		ConnectToDB()
	}

	// Null student for testing
	/*
		if err := db.Table(schoolShortName + "_StudentUsers").Create(&StudentUser{
			ID:              0,
			StudentID:       "oppi0",
			AnonID:          "Anon0",
			StudentName:     "Oppilas 0",
			StudentSegments: StudentSegment{},
			StudentEmail:    "oppilas0@oppilaitos.fi",
			StudentClass:    "tit0",
		}).Error; err != nil {
			log.Panic("Problems populating table of StudentUsers. <go/populate.go->populateStudents>")
		}
	*/

	for i := 0; i < p; i = i + 1 {

		// Switching auto-generated classes
		classToAdd := ""
		if i%2 == 0 {
			classToAdd = "tit2"
		} else {
			classToAdd = "tit1"
		}

		if err := db.Table(schoolShortName + "_StudentUsers").Create(&StudentUser{
			ID:              0,
			StudentID:       "oppi" + strconv.Itoa(i),
			AnonID:          "Anon" + strconv.Itoa(i),
			StudentName:     "Oppilas " + strconv.Itoa(i),
			StudentSegments: StudentSegment{},
			StudentEmail:    "oppilas" + strconv.Itoa(i) + "@oppilaitos.fi",
			StudentClass:    classToAdd,
		}).Error; err != nil {
			log.Panic("Problems populating table of StudentUsers. <go/populate.go->populateStudents>")
		}
	}

}

// desc: Auto-generating courses for testing purposes
func PopulateCourses(p int) {
	if db == nil {
		ConnectToDB()
	}
	for i := 0; i < p; i = i + 1 {

		// Auto-generating archived status
		archivedToAdd := false
		if i%2 == 0 {
			archivedToAdd = true
		} else {
			archivedToAdd = false
		}

		if err := db.Table(schoolShortName + "_Courses").Create(&Course{
			ID:              0,
			ResourceID:      0,
			Schools:         School{},
			CourseCode:      "GTC" + strconv.Itoa(i),
			CourseName:      "Generated Test Course " + strconv.Itoa(i),
			CourseStartDate: strconv.Itoa(i) + "." + strconv.Itoa(i) + ".2020",
			CourseEndDate:   strconv.Itoa(i) + "." + strconv.Itoa(i) + ".2021",
			Archived:        archivedToAdd,
			Segment:         []Segment{},
		}).Error; err != nil {
			log.Panic("Problems populating Courses table. <go/populate.go->populateCourses>")
		}
	}

}
