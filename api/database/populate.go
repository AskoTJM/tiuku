package database

import (
	"log"
	"strconv"
	//"github.com/AskoTJM/tiuku/api/database"
)

// Place for scripts to initalization and for populating database with test data
// Stuff that should not be needed when in use.

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
		log.Printf("Problems populating table of Schools. <go/populate.go->populateSchool>")
	}

	if err := db.Create(&MainCategory{
		ID:       0,
		Shorthad: "Lähi",
		Finnish:  "Lähiopetus",
		English:  "Classroom study",
	}).Error; err != nil {
		log.Printf("Problems creating main categories. <go/populate.go->populateSchool>")
	}
	if err := db.Create(&MainCategory{
		ID:       0,
		Shorthad: "Etä",
		Finnish:  "Etäopetus",
		English:  "Virtualroom study",
	}).Error; err != nil {
		log.Printf("Problems creating main categories. <go/populate.go->populateSchool>")
	}
	if err := db.Create(&MainCategory{
		ID:       0,
		Shorthad: "Itse",
		Finnish:  "Itsenäinen opiskelu",
		English:  "Independent Study",
	}).Error; err != nil {
		log.Printf("Problems creating main categories. <go/populate.go->populateSchool>")
	}
}

// desc: Auto-generating student users for testing purposes
func PopulateStudents(p int) {
	if db == nil {
		ConnectToDB()
	}

	for i := 0; i < p; i = i + 1 {

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
			log.Println("Problems populating table of StudentUsers. <go/populate.go->populateStudents>")
		}
		CreateStudentSegmentTable("oppi" + strconv.Itoa(i))
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
			log.Println("Problems populating Courses table. <go/populate.go->populateCourses>")
		}
		courseToAdd := FindCourseTableById(strconv.Itoa(i))
		result := AutoCreateSegments(courseToAdd)
		log.Println(result.CourseName)
	}

}

// desc: Testing purposes generates faculty users
//
func PopulateFaculty(p int) {
	if db == nil {
		ConnectToDB()
	}

	for i := 0; i < p; i = i + 1 {

		//if err := db.Table(schoolShortName + "_StudentUsers").Create(&StudentUser{
		if err := db.Create(&FacultyUser{
			ID:           0,
			FacultyID:    "ope" + strconv.Itoa(i),
			FacultyName:  "opettaja" + strconv.Itoa(i),
			FacultyEmail: "opettaja" + strconv.Itoa(i) + "@oppilaitos.fi",
			//School:         School{},
			Apartment:      Apartment{},
			FacultySegment: "",
			//FacultySegment: FacultySegment{},
		}).Error; err != nil {
			log.Println("Problems populating table of StudentUsers. <go/populate.go->populateStudents>")
		}
		CreateFacultySegmentTable("ope" + strconv.Itoa(i))
	}

}

//	desc: AutoCreateSegments for Courses
//	comment: Couldn't use
func AutoCreateSegments(courseToAdd Course) Course {
	if db == nil {
		ConnectToDB()
	}
	//For what course is this
	var c = 0
	log.Printf("CourseCode is: %s", courseToAdd.CourseCode)
	//getCourseData := FindCourseTableById(courseToAdd.ID)
	for c < 3 {
		newSegment := &Segment{
			ID:                    0,
			CourseID:              courseToAdd.ID,
			SegmentName:           "segment " + strconv.Itoa(c),
			TeacherID:             0,
			Scope:                 3,
			SegmentCategories:     SegmentCategory{},
			ExpectedAttendance:    15,
			SchoolSegmentsSession: SchoolSegmentsSession{},
		}
		c++
		db.Model(&courseToAdd).Association("Segment").Append(newSegment)
		db.Save(&courseToAdd)
	}

	return courseToAdd
	//dec := json.NewDecoder(r.Body)
	/*
		newSeg.DisallowUnknownFields()
		log.Println(dec)

		var newSegment Segment
		err := dec.Decode(&newSegment)
		if err != nil {
			log.Println("Problem with json decoding <database/database_create->CreateSegment")
		}*/
	//getCourseData.Segment[0] = newSegment

}
