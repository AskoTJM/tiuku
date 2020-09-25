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
	tableToEdit := schoolShortName + "_StudentUsers"
	var tempStudent StudentUser

	result := db.Table(tableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)
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
	tableToEdit := schoolShortName + "_StudentUsers"
	var tempStudent StudentUser

	result := db.Table(tableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)
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
	tableToEdit := schoolShortName + "_StudentUsers"
	var tempStudent StudentUser

	//result :=
	db.Table(tableToEdit).Where("student_id = ?", StudentID).First(&tempStudent)

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
// Status: Problems getting []struct to fill with rows
func GetCourses(r *http.Request) []Course {
	if db == nil {
		ConnectToDB()
	}

	tableToEdit := schoolShortName + "_Courses"
	var tempCourses []Course
	//var returnCourses []Course
	result := db.Table(tableToEdit)
	//result2, _ := db.Table(tableToEdit).Rows()
	paramTest := r.URL.Query()
	filter, params := paramTest["archived"]
	// Ye olde one, that worked with JSON

	if !params || len(filter) == 0 {
		result = db.Table(tableToEdit).Where("archived = ?", false).Find(&tempCourses)
		if result != nil {
			log.Println(result.Error)
		}
	} else if paramTest.Get("archived") == "yes" {
		result = db.Table(tableToEdit).Find(&tempCourses)
		if result != nil {
			log.Println(result.Error)
		}
	} else if paramTest.Get("archived") == "only" {
		result = db.Table(tableToEdit).Where("archived = ?", true).Find(&tempCourses)
		if result != nil {
			log.Println(result)
		}
	} else {
		fmt.Println("Error: Invalid parameters.")
	}
	/*
		// New one trying to get to blasted rows.Next() to work right
		if !params || len(filter) == 0 {
			result2, _ = db.Table(tableToEdit).Where("archived = ?", false).Find(&tempCourses).Rows()
			if result2 != nil {
				log.Println(result2)
			}
		} else if paramTest.Get("archived") == "yes" {
			result2, _ = db.Table(tableToEdit).Find(&tempCourses).Rows()
			if result2 != nil {
				log.Println(result2)
			}
		} else if paramTest.Get("archived") == "only" {
			result2, _ = db.Table(tableToEdit).Where("archived = ?", true).Find(&tempCourses).Rows()
			if result2 != nil {
				log.Println(result2)
			}
		} else {
			fmt.Println("Error: Invalid parameters.")
		}
	*/
	// Test what we got from db
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])
	log.Println(s)
	/*
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])

		tempJSON := gjson.Get(s, "Value")
		//return tempJSON.String()
		w.Write(tempJSON)
		fmt.Fprintf(w, "%s", tempJSON)
	*/
	var i int64
	var c int64 = 0
	//defer result.Close()
	returnCourses := make([]Course, 0)
	result2, _ := result.Rows()
	//log.Println(result2)
	//i = result.RowsAffected
	log.Println("Value of i is %i", i)
	var tempCourse2 Course
	for result2.Next() {
		c++
		log.Println("Value of c is %i", c)
		//result2.Next() c < i { //

		if err3 := result.ScanRows(result2, &tempCourse2); err3 != nil {
			log.Println(err3)
		}

		/*
			if err3 := result2.Scan(&tempCourse2.ID, &tempCourse2.ResourceID, &tempCourse2.CourseCode,
				&tempCourse2.CourseName, &tempCourse2.CourseStartDate, &tempCourse2.CourseEndDate, &tempCourse2.Archived); err3 != nil {
				log.Fatal(err3)
			}
		*/
		//log.Println(result2)
		log.Println(tempCourse2)
		/*
			if err3 := scripts.StructScan(result2, &tempCourse2); err3 != nil {
				log.Fatal(err3)
			}
		*/
		returnCourses = append(returnCourses, tempCourse2)
		//log.Println(returnCourses2)
		log.Println(returnCourses)
		//returnCourses = returnCourses2

		//returnCourses = appe
	}

	return returnCourses
}
