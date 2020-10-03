package database

import "log"

// initDB() for creating needed tables for database
func InitDB() {
	if db == nil {
		ConnectToDB()
	}

	if debugMode {
		log.Printf("Trying to AutoMigrate Course table to database. <database/maintenance_database.go->initDB>")
		//if err := db.Table(schoolShortName + "_Courses").AutoMigrate(&Course{
	}
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
	if debugMode {
		log.Printf("Trying to AutoMigrate Segment table to database. <database/maintenance_database.go->initDB>")
	}
	if err := db.AutoMigrate(&Segment{
		ID:          0,
		CourseID:    0,
		SegmentName: "",
		TeacherID:   0,
		Scope:       0,
		//SegmentCategories:     "", //SegmentCategory{},
		ExpectedAttendance:    0,
		SchoolSegmentsSession: SchoolSegmentsSession{},
	}).Error; err != nil {
		log.Println("Problems creating table for Segment. <database/maintenance_database.go->initDB>")
	}

	// desc: Creating MainCategory Table
	if err := db.AutoMigrate(&MainCategory{
		ID:        0,
		Shorthand: "",
		Finnish:   "",
		English:   "",
	}).Error; err != nil {
		log.Println("Problems creating Main Category table. <database/maintenance_database.go->initDB>")
	}

	// desc: Table for SegmentCategories
	if err := db.AutoMigrate(&SegmentCategory{
		ID:                 0,
		MainCategory:       0,
		SubCategory:        "",
		MandatoryToTrack:   false,
		MandatoryToComment: false,
		Tickable:           false,
		LocationNeeded:     false,
		Active:             false,
	}).Error; err != nil {
		log.Println("Problems creating table for Categories. <database/maintenance_database.go->initDB>")
	}

	if err := db.AutoMigrate(&SchoolSegmentsSession{
		ID:                      0,
		SegmentID:               0,
		AnonID:                  "",
		StudentSegmentsSessions: "",
		Privacy:                 "",
	}).Error; err != nil {
		log.Println("Problems creating table for Participant Session table. <database/maintenance_database.go->initDB>")
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

	initData()

}

// desc: inital data for tables
func initData() {
	initSchool()
	initMainCategoryTable()
	initSegmentCategoryTable()
}

func initMainCategoryTable() {
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

func initSegmentCategoryTable() {
	if err := db.Create(&SegmentCategory{
		ID:                 0,
		SegmentID:          0,
		MainCategory:       1,
		SubCategory:        "Lähi",
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
		SegmentID:          0,
		MainCategory:       2,
		SubCategory:        "Etä",
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
		SegmentID:          0,
		MainCategory:       3,
		SubCategory:        "Itse",
		MandatoryToTrack:   false,
		MandatoryToComment: false,
		Tickable:           false,
		LocationNeeded:     false,
		Active:             true,
		Archived:           false,
	}).Error; err != nil {
		log.Println("Problems populating categories table. <database/populate.go->populateCategories>")
	}
}

func initSchool() {
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
}
