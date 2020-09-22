package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
	"gorm.io/drive/mysql"
	"gorm.io/gorm"
*/
//
// Global variable for database
var db *gorm.DB

// Global variable for School,
// Temporary solution needs to be replaced by smarter solution
// After getting at basic functionality inplace.
var schoolShortName = "OAMK"

func ConnectToDB() {
	var err error

	//Maybe use global variables for database settings. But for now this is fine...
	//var dbconn = "\"" + usernamedb + ":" + userpassdb + "@tcp(db:3306)/tiukuDB?charset=utf8mb4"
	log.Printf("Trying to connect to database. <go/database.go->connectToDB>")

	// For GORM v2 following should be used, but doesn't seem to work.
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

/* initDB() for creating needed tables for database
// returns 0 if problem, 1 if everything fine.
*/
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
	//if err := db.CreateTable("")
}

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
		ResourceID:            "",
		Course:                Course{},
		SegmentNumber:         0,
		SchoolSegmentsSession: SchoolSegmentsSession{},
		SegmentCategories:     SegmentCategory{},
		Archived:              false,
	}, &StudentSegment{
		ID:                     0,
		ResourceID:             "",
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

func InitDBv3() {

	if db == nil {
		ConnectToDB()
	}

	if err := db.AutoMigrate(&Degree{
		ID:      0,
		Finnish: "",
		English: "",
	}).Error; err != nil {
		log.Panic("Problems creating table for Degrees. <go/database.go->initDB>")
	}

	if err := db.AutoMigrate(&Apartment{
		ID:      0,
		Finnish: "",
		English: "",
		Degrees: []Degree{},
	}).Error; err != nil {
		log.Panic("Problems creating table for Apartments. <go/database.go->initDB>")
	}

	//db.Model(&Apartment).AddForeignKey("primary_key", "Degrees", "RESTRICT", "RESTRICT")

	if err := db.AutoMigrate(&Campus{
		ID:         0,
		Finnish:    "",
		English:    "",
		Apartments: []Apartment{},
	}).Error; err != nil {
		log.Panic("Problems creating table for Campuses. <go/database.go->initDB>")
	}

	log.Printf("Trying to AutoMigrate Schools table to database. <go/database.go->initDB>")
	if err := db.AutoMigrate(&School{
		ID:       0,
		Finnish:  "",
		English:  "",
		Campuses: []Campus{},
	}).Error; err != nil {
		log.Panic("Problems creating table for School. <go/database.go->initDB>")
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
	}).Error; err != nil {
		log.Panic("Problems creating table for Course. <go/database.go->initDB>")
	}

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

	log.Printf("Trying to AutoMigrate StudentUsers table to database. <go/database.go->initDB>")
	if err := db.AutoMigrate(&StudentUser{
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
	if err := db.AutoMigrate(&FacultyUser{
		ID:             0,
		FacultyID:      "",
		FacultyName:    "",
		FacultyEmail:   "",
		FacultySegment: FacultySegment{},
	}).Error; err != nil {
		log.Panic("Problems creating table for FacultyUsers. <go/database.go->initDB>")
	}
	// Tables for School data

	//if err := db.CreateTable("")
}

func TestAdd() {
	db.Create(&Degree{
		ID:      0,
		Finnish: "IT",
		English: "ITC",
	})
}

/*
func testCreate() {
	test := Testi{dii: 5, fff: "No totta kai"}
	db.Create(&test)
}

func testRead() {
	var testi []Testi
	db.Find(&testi)
}
*/
