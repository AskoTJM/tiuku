package database

import (
	"log"
	"strconv"
	//"github.com/AskoTJM/tiuku/api/database"
)

// Should only contain scripts to autopopulate data for development and testing
// Place for scripts to initalization and for populating database with test data
// Stuff that should not be needed when in use.

// OBSOLETE! Replaced by initDB scripts.
// desc: Populating School data and maincategories
func PopulateSchool() {
	if db == nil {
		ConnectToDB()
	}

	if err := db.Create(&School{
		ID:        0,
		Shorthand: "OAMK",
		Finnish:   "Oulun Ammattikorkeakoulu",
		English:   "Oulu University of Applied Sciences",
		Campuses: []Campus{{
			ID:        0,
			Shorthand: "Linna",
			Finnish:   "Linnanmaan Kampus",
			English:   "Campus Linnanmaa",
			Apartments: []Apartment{{
				ID:        0,
				Shorthand: "ICT",
				Finnish:   "Informaatioteknologia",
				English:   "Information Technology",
				Degrees: []Degree{{
					ID:        0,
					Shorthand: "bEng",
					Finnish:   "Insinööri (AMK), tieto- ja viestintätekniikka",
					English:   "Bachelor of Engineering, Information Technology",
				}},
			}},
		}},
	}).Error; err != nil {
		log.Printf("Problems populating table of Schools. <database/populate.go->populateSchool>")
	}

	if err := db.Create(&MainCategory{
		ID:        0,
		Shorthand: "Lähi",
		Finnish:   "Lähiopetus",
		English:   "Classroom study",
	}).Error; err != nil {
		log.Printf("Problems creating main categories. <database/populate.go->populateSchool>")
	}
	if err := db.Create(&MainCategory{
		ID:        0,
		Shorthand: "Etä",
		Finnish:   "Etäopetus",
		English:   "Virtualroom study",
	}).Error; err != nil {
		log.Printf("Problems creating main categories. <database/populate.go->populateSchool>")
	}
	if err := db.Create(&MainCategory{
		ID:        0,
		Shorthand: "Itse",
		Finnish:   "Itsenäinen opiskelu",
		English:   "Independent Study",
	}).Error; err != nil {
		log.Printf("Problems creating main categories. <database/populate.go->populateSchool>")
	}
}

// desc: Auto-generating student users for testing purposes
func PopulateStudents(p int) {
	if db == nil {
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
		//if err := db.Table(schoolShortName + "_StudentUsers").Create(&StudentUser{
		if err := db.Create(&StudentUser{
			ID:          0,
			StudentID:   "oppi" + strconv.Itoa(i),
			AnonID:      "Anon" + strconv.Itoa(i),
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
	if db == nil {
		ConnectToDB()
	}
	numberOfStudentUsers := CountStudentUsers()
	i := 1
	for i < (numberOfStudentUsers + 1) {
		newStudent := GetStudentUser("oppi" + strconv.Itoa(i))
		if debugMode {
			log.Printf("newStudent has value of: %d", i)
			log.Printf("newStudent AnonID is : %s", newStudent.AnonID)
		}
		CreateStudentSegmentTable(newStudent)
		CreateActiveSegmentSessionsTable(newStudent)
		CreateSegmentsSessionsArchive(newStudent)
		i++
	}
}

// desc: Auto-generating courses for testing purposes
func PopulateCourses(p int) {
	if db == nil {
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
		if err := db.Create(&Course{
			ID: 0,
			//ResourceID:      0,
			Degree: Degree{
				ID:          0,
				ApartmentID: 0,
				Shorthand:   "",
				Finnish:     "",
				English:     "",
			},
			CourseCode:      "GTC" + strconv.Itoa(i),
			CourseName:      "Generated Test Course " + strconv.Itoa(i),
			CourseStartDate: strconv.Itoa(i) + "." + strconv.Itoa(i) + ".2020",
			CourseEndDate:   strconv.Itoa(i) + "." + strconv.Itoa(i) + ".2021",
			Archived:        archivedToAdd,
			Segment:         []Segment{},
		}).Error; err != nil {
			log.Println("Problems populating Courses table. <database/populate.go->populateCourses>")
		}
		//tempFaculty := GetFacultyUser(strconv.Itoa(i))
		//CreateFacultySegmentTable(tempFaculty)
	}

}

//	desc: AutoCreateSegments for Courses
//	comment: Modified code from CreateSegments
func AutoCreateSegments() {
	if db == nil {
		ConnectToDB()
	}
	numberOfCourses := CountCourses()
	i := 1
	for i < numberOfCourses {
		courseToAdd := GetCourseTableById(strconv.Itoa(i))
		log.Print(courseToAdd.ID)
		c := 1
		for c < 4 {
			newSegment := &Segment{
				ID:          0,
				CourseID:    courseToAdd.ID,
				SegmentName: "segment " + strconv.Itoa(c),
				TeacherID:   0,
				Scope:       3,
				//SegmentCategories:     "", //SegmentCategory{},
				ExpectedAttendance:    15,
				SchoolSegmentsSession: SchoolSegmentsSession{},
			}
			c++
			newSegment.CourseID = courseToAdd.ID
			db.Model(&courseToAdd).Association("Segment").Append(newSegment)
			db.Save(&courseToAdd)
			// Re-thinkin about categories, maybe only create when using other than 3 stock ones?
			//newSegment.SegmentCategories = AutoCreateCategoriesForSegment(newSegment.ID)
		}

		i++
	}
	//return courseToAdd
}

// desc: Auto creating catergories for segments
// status: works, but decided to go with one shared table for categories.
func AutoCreateCategoriesForSegment(segmentToAdd uint) string { //segmentToAdd Segment) string {
	if db == nil {
		ConnectToDB()
	}

	segmentID := segmentToAdd //.ID
	log.Println(segmentID)
	tableToCreate := strconv.FormatUint(uint64(segmentID), 10) + "_categories"
	if err := db.Table(tableToCreate).AutoMigrate(&SegmentCategory{
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

	db.Table(tableToCreate).AddForeignKey("main_category", "main_categories(id)", "RESTRICT", "RESTRICT")
	return tableToCreate
}

// desc: Auto Populate categories with test categories.
func PopulateCategories() {
	if db == nil {
		ConnectToDB()
	}
	c := CountSegments()
	i := 1
	for i < c {
		if err := db.Create(&SegmentCategory{
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
		if err := db.Create(&SegmentCategory{
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
		if err := db.Create(&SegmentCategory{
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

// desc: Testing purposes generates faculty users
//
func PopulateFaculty(p int) {
	if db == nil {
		ConnectToDB()
	}
	i := 0
	for i < p {
		i = i + 1
		//if err := db.Table(schoolShortName + "_StudentUsers").Create(&StudentUser{
		if err := db.Create(&FacultyUser{
			ID:           0,
			FacultyID:    "ope" + strconv.Itoa(i),
			FacultyName:  "opettaja" + strconv.Itoa(i),
			FacultyEmail: "opettaja" + strconv.Itoa(i) + "@oppilaitos.fi",
			//School:         School{},
			Apartment:      Apartment{},
			FacultySegment: "", //CreateFacultySegmentTable("ope" + strconv.Itoa(i))"",
			//FacultySegment: FacultySegment{},
		}).Error; err != nil {
			log.Println("Problems populating table of StudentUsers. <database/populate.go->populateStudents>")
		}
	}
}

// desc: Auto create table for Faculty Users
func AutoCreateFacultyUserTables() {
	if db == nil {
		ConnectToDB()
	}
	numberOfFacultyUsers := CountFacultyUsers()
	i := 1
	for i < numberOfFacultyUsers {
		newFaculty := GetFacultyUser("ope" + strconv.Itoa(i))
		CreateFacultySegmentTable(newFaculty)
		i++
	}
}
