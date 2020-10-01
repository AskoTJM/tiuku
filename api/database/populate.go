package database

import (
	"log"
	"strconv"
	//"github.com/AskoTJM/tiuku/api/database"
)

// Place for scripts to initalization and for populating database with test data
// Stuff that should not be needed when in use.
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
		tempStudent := GetStudentUser(strconv.Itoa(i))
		CreateStudentSegmentTable(tempStudent)

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
		tempFaculty := GetFacultyUser(strconv.Itoa(i))
		CreateFacultySegmentTable(tempFaculty)
	}

}

//	desc: AutoCreateSegments for Courses
//	comment: Couldn't use
func AutoCreateSegments() {
	if db == nil {
		ConnectToDB()
	}
	numberOfCourses := CountCourses()
	i := 1
	for i < numberOfCourses {
		courseToAdd := FindCourseTableById(strconv.Itoa(i))
		log.Print(courseToAdd.ID)
		c := 1
		for c < 4 {
			newSegment := &Segment{
				ID:                    0,
				CourseID:              courseToAdd.ID,
				SegmentName:           "segment " + strconv.Itoa(c),
				TeacherID:             0,
				Scope:                 3,
				SegmentCategories:     "", //SegmentCategory{},
				ExpectedAttendance:    15,
				SchoolSegmentsSession: SchoolSegmentsSession{},
			}
			c++
			newSegment.CourseID = courseToAdd.ID
			db.Model(&courseToAdd).Association("Segment").Append(newSegment)
			db.Save(&courseToAdd)
			newSegment.SegmentCategories = AutoCreateCategoriesForSegment(newSegment.ID)
		}

		i++
	}
	//return courseToAdd
}

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
			log.Println("Problems populating table of StudentUsers. <go/populate.go->populateStudents>")
		}
		//CreateFacultySegmentTable("ope" + strconv.Itoa(i))
	}

}
