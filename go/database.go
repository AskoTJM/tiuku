package swagger

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

/*
	"gorm.io/drive/mysql"
	"orm.io/gorm"
*/
//
// Global variable for database
var db *gorm.DB

func connectToDB() {
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
 */
func initDB() {

	log.Printf("Trying to AutoMigrate Course table to database. <go/database.go->initDB>")
	if err := db.AutoMigrate(&Course{
		ID:              0,
		ResourceID:      0,
		Schools:         Schools{},
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
		ID:                     0,
		SegmentName:            "",
		TeacherID:              "",
		Scope:                  0,
		SegmentCategories:      SegmentCategories{},
		ExpectedAttendance:     0,
		SchoolSegmentsSessions: SchoolSegmentsSessions{},
	}).Error; err != nil {
		log.Panic("Problems creating table for Segment. <go/database.go->initDB>")
	}

	log.Printf("Trying to AutoMigrate Faculty table to database. <go/database.go->initDB>")
	if err := db.AutoMigrate(&FacultySegments{
		ID:                     0,
		ResourceID:             "",
		Course:                 Course{},
		SegmentNumber:          0,
		SchoolSegmentsSessions: SchoolSegmentsSessions{},
		SegmentCategories:      SegmentCategories{},
		Archived:               false,
	}).Error; err != nil {
		log.Panic("Problems creating table for FacultySegments. <go/database.go->initDB>")
	}

	log.Printf("Trying to AutoMigrate StudentSegments table to database. <go/database.go->initDB>")
	if err := db.AutoMigrate(&StudentSegments{
		ID:                     0,
		ResourceID:             "",
		Course:                 Course{},
		SegmentNumber:          0,
		StudentSegmentSessions: StudentSegmentSessions{},
		SegmentCategories:      SegmentCategories{},
		Archived:               false,
	}).Error; err != nil {
		log.Panic("Problems creating table for StudentSegments. <go/database.go->initDB>")
	}

	log.Printf("Trying to AutoMigrate StudentUsers table to database. <go/database.go->initDB>")
	if err := db.AutoMigrate(&StudentUsers{
		ID:              0,
		StudentID:       "",
		AnonID:          "",
		StudentName:     "",
		StudentSegments: StudentSegments{},
		StudentEmail:    "",
		StudentClass:    "",
	}).Error; err != nil {
		log.Panic("Problems creating table for StudentUsers. <go/database.go->initDB>")
	}

	log.Printf("Trying to AutoMigrate FacultyUsers table to database. <go/database.go->initDB>")
	if err := db.AutoMigrate(&FacultyUsers{
		ID:              0,
		FacultyID:       "",
		FacultyName:     "",
		FacultyEmail:    "",
		FacultySegments: FacultySegments{},
	}).Error; err != nil {
		log.Panic("Problems creating table for FacultyUsers. <go/database.go->initDB>")
	}
	// Tables for School data

	log.Printf("Trying to AutoMigrate Schools table to database. <go/database.go->initDB>")
	if err := db.AutoMigrate(&Schools{
		ID:       0,
		Finnish:  "",
		English:  "",
		Campuses: []Campuses{},
	}).Error; err != nil {
		log.Panic("Problems creating table for School. <go/database.go->initDB>")
	}

	if err := db.Table("OAMK").AutoMigrate(&Campuses{
		ID:         0,
		Finnish:    "",
		English:    "",
		Apartments: []Apartments{},
	}).Error; err != nil {
		log.Panic("Problems creating table for Campuses. <go/database.go->initDB>")
	}

	//if err := db.CreateTable("")

}

func testAdd() {
	db.Create(&Degrees{
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
