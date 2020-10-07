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
	if debugMode {
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
		if debugMode {
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

func CreateCategory(newCategory SegmentCategory, tableToEdit string) string {
	if Tiukudb == nil {
		ConnectToDB()
	}
	/*
		if !CheckIfRequiredTablesExist() {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			response := "Problems creating new Category, required tables do not exist. <database/create_database->func CreateCategory>"
			return response
		} else {
			// Check if content type is set.
			if r.Header.Get("Content-Type") == "" {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusNoContent)
				response := "Problems creating new Category, no body in request information. <database/create_database->CreateCategory> Error: No body information available."
				return response
			} else {
				//rbody, _ := header.ParseValueAndParams(r.Header, "Content-Type")
				rbody := r.Header.Get("Content-Type")
				// Check if content type is correct one.
				if rbody != "application/json" {
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(http.StatusNotAcceptable)
					response := "Error: Content-Type is not application/json."
					return response
				}

			}

		}
		dec := json.NewDecoder(r.Body)

		dec.DisallowUnknownFields()
		//log.Println(dec)

		var newCategory SegmentCategory
		err := dec.Decode(&newCategory)
		if err != nil {
			log.Println("Problem with json decoding <database/create_database->CreateCategory")
		}

		vars := mux.Vars(r)
		segmentID := vars["segment"]
		if debugMode {
			log.Println(segmentID)
		}
		tempSegId := scripts.StringToUint(segmentID)
		//tableToCreate := segmentID + "_categories"
		if err := Tiukudb.Table(SegmentTableToEdit).AutoMigrate(&SegmentCategory{
			ID:                 0,
			SegmentID:          tempSegId,
			MainCategory:       newCategory.MainCategory,
			SubCategory:        newCategory.SubCategory,
			MandatoryToTrack:   newCategory.MandatoryToTrack,
			MandatoryToComment: newCategory.MandatoryToComment,
			Tickable:           newCategory.Tickable,
			LocationNeeded:     newCategory.LocationNeeded,
			Active:             newCategory.Active,
			Archived:           newCategory.Archived,
		}).Error; err != nil {
			log.Println("Problems creating categories table for segment. <database/database_create->CreateCategories>")
		}

		//db.Table(segmentID).AddForeignKey("main_category", "main_categories(id)", "RESTRICT", "RESTRICT")
		// For some reason have to manually set the
		newCategory.SegmentID = tempSegId
	*/
	Tiukudb.Save(&newCategory)
	response := "Created categories."
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
		log.Println("Problems creating School sessions list for segment. <database/create_database->CreateSchoolSegmentsList>")
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
		log.Println("Problems creating active Session table for student user. <database/create_database->CreateActiveSegmentSessionsTable>")
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
		log.Println("Problems creating active Session table for student user. <database/create_database->CreateActiveSegmentSessionsTable>")
	}

	return tableToCreate
}
