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

	result := db.Table(courseTableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)
	if result == nil {
		log.Panic(result)
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
		log.Panic(result)
	}
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	tempJSON := gjson.Get(s, "Value.StudentName")
	return tempJSON.String()
}

// Desc: Get Students data
// Status: Works, but needs more. Return value and obfuscing of AnonID if used outside
func GetStudent(StudentID string) StudentUser {
	if db == nil {
		ConnectToDB()
	}

	var tempStudent StudentUser

	//result :=
	db.Table(courseTableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)

	/*
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
	*/
	//if result.Error != nil {
	//	log.Panic(result)
	//}

	return tempStudent
}

// Get Courses, default only active. with "archived=yes" all courses
// , "archived=only" to get only archived ones.
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

	/*
		// Test what we got from db
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		log.Println(s)
	*/
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

// desc: Find course by courseCode
// Status: Unclear
func FindCourseTableByCode(courseCode string) Course {
	if db == nil {
		ConnectToDB()
	}

	var tempCourse Course

	db.Table(courseTableToEdit).Where("course_code = ?", courseCode).Find(&tempCourse).Row()

	return tempCourse
}

// desc: Find course by id
// Status: Unclear
func FindCourseTableById(id string) Course {
	if db == nil {
		ConnectToDB()
	}

	var tempCourse Course

	db.Table(courseTableToEdit).Where("id = ?", id).Find(&tempCourse).Row()

	return tempCourse
}
