package database

/*
// create_database.go
// Description: Creating tables on database
*/
import (
	"log"
	"net/http"

	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/dchest/uniuri"
)

// For creating Segments table for new Student users and adding it to student_user list
// Status: Works
func CreateStudentSegmentTable(student StudentUser) string {
	if Tiukudb == nil {
		ConnectToDB()
	}
	// Get Students data
	tempStudent := student //.StudentID)
	// Get AnonID for data
	myAnonID := tempStudent.AnonID
	tableToEdit := myAnonID + "_segments"
	result := Tiukudb.HasTable(tableToEdit)
	if result {
		log.Println("Error: Table already exists. <database/database_maintenance.go->CreateStudentSegmentTable>")
		return "Error: Table already exists."
	} else {
		if err := Tiukudb.Table(tableToEdit).AutoMigrate(&StudentSegment{
			ID: 0,
			//Course:    Course{},
			SegmentID: 0,
			//StudentSegmentSessions: StudentSegmentSession{},
			//SegmentCategory:        SegmentCategory{},
			StudentSegmentSessions: "",
			//SegmentCategory:        "",
			Archived: false,
		}).Error; err != nil {
			log.Println("Problems creating Segment table of StudentUsers. <database/database_create->CreateStudentSegmentTable>")
		}
		// Update the Student data with the name of the segment table
		Tiukudb.Model(&tempStudent).Where("student_id = ? ", tempStudent.StudentID).Update("student_segments", tableToEdit)
		return tableToEdit

	}
}

// For creating Archive Segment table for containing old segments
// Status: Works
// comment: Most likely unnecessary. Amount of segments for one student user shouldn't be that much
// that we need another table for archiving.
func CreateStudentSegmentTableArchived(newStudent StudentUser) string {
	if Tiukudb == nil {
		ConnectToDB()
	}
	// Get Students data
	//tempStudent := GetStudent(StudentID)
	// Get AnonID for data
	myAnonID := newStudent.AnonID
	tableToEdit := myAnonID + "_segments_archived"
	// Check if the table already exists
	result := Tiukudb.HasTable(tableToEdit)
	if result {
		log.Println("Error: Table already exists. <database/database_maintenance.go->CreateStudentSegmentTable>")
		return "Error: Table already exists."
	} else {
		if err := Tiukudb.Table(tableToEdit).AutoMigrate(&StudentSegment{
			ID:                     0,
			SegmentID:              0,
			StudentSegmentSessions: "",
			Archived:               false,
			//Course:    Course{},
			//StudentSegmentSessions: StudentSegmentSession{},
			//SegmentCategory:        SegmentCategory{},
			//SegmentCategory:        "",
		}).Error; err != nil {
			log.Println("Problems creating Segment table of StudentUsers. <database/database_create->CreateStudentSegmentTable>")
		}
		// Update the Student data with the name of the segment table
		// db.Model(&tempStudent).Update("student_segments", tableToEdit)
		return tableToEdit
	}
}

// Create new course in Courses table.
// Status: Working, but not finished. Needs checking.
func CreateCourse(newCourse Course, tableToEdit string) string {
	// Check if there is connection to database if not connect to it
	if Tiukudb == nil {
		ConnectToDB()
	}

	Tiukudb.Table(tableToEdit).Create(&newCourse)

	response := "Course created " + newCourse.CourseCode
	return response
}

// Create new Segment for course
// Status: works
func CreateSegment(newSegment Segment, tableToEdit string) string {
	if Tiukudb == nil {
		ConnectToDB()
	}
	getCourseData := GetCourseTableById(newSegment.CourseID)
	Tiukudb.Model(&getCourseData).Association("Segment").Append(newSegment)
	Tiukudb.Save(&getCourseData)
	response := "Segment created " + newSegment.SegmentName
	return response //getCourseData
}

// Create new Category for Segment, return True on success, else False
func CreateCategory(newCategory SegmentCategory, tableToEdit string) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var response bool
	if err := Tiukudb.Table(tableToEdit).Save(&newCategory).Error; err != nil {
		log.Printf("Problems creating new Category for segment. <database/create_database->CreateCategory> Error: %v \n", err)
		response = false
	} else {
		log.Printf("Succesfully created new Category %v for Segment %v \n", newCategory.ID, newCategory.SegmentID)
		response = true
	}
	return response
}

// Create List to

func CreateSchoolSegmentSession(newSeg Segment) string {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var returnString string
	tableToCreate := scripts.UintToString(newSeg.ID) + "_session"
	if err := Tiukudb.Table(tableToCreate).AutoMigrate(&SchoolSegmentsSession{
		ID:                      0,
		AnonID:                  "",
		StudentSegmentsSessions: "",
		Privacy:                 "",
	}).Error; err != nil {
		log.Printf("Problems creating School sessions list for segment. <database/create_database->CreateSchoolSegmentsList>Error: %v \n", err)
	}

	return returnString
}

// Create SegmentSessionTable for active segments for new student user
func CreateActiveSegmentSessionsTable(user StudentUser) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	//var newTable StudentSegmentSession
	var response bool
	tableToCreate := user.AnonID + "_sessions"

	if err := Tiukudb.Table(tableToCreate).AutoMigrate(&StudentSegmentSession{}).Error; err != nil {
		log.Printf("Problems creating active Session table for student user. <database/create_database->CreateActiveSegmentSessionsTable> Error: %v \n", err)
		response = false
	} else {
		response = true
	}

	return response
}

// Create Archive SegmentSessionTable for student user
func CreateSegmentsSessionsArchive(user StudentUser) string {
	if Tiukudb == nil {
		ConnectToDB()
	}
	tableToCreate := user.AnonID + "_sessions_archived"

	if err := Tiukudb.Table(tableToCreate).AutoMigrate(&StudentSegmentSession{}).Error; err != nil {
		log.Printf("Problems creating active Session table for student user. <database/create_database->CreateActiveSegmentSessionsTable>Error: % \n", err)
	}

	return tableToCreate
}

// Add/Start Session, Returns True on success  / False on Error
// W0rks
func CreateNewSessionOnSegment(student string, newSession StudentSegmentSession) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	studentNow := GetStudentUser(student)
	var response bool
	tableToEdit := studentNow.AnonID + "_sessions"
	if err := Tiukudb.Table(tableToEdit).Create(&newSession).Error; err != nil {
		log.Printf("Error in starting Session %v \n", err)
		response = false
	} else {
		//newSession.ResourceID = newSession.ID
		if err2 := Tiukudb.Table(tableToEdit).Where("id = ?", newSession.ID).Updates(StudentSegmentSession{ResourceID: newSession.ID}).Error; err2 != nil {
			log.Printf("Error setting ResourceID for starting Session %v \n", err)
			response = false
			return response
		}
		response = true
	}
	return response
}

// Create new Student User
// W0rks
func CreateNewStudentUser(newStudent StudentUser) (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int
	var responseString string
	err2, resString := CreateNewAnonID()

	if err2 != http.StatusCreated {
		responseCode = http.StatusInternalServerError
		responseString = resString
	} else {
		newStudent.AnonID = resString
		newStudent.StudentSegments = newStudent.AnonID + "_segments"
		if err := Tiukudb.Table(StudentsTableToEdit).Create(&newStudent).Error; err != nil {
			responseCode = http.StatusInternalServerError
			responseString = "Error creating new student user. "
		} else {
			CreateStudentSegmentTable(newStudent)
			CreateActiveSegmentSessionsTable(newStudent)
			CreateSegmentsSessionsArchive(newStudent)
			responseCode = http.StatusOK
			responseString = "Created new student user."

		}

	}
	return responseCode, responseString
}

// Create new Faculty User
// W0rks?
func CreateNewFacultyUser(newFaculty FacultyUser) (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int
	var responseString string

	if err := Tiukudb.Table(FacultyTableToEdit).Create(&newFaculty).Error; err != nil {
		responseCode = http.StatusInternalServerError
		responseString = "Error creating new Faculty user. "
	} else {
		responseCode = http.StatusOK
		responseString = "Created new Faculty user."
	}
	return responseCode, responseString
}

// Create new AnonID, generates, check if AnonID is already in use, tries 5 times to generate unique
// W0rks
func CreateNewAnonID() (int, string) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var responseCode int = 0
	var responseString string
	c := 0
	for c < 5 {
		newAnon := "AID:" + uniuri.NewLen(20)
		//newAnon := "IOEi17L6abYv4SaT84l2"
		// Check if AnonID already exists
		err := GetStudentUserWithAnonID(newAnon)
		if err == (StudentUser{}) {
			//log.Println("Generating AnonID")
			responseCode = http.StatusCreated
			responseString = newAnon
			c = 6
		} else {
			log.Printf("Error with generating Anon ID. ID already exists <database/create_database.go->CreateNewAnonID")
			//log.Printf("C is now %v", c)
			responseCode = http.StatusInternalServerError
			responseString = "Problems witht the Server"
			c = c + 1
		}
	}

	return responseCode, responseString
}
