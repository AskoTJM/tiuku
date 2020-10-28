package database

/*
// populate.go
// Description: Should only contain scripts to autopopulate data for development and testing
// Place for scripts for populating database with test data
// Stuff that should not be needed when in use.
*/
import (
	"log"
	"strconv"

	"github.com/AskoTJM/tiuku/api/scripts"
	//"github.com/AskoTJM/tiuku/api/database"
)

// OBSOLETE! Replaced by initDB scripts.
// Populating School data and maincategories
func PopulateSchool() {
	if Tiukudb == nil {
		ConnectToDB()
	}

	if err := Tiukudb.Create(&School{
		ID:        1,
		Shorthand: "OAMK",
		Finnish:   "Oulun Ammattikorkeakoulu",
		English:   "Oulu University of Applied Sciences",
		Campuses: []Campus{{
			ID:        1,
			Shorthand: "Linna",
			Finnish:   "Linnanmaan Kampus",
			English:   "Campus Linnanmaa",
			Apartments: []Apartment{{
				ID:        1,
				Shorthand: "ICT",
				Finnish:   "Informaatioteknologia",
				English:   "Information Technology",
				Degrees: []Degree{{
					ID:        1,
					Shorthand: "bEng",
					Finnish:   "Insinööri (AMK), tieto- ja viestintätekniikka",
					English:   "Bachelor of Engineering, Information Technology",
				}},
			}},
		}},
	}).Error; err != nil {
		log.Printf("Problems populating table of Schools. <database/populate.go->populateSchool>")
	}

	if err := Tiukudb.Create(&MainCategory{
		ID:        0,
		Shorthand: "Lähi",
		Finnish:   "Lähiopetus",
		English:   "Classroom study",
	}).Error; err != nil {
		log.Printf("Problems creating main categories. <database/populate.go->populateSchool>")
	}
	if err := Tiukudb.Create(&MainCategory{
		ID:        0,
		Shorthand: "Etä",
		Finnish:   "Etäopetus",
		English:   "Virtualroom study",
	}).Error; err != nil {
		log.Printf("Problems creating main categories. <database/populate.go->populateSchool>")
	}
	if err := Tiukudb.Create(&MainCategory{
		ID:        0,
		Shorthand: "Itse",
		Finnish:   "Itsenäinen opiskelu",
		English:   "Independent Study",
	}).Error; err != nil {
		log.Printf("Problems creating main categories. <database/populate.go->populateSchool>")
	}
}

// Auto-generating student users for testing purposes
func PopulateStudents(p int) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	i := 0
	for i < p {
		i = i + 1
		// Switching auto-generated classes
		classToAdd := ""
		if i%2 == 0 {
			classToAdd = "tit2"
		} else {
			classToAdd = "tit1"
		}
		_, tempAnon := CreateNewAnonID()
		//if err := db.Table(schoolShortName + "_StudentUsers").Create(&StudentUser{
		if err := Tiukudb.Create(&StudentUser{
			ID:          0,
			StudentID:   "oppi" + strconv.Itoa(i),
			AnonID:      tempAnon,
			StudentName: "Oppilas " + strconv.Itoa(i),
			//StudentSegments: StudentSegment{},
			StudentSegments: "",
			StudentEmail:    "oppilas" + strconv.Itoa(i) + "@oppilaitos.fi",
			StudentClass:    classToAdd,
		}).Error; err != nil {
			log.Println("Problems populating table of StudentUsers. <database/populate.go->populateStudents>")
		}

	}

}

func AutoCreateStudentUserTables() {
	if Tiukudb == nil {
		ConnectToDB()
	}
	numberOfStudentUsers := CountTableRows(StudentsTableToEdit)
	i := 1
	for i < (numberOfStudentUsers + 1) {
		newStudent := GetStudentUserWithStudentID("oppi" + strconv.Itoa(i))
		CreateStudentSegmentTable(newStudent)
		CreateActiveSegmentSessionsTable(newStudent)
		CreateSegmentsSessionsArchive(newStudent)
		i++
	}
}

// Auto-generating courses for testing purposes
func PopulateCourses(p int) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	i := 0
	for i < p {
		i++
		// Auto-generating archived status
		archivedToAdd := false
		if i%2 == 0 {
			archivedToAdd = true
		} else {
			archivedToAdd = false
		}
		//if err := db.Table(schoolShortName + "_Courses").Create(&Course{
		if err := Tiukudb.Create(&Course{
			ID: 0,
			//ResourceID:      0,
			Degree:          1,
			CourseCode:      "GTC" + strconv.Itoa(i),
			CourseName:      "Generated Test Course " + strconv.Itoa(i),
			CourseStartDate: strconv.Itoa(i) + "." + strconv.Itoa(i) + ".2020",
			CourseEndDate:   strconv.Itoa(i) + "." + strconv.Itoa(i) + ".2021",
			Archived:        archivedToAdd,
			//Segment:         []Segment{},
		}).Error; err != nil {
			log.Println("Problems populating Courses table. <database/populate.go->populateCourses>")
		}
		//tempFaculty := GetFacultyUser(strconv.Itoa(i))
		//CreateFacultySegmentTable(tempFaculty)
	}

}

//	AutoCreateSegments for Courses
//	comment: Modified code from CreateSegments
func AutoCreateSegments() {
	if Tiukudb == nil {
		ConnectToDB()
	}
	numberOfCourses := CountTableRows(CourseTableToEdit)
	i := 1
	for i < numberOfCourses {
		courseToAdd := GetCourseTableById(scripts.IntToUint(i))
		log.Print(courseToAdd.ID)
		c := 1
		for c < 7 {
			newSegment := &Segment{
				ID:          0,
				CourseID:    courseToAdd.ID,
				SegmentName: "segment " + strconv.Itoa(c),
				TeacherID:   scripts.IntToUint(c),
				Scope:       3,
				//SegmentCategories:     "", //SegmentCategory{},
				ExpectedAttendance: 15,
				//SchoolSegmentsSession: SchoolSegmentsSession{},
			}
			c++
			newSegment.CourseID = courseToAdd.ID
			Tiukudb.Model(&courseToAdd).Association("Segment").Append(newSegment)
			Tiukudb.Save(&courseToAdd)
			// Re-thinkin about categories, maybe only create when using other than 3 stock ones?
			//newSegment.SegmentCategories = AutoCreateCategoriesForSegment(newSegment.ID)
		}

		i++
	}
	//return courseToAdd
}

// Auto creating catergories for segments
// status: works, but decided to go with one shared table for categories.
func AutoCreateCategoriesForSegment(segmentToAdd uint) string { //segmentToAdd Segment) string {
	if Tiukudb == nil {
		ConnectToDB()
	}

	segmentID := segmentToAdd //.ID
	log.Println(segmentID)
	tableToCreate := strconv.FormatUint(uint64(segmentID), 10) + "_categories"
	if err := Tiukudb.Table(tableToCreate).AutoMigrate(&SegmentCategory{
		ID:                 0,
		MainCategory:       0,
		SubCategory:        "",
		MandatoryToTrack:   false,
		MandatoryToComment: false,
		Tickable:           false,
		LocationNeeded:     false,
		Active:             false,
	}).Error; err != nil {
		log.Println("Problems creating categories table for segment. <database/database_create->AutoCreateCategoriesForSegment>")
	}

	Tiukudb.Table(tableToCreate).AddForeignKey("main_category", "main_categories(id)", "RESTRICT", "RESTRICT")
	return tableToCreate
}

// Auto Populate categories with test categories.
func PopulateCategories() {
	if Tiukudb == nil {
		ConnectToDB()
	}
	c := CountTableRows(SegmentTableToEdit)
	i := 1
	for i < c {
		if err := Tiukudb.Create(&SegmentCategory{
			ID:                 0,
			SegmentID:          uint(i),
			MainCategory:       1,
			SubCategory:        "Lähi tunti " + strconv.Itoa(i),
			MandatoryToTrack:   false,
			MandatoryToComment: false,
			Tickable:           false,
			LocationNeeded:     false,
			Active:             true,
			Archived:           false,
		}).Error; err != nil {
			log.Println("Problems populating categories table. <database/populate.go->populateCategories>")
		}
		if err := Tiukudb.Create(&SegmentCategory{
			ID:                 0,
			SegmentID:          uint(i),
			MainCategory:       2,
			SubCategory:        "Videoluento " + strconv.Itoa(i),
			MandatoryToTrack:   false,
			MandatoryToComment: false,
			Tickable:           false,
			LocationNeeded:     false,
			Active:             true,
			Archived:           false,
		}).Error; err != nil {
			log.Println("Problems populating categories table. <database/populate.go->populateCategories>")
		}
		if err := Tiukudb.Create(&SegmentCategory{
			ID:                 0,
			SegmentID:          uint(i),
			MainCategory:       3,
			SubCategory:        "Kotitehtävä " + strconv.Itoa(i),
			MandatoryToTrack:   false,
			MandatoryToComment: false,
			Tickable:           false,
			LocationNeeded:     false,
			Active:             true,
			Archived:           false,
		}).Error; err != nil {
			log.Println("Problems populating categories table. <database/populate.go->populateCategories>")
		}
		i++
	}

}

// Testing purposes generates faculty users
//
func PopulateFaculty(p int) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	i := 0
	for i < p {
		i = i + 1
		//if err := db.Table(schoolShortName + "_StudentUsers").Create(&StudentUser{
		if err := Tiukudb.Create(&FacultyUser{
			ID:           0,
			FacultyID:    "ope" + strconv.Itoa(i),
			FacultyName:  "opettaja" + strconv.Itoa(i),
			FacultyEmail: "opettaja" + strconv.Itoa(i) + "@oppilaitos.fi",
			Apartment:    1,
			Active:       true,
			Teacher:      true,
			Admin:        true,
		}).Error; err != nil {
			log.Println("Problems populating table of StudentUsers. <database/populate.go->populateStudents>")
		}
	}
}

// Auto create table for Faculty Users
// status: works, but not in use
/*
func AutoCreateFacultyUserTables() {
	if Tiukudb == nil {
		ConnectToDB()
	}
	numberOfFacultyUsers := CountTableRows(facultyTableToEdit)
	i := 1
	for i < numberOfFacultyUsers {
		newFaculty := GetFacultyUser("ope" + strconv.Itoa(i))
		CreateFacultySegmentTable(newFaculty)
		i++
	}
}
*/
