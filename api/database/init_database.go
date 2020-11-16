package database

/*
// init_database.go
// Description: Initializing database for use.
*/
import "log"

// initDB() for creating needed tables for database
func InitDB() {
	if Tiukudb == nil {
		ConnectToDB()
	}

	if DebugMode {
		log.Printf("Trying to AutoMigrate Course table to database. <database/maintenance_database.go->initDB>")
		//if err := db.Table(schoolShortName + "_Courses").AutoMigrate(&Course{
	}
	//if err := db.Table(schoolShortName + "_Courses").AutoMigrate(&Course{
	if err := Tiukudb.AutoMigrate(&Course{
		ID: 0,
		//ResourceID:      0,
		Degree:          0,
		CourseCode:      "",
		CourseName:      "",
		CourseStartDate: "",
		CourseEndDate:   "",
		Archived:        false,
		//Segment:         []Segment{},
	}).Error; err != nil {
		log.Println("Problems creating table for Course. <database/maintenance_database.go->initDB>")
	}

	// This should be created and named when creating Segment
	if DebugMode {
		log.Printf("Trying to AutoMigrate Segment table to database. <database/maintenance_database.go->initDB>")
	}
	if err := Tiukudb.AutoMigrate(&Segment{
		ID:          0,
		CourseID:    0,
		SegmentName: "",
		TeacherID:   0,
		Scope:       0,
		//SegmentCategories:     "", //SegmentCategory{},
		ExpectedAttendance: 0,
		//: SchoolSegmentsSession{},
	}).Error; err != nil {
		log.Println("Problems creating table for Segment. <database/maintenance_database.go->initDB>")
	}

	// Creating MainCategory Table, Mandatory options in all Segments
	if err := Tiukudb.AutoMigrate(&MainCategory{
		ID:        0,
		Shorthand: "",
		Finnish:   "",
		English:   "",
	}).Error; err != nil {
		log.Println("Problems creating Main Category table. <database/maintenance_database.go->initDB>")
	}

	// Table for SegmentCategories
	if err := Tiukudb.AutoMigrate(&SegmentCategory{
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
	// Create table for used for tracking participation to Segments
	if err := Tiukudb.AutoMigrate(&SchoolSegmentsSession{
		ID:                      0,
		SegmentID:               0,
		AnonID:                  "",
		StudentSegmentsSessions: "",
		Privacy:                 "",
	}).Error; err != nil {
		log.Println("Problems creating table for Participant Session table. <database/maintenance_database.go->initDB>")
	}
	// Create table for Archiving Sessions
	if err := Tiukudb.AutoMigrate(&ArchivedSessionsTable{
		ID:                 0,
		SchoolID:           0,
		CampusID:           0,
		ApartmentID:        0,
		DegreeID:           0,
		CourseCode:         "",
		CourseName:         "",
		CourseStartDate:    "",
		CourseEndDate:      "",
		SegmentName:        "",
		TeacherID:          0,
		Scope:              0,
		ExpectedAttendance: 0,
		MainCategory:       0,
		SubCategoryID:      0,
		SubCategory:        "",
		MandatoryToTrack:   false,
		MandatoryToComment: false,
		Tickable:           false,
		AnonID:             "",
		StartTime:          "",
		EndTime:            "",
		Created:            "",
		Updated:            "",
		Deleted:            "",
		Comment:            "",
		Version:            0,
		Locations:          "",
		Privacy:            false,
	}).Error; err != nil {
		log.Println("Problems creating table for Degrees. <database/maintenance_database.go->initDB>")
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
	if err := Tiukudb.AutoMigrate(&StudentUser{
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
	if err := Tiukudb.AutoMigrate(&FacultyUser{
		ID:           0,
		FacultyID:    "",
		FacultyName:  "",
		FacultyEmail: "",
		//FacultySegment: FacultySegment{},
		//FacultySegment: "",
	}).Error; err != nil {
		log.Println("Problems creating table for FacultyUsers. <database/maintenance_database.go->initDB>")
	}
	// Tables for School data

	log.Printf("Trying to AutoMigrate Schools table to database. <database/maintenance_database.go->initDB>")
	if err := Tiukudb.AutoMigrate(&School{
		ID:        0,
		Shorthand: "",
		Finnish:   "",
		English:   "",
		Campuses:  []Campus{},
	}).Error; err != nil {
		log.Println("Problems creating table for School. <database/maintenance_database.go->initDB>")
	}
	//if err := db.Table(schoolShortName + "_Campuses").AutoMigrate(&Campus{
	if err := Tiukudb.AutoMigrate(&Campus{
		ID:         0,
		Shorthand:  "",
		Finnish:    "",
		English:    "",
		Apartments: []Apartment{},
	}).Error; err != nil {
		log.Println("Problems creating table for Campuses. <database/maintenance_database.go->initDB>")
	}
	//if err := db.Table(schoolShortName + "_Apartments").AutoMigrate(&Apartment{
	if err := Tiukudb.AutoMigrate(&Apartment{
		ID:        0,
		Shorthand: "",
		Finnish:   "",
		English:   "",
		Degrees:   []Degree{},
	}).Error; err != nil {
		log.Println("Problems creating table for Apartments. <database/maintenance_database.go->initDB>")
	}
	//if err := db.Table(schoolShortName + "_Degrees").AutoMigrate(&Degree{
	if err := Tiukudb.AutoMigrate(&Degree{
		ID:        0,
		Shorthand: "",
		Finnish:   "",
		English:   "",
	}).Error; err != nil {
		log.Println("Problems creating table for Degrees. <database/maintenance_database.go->initDB>")
	}

	initData()

}

// inital data for tables
func initData() {
	initSchool()
	initMainCategoryTable()
	initSegmentCategoryTable()
}

func initMainCategoryTable() {
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

func initSegmentCategoryTable() {
	if err := Tiukudb.Create(&SegmentCategory{
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
	if err := Tiukudb.Create(&SegmentCategory{
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
	if err := Tiukudb.Create(&SegmentCategory{
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
	if err := Tiukudb.Create(&School{
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
