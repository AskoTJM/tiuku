package database

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tidwall/gjson"
)

// Desc: GetAnonId with StudentID
// HOX! AnonID SHOULD NOT LEAVE OUTSIDE OF THE API
// Status: Done
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

// desc: Get name of the student with StudentID
// Status: Should work, c+p
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
	myAnonID := GetAnonId(r.Header.Get("X-User"))
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
	result := db.Table(courseTableToEdit)
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
// status: works
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
}
