package database

/*
// create_database.go
// Description: Creating tables on database
*/
import (
	"log"

	"github.com/AskoTJM/tiuku/api/scripts"
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
	if DebugMode {
		log.Printf("Anon Id is: %s", student.StudentID)
	}
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
		if DebugMode {
			log.Println(tempStudent)
		}
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

// Create Segment table for new Faculty users
// Status: No clue, just copy+pasted and edited from CreateStudentSegmentTable
// Not in use. Decided to go with one table for all faculty users
/*
func CreateFacultySegmentTable(newFaculty FacultyUser) string {
	if Tiukudb == nil {
		ConnectToDB()
	}
	tableToEdit := newFaculty.FacultyID + "_segments"
	result := Tiukudb.HasTable(tableToEdit)
	if result {
		log.Println("Error: Table already exists. <database/create_database.go->CreateFacultySegmentTable>")
		return "Error: Table already exists."
	} else {
		if err := Tiukudb.Table(tableToEdit).AutoMigrate(&FacultySegment{
			ID:                    0,
			Course:                Course{},
			SegmentNumber:         0,
			SchoolSegmentsSession: SchoolSegmentsSession{},
			SegmentCategories:     SegmentCategory{},
			Archived:              false,
		}).Error; err != nil {
			log.Println("Problems creating Segment table of FacultyUsers. <database/create_database->CreateFacultySegmentTable>")
		}
		return tableToEdit
	}
}
*/

// Create new course in Courses table.
// Status: Working, but not finished. Needs checking.
func CreateCourse(newCourse Course, tableToEdit string) string {
	// Check if there is connection to database if not connect to it
	if Tiukudb == nil {
		ConnectToDB()
	}

	Tiukudb.Table(tableToEdit).Create(&newCourse)

	response := "Course created" + newCourse.CourseCode
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
func CreateActiveSegmentSessionsTable(user StudentUser) string {
	if Tiukudb == nil {
		ConnectToDB()
	}
	//var newTable StudentSegmentSession
	tableToCreate := user.AnonID + "_sessions"

	if err := Tiukudb.Table(tableToCreate).AutoMigrate(&StudentSegmentSession{}).Error; err != nil {
		log.Printf("Problems creating active Session table for student user. <database/create_database->CreateActiveSegmentSessionsTable> Error: %v \n", err)
	}

	return tableToCreate
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
func StartSessionOnSegment(student string, newSession StudentSegmentSession) bool {
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
