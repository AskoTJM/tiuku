package database

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

// Desc: Establish connection to database
// Status: Done
func ConnectToDB() {
	var err error

	//Maybe use global variables for database settings. But for now this is fine...
	//var dbconn = "\"" + usernamedb + ":" + userpassdb + "@tcp(db:3306)/tiukuDB?charset=utf8mb4"
	log.Printf("Trying to connect to database. <database/database.go->connectToDB>")

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

// initDB() for creating needed tables for database
func InitDB() {

	if db == nil {
		ConnectToDB()
	}

	log.Printf("Trying to AutoMigrate Course table to database. <database/maintenance_database.go->initDB>")
	//if err := db.Table(schoolShortName + "_Courses").AutoMigrate(&Course{
	if err := db.AutoMigrate(&Course{
		ID: 0,
		//ResourceID:      0,
		Degree:          Degree{},
		CourseCode:      "",
		CourseName:      "",
		CourseStartDate: "",
		CourseEndDate:   "",
		Archived:        false,
		Segment:         []Segment{},
	}).Error; err != nil {
		log.Println("Problems creating table for Course. <database/maintenance_database.go->initDB>")
	}

	// This should be created and named when creating Segment

	log.Printf("Trying to AutoMigrate Segment table to database. <database/maintenance_database.go->initDB>")
	if err := db.AutoMigrate(&Segment{
		ID:                    0,
		CourseID:              0,
		SegmentName:           "",
		TeacherID:             0,
		Scope:                 0,
		SegmentCategories:     "", //SegmentCategory{},
		ExpectedAttendance:    0,
		SchoolSegmentsSession: SchoolSegmentsSession{},
	}).Error; err != nil {
		log.Println("Problems creating table for Segment. <database/maintenance_database.go->initDB>")
	}

	log.Printf("Trying to AutoMigrate Course table to database. <database/maintenance_database.go->initDB>")
	//if err := db.Table(schoolShortName + "_Courses").AutoMigrate(&Course{
	if err := db.AutoMigrate(&MainCategory{
		ID:        0,
		Shorthand: "",
		Finnish:   "",
		English:   "",
	}).Error; err != nil {
		log.Println("Problems creating table for Course. <database/maintenance_database.go->initDB>")
	}

	// This should be created and named when creating new Faculty User
	/*
		log.Printf("Trying to AutoMigrate Faculty table to database. <database/database.go->initDB>")
		if err := db.AutoMigrate(&FacultySegment{
			ID:                    0,
			ResourceID:            "",
			Course:                Course{},
			SegmentNumber:         0,
			SchoolSegmentsSession: SchoolSegmentsSession{},
			SegmentCategories:     SegmentCategory{},
			Archived:              false,
		}).Error; err != nil {
			log.Println("Problems creating table for FacultySegments. <database/maintenance_database.go->initDB>")
		}
	*/

	// This should be created and named when new Student Users is created
	/*
		log.Printf("Trying to AutoMigrate StudentSegments table to database. <database/maintenance_database.go->initDB>")
		if err := db.AutoMigrate(&StudentSegment{
			ID:                     0,
			ResourceID:             "",
			Course:                 Course{},
			SegmentNumber:          0,
			StudentSegmentSessions: StudentSegmentSession{},
			SegmentCategory:        SegmentCategory{},
			Archived:               false,
		}).Error; err != nil {
			log.Println("Problems creating table for StudentSegments. <database/maintenance_database.go->initDB>")
		}
	*/
	log.Printf("Trying to AutoMigrate StudentUsers table to database. <database/maintenance_database.go->initDB>")
	//if err := db.Table(schoolShortName + "_StudentUsers").AutoMigrate(&StudentUser{
	if err := db.AutoMigrate(&StudentUser{
		ID:          0,
		StudentID:   "",
		AnonID:      "",
		StudentName: "",
		//StudentSegments: StudentSegment{},
		StudentSegments: "",
		StudentEmail:    "",
		StudentClass:    "",
	}).Error; err != nil {
		log.Println("Problems creating table for StudentUsers. <database/maintenance_database.go->initDB>")
	}

	log.Printf("Trying to AutoMigrate FacultyUsers table to database. <database/maintenance_database.go->initDB>")
	//if err := db.Table(schoolShortName + "_FacultyUsers").AutoMigrate(&FacultyUser{
	if err := db.AutoMigrate(&FacultyUser{
		ID:           0,
		FacultyID:    "",
		FacultyName:  "",
		FacultyEmail: "",
		//FacultySegment: FacultySegment{},
		FacultySegment: "",
	}).Error; err != nil {
		log.Println("Problems creating table for FacultyUsers. <database/maintenance_database.go->initDB>")
	}
	// Tables for School data

	log.Printf("Trying to AutoMigrate Schools table to database. <database/maintenance_database.go->initDB>")
	if err := db.AutoMigrate(&School{
		ID:        0,
		Shorthand: "",
		Finnish:   "",
		English:   "",
		Campuses:  []Campus{},
	}).Error; err != nil {
		log.Println("Problems creating table for School. <database/maintenance_database.go->initDB>")
	}
	//if err := db.Table(schoolShortName + "_Campuses").AutoMigrate(&Campus{
	if err := db.AutoMigrate(&Campus{
		ID:         0,
		Shorthand:  "",
		Finnish:    "",
		English:    "",
		Apartments: []Apartment{},
	}).Error; err != nil {
		log.Println("Problems creating table for Campuses. <database/maintenance_database.go->initDB>")
	}
	//if err := db.Table(schoolShortName + "_Apartments").AutoMigrate(&Apartment{
	if err := db.AutoMigrate(&Apartment{
		ID:        0,
		Shorthand: "",
		Finnish:   "",
		English:   "",
		Degrees:   []Degree{},
	}).Error; err != nil {
		log.Println("Problems creating table for Apartments. <database/maintenance_database.go->initDB>")
	}
	//if err := db.Table(schoolShortName + "_Degrees").AutoMigrate(&Degree{
	if err := db.AutoMigrate(&Degree{
		ID:        0,
		Shorthand: "",
		Finnish:   "",
		English:   "",
	}).Error; err != nil {
		log.Println("Problems creating table for Degrees. <database/maintenance_database.go->initDB>")
	}

}

// Alternative test for InitDB initalization
func InitDBv2() {

	if db == nil {
		ConnectToDB()
	}

	log.Printf("Trying to AutoMigrate Course table to database. <database/maintenance_database.go->initDB>")
	if err := db.AutoMigrate(&Course{
		ID: 0,
		//ResourceID:      0,
		Degree:          Degree{},
		CourseCode:      "",
		CourseName:      "",
		CourseStartDate: "",
		CourseEndDate:   "",
		Archived:        false,
		Segment:         []Segment{},
	}, &Segment{
		ID:                    0,
		SegmentName:           "",
		TeacherID:             0,
		Scope:                 0,
		SegmentCategories:     "", //SegmentCategory{},
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
		ID:            0,
		Course:        Course{},
		SegmentNumber: 0,
		//StudentSegmentSessions: StudentSegmentSession{},
		//SegmentCategory:        SegmentCategory{},
		StudentSegmentSessions: "",
		SegmentCategory:        "",
		Archived:               false,
	}, &StudentUser{
		ID:          0,
		StudentID:   "",
		AnonID:      "",
		StudentName: "",
		//StudentSegments: StudentSegment{},
		StudentSegments: "",
		StudentEmail:    "",
		StudentClass:    "",
	}, &FacultyUser{
		ID:           0,
		FacultyID:    "",
		FacultyName:  "",
		FacultyEmail: "",
		//FacultySegment: FacultySegment{},
		FacultySegment: "",
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
		log.Println("Problems creating initial tables. <database/maintenance_database.go->initDBv2>")
	}
	//db.Model(&School.Campuses{}).AddForeignKey("")

	//if err := db.CreateTable("")
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

func CountCourses() int {
	if db == nil {
		ConnectToDB()
	}
	var numberOfRows int
	db.Table(courseTableToEdit).Count(&numberOfRows)
	return numberOfRows
}

func CountSegments() int {
	if db == nil {
		ConnectToDB()
	}
	var numberOfRows int
	db.Table(segmentTableToEdit).Count(&numberOfRows)
	return numberOfRows
}
