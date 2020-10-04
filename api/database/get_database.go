package database

/*
// get_database.go
// Description: Code for retrieving data from database
*/
import (
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Desc: Get Students data
// Status: Works, but needs more. Return value and obfuscing of AnonID if used outside
func GetStudentUser(StudentID string) StudentUser {
	if db == nil {
		ConnectToDB()
	}

	var tempStudent StudentUser

	result := db.Table(studentsTableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)
	if result == nil {
		log.Println(result)
	}
	if debugMode {
		log.Printf("tempStudent has value of: %s", tempStudent.AnonID)
	}
	return tempStudent
}

// desc: Get segments of student user
// status:
func GetUserSegments(r *http.Request) []StudentSegment {
	if db == nil {
		ConnectToDB()
	}

	var tempSegment []StudentSegment
	myAnonID := GetStudentUser(r.Header.Get("X-User")).AnonID
	tableToEdit := myAnonID + "_segments"
	result := db.Table(tableToEdit)
	paramTest := r.URL.Query()
	filter, params := paramTest["archived"]

	if !params || len(filter) == 0 {
		result = db.Table(tableToEdit).Where("archived = ?", false).Find(&tempSegment)
		if result != nil {
			log.Println(result.Error)
		}
	} else if paramTest.Get("archived") == "yes" {
		result = db.Table(tableToEdit).Find(&tempSegment)
		if result != nil {
			log.Println(result.Error)
		}
	} else if paramTest.Get("archived") == "only" {
		result = db.Table(tableToEdit).Where("archived = ?", true).Find(&tempSegment)
		if result != nil {
			log.Println(result)
		}
	} else {
		fmt.Println("Error: Invalid parameters.")
	}

	returnSegments := make([]StudentSegment, 0)
	result2, _ := result.Rows()

	var tempSegments2 StudentSegment
	for result2.Next() {

		if err3 := result.ScanRows(result2, &tempSegments2); err3 != nil {
			log.Println(err3)
		}
		returnSegments = append(returnSegments, tempSegments2)
	}

	return returnSegments
}

// Get Courses, default only active. with "archived=yes" all courses, "archived=only" to get only archived ones.
// Status: Works
func GetCourses(r *http.Request) []Course {
	if db == nil {
		ConnectToDB()
	}

	var tempCourses []Course
	var result *gorm.DB //db.Table(courseTableToEdit)
	paramTest := r.URL.Query()
	filter, params := paramTest["archived"]

	if !params || len(filter) == 0 {
		result = db.Table(courseTableToEdit).Where("archived = ?", false).Find(&tempCourses)
		if result != nil {
			log.Println(result.Error)
		}
	} else if paramTest.Get("archived") == "yes" {
		result = db.Table(courseTableToEdit).Find(&tempCourses)
		if result != nil {
			log.Println(result.Error)
		}
	} else if paramTest.Get("archived") == "only" {
		result = db.Table(courseTableToEdit).Where("archived = ?", true).Find(&tempCourses)
		if result != nil {
			log.Println(result)
		}
	} else {
		fmt.Println("Error: Invalid parameters.")
	}

	returnCourses := make([]Course, 0)
	result2, _ := result.Rows()

	var tempCourse2 Course
	for result2.Next() {

		if err3 := result.ScanRows(result2, &tempCourse2); err3 != nil {
			log.Println(err3)
		}
		returnCourses = append(returnCourses, tempCourse2)
	}

	return returnCourses
}

// FacultyUserSpesifics

// Desc: Get Faculty User
// Status: Works
func GetFacultyUser(FacultyID string) FacultyUser {
	if db == nil {
		ConnectToDB()
	}

	var tempFaculty FacultyUser

	result := db.Table(facultyTableToEdit).Where("faculty_id = ?", FacultyID).First(&tempFaculty)
	if result == nil {
		log.Println(result)
	}

	return tempFaculty
}

// desc: Get segments of faculty user, active and with parameters, archived=yes and archived=only
// status:
// comment: New version as added Archived bool, makes search simpler, doesn't require reading Courses tables.
func GetFacultyUserSegments(r *http.Request) []Segment {
	if db == nil {
		ConnectToDB()
	}
	var tempSegment []Segment
	var result *gorm.DB //db.Table(segmentTableToEdit)
	user := r.Header.Get("X-User")
	// Get teachers ID number
	teacher := GetFacultyUser(user)

	paramTest := r.URL.Query()
	filter, params := paramTest["archived"]
	if !params || len(filter) == 0 {
		result = db.Table(segmentTableToEdit).Where("teacher_id = ?", teacher.ID).Where("archived = ?", false).Find(&tempSegment)
		if result != nil {
			log.Println(result.Error)
		}
	} else if paramTest.Get("archived") == "yes" {
		result = db.Table(segmentTableToEdit).Where("teacher_id = ?", teacher.ID).Find(&tempSegment)
		if result != nil {
			log.Println(result.Error)
		}
	} else if paramTest.Get("archived") == "only" {
		result = db.Table(segmentTableToEdit).Where("teacher_id = ?", teacher.ID).Where("archived = ?", true).Find(&tempSegment)
		if result != nil {
			log.Println(result.Error)
		}
	} else {
		fmt.Println("Error: Invalid parameters.")
	}

	returnSegments := make([]Segment, 0)
	result2, _ := result.Rows()

	var tempSegments2 Segment
	for result2.Next() {
		//Read row to tempSegments2
		if err3 := result.ScanRows(result2, &tempSegments2); err3 != nil {
			log.Println(err3)
		}
		returnSegments = append(returnSegments, tempSegments2)
	}

	return returnSegments
}

// desc: Find course by courseCode
// Status: Unclear
func GetCourseTableByCode(courseCode string) Course {
	if db == nil {
		ConnectToDB()
	}
	var tempCourse Course
	db.Table(courseTableToEdit).Where("course_code = ?", courseCode).Find(&tempCourse).Row()
	return tempCourse
}

// desc: Find course by id
// Status: Unclear
func GetCourseTableById(id string) Course {
	if db == nil {
		ConnectToDB()
	}
	var tempCourse Course
	db.Table(courseTableToEdit).Where("id = ?", id).Find(&tempCourse).Row()
	return tempCourse
}

// desc: Find segment by id
// status:
func GetSegmentDataById(id uint) Segment {
	if db == nil {
		ConnectToDB()
	}
	var tempSegment Segment
	db.Table(segmentTableToEdit).Where("id = ?", id).Find(&tempSegment).Row()
	return tempSegment
}

func GetSegmentTableByCourseId(courseID uint) []Segment {
	if db == nil {
		ConnectToDB()
	}
	var tempSegment []Segment
	//tempSegment := make([]Segment, 0)
	result := db.Table(segmentTableToEdit).Where("course_id = ?", courseID).Find(&tempSegment)
	if result != nil {
		log.Println(result)
	}
	returnSegment := make([]Segment, 0)
	result2, _ := result.Rows()
	var tempCourse2 Segment
	for result2.Next() {
		if err3 := result.ScanRows(result2, &tempCourse2); err3 != nil {
			log.Println(err3)
		}
		returnSegment = append(returnSegment, tempCourse2)
	}
	return returnSegment
}

// desc: Find ALL categories belonging to segment
// comment: If using categories table for all segments
func GetCategoriesBySegmentId(segmentID uint) []SegmentCategory {
	if db == nil {
		ConnectToDB()
	}
	var tempSegment []SegmentCategory
	result := db.Table(categoriesTableToEdit).Where("segment_id = ?", segmentID).Find(&tempSegment)
	if result != nil {
		log.Println(result)
	}
	returnSegment := make([]SegmentCategory, 0)
	result2, _ := result.Rows()
	var tempSegment2 SegmentCategory
	for result2.Next() {
		if err3 := result.ScanRows(result2, &tempSegment2); err3 != nil {
			log.Println(err3)
		}
		returnSegment = append(returnSegment, tempSegment2)
	}
	return returnSegment
}

// Desc: GetAnonId with StudentID
// HOX! AnonID SHOULD NOT LEAVE OUTSIDE OF THE API
// Status: Done but REMOVED, smarter to use GetStudentUser
/*
func GetAnonId(StudentID string) string {
	if db == nil {
		ConnectToDB()
	}

	var tempStudent StudentUser

	result := db.Table(studentsTableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)
	if result == nil {
		log.Println(result)
	}
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	tempJSON := gjson.Get(s, "Value.AnonID")
	return tempJSON.String()
}
*/

// desc: Get name of the student with StudentID
// Status: Removed,not in use and GetStudentUser is smarter to use
/*
func GetStudentName(StudentID string) string {
	if db == nil {
		ConnectToDB()
	}

	var tempStudent StudentUser

	result := db.Table(courseTableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)
	if result == nil {
		log.Println(result)
	}
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	tempJSON := gjson.Get(s, "Value.StudentName")
	return tempJSON.String()
}
*/

// desc: Get segments of faculty user, active and with parameters, archived=yes and archived=only
// status: Works, but replaced by new version as Archived added possibility to archive invidual segments
// left if we happen to need to go back in this.
// HOX! Weird issues with filtering archived/non-archived results.
/*
func GetFacultyUserSegments(r *http.Request) []Segment {
	if db == nil {
		ConnectToDB()
	}
	var tempSegment []Segment

	user := r.Header.Get("X-User")
	// Get teachers ID number
	teacher := GetFacultyUser(user)
	result := db.Table(segmentTableToEdit).Where("teacher_id = ?", teacher.ID).Find(&tempSegment)
	paramTest := r.URL.Query()
	filter, params := paramTest["archived"]

	returnSegments := make([]Segment, 0)
	result2, _ := result.Rows()

	var tempSegments2 Segment
	for result2.Next() {

		//Read row to tempSegments2
		if err3 := result.ScanRows(result2, &tempSegments2); err3 != nil {
			log.Println(err3)
		}
		byID := tempSegments2.CourseID
		var tempCourse Course
		if err4 := db.Table(courseTableToEdit).First(&tempCourse, byID); err4 != nil {
			log.Println(err4)
		}
		log.Println("tempSegments2 value: ")
		log.Println(byID)
		if !params || len(filter) == 0 {
			if !tempCourse.Archived {
				returnSegments = append(returnSegments, tempSegments2)
			}
			if result != nil {
				log.Println(result.Error)
			}
		} else if paramTest.Get("archived") == "yes" {
			returnSegments = append(returnSegments, tempSegments2)
			if result != nil {
				log.Println(result.Error)
			}
		} else if paramTest.Get("archived") == "only" {
			if tempCourse.Archived {
				returnSegments = append(returnSegments, tempSegments2)
			}
			if result != nil {
				log.Println(result)
			}
		} else {
			fmt.Println("Error: Invalid parameters.")
		}
	}
	return returnSegments
}*/
