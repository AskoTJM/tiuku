package tiuku

import (
	"net/http"
	"strconv"

	database "github.com/AskoTJM/tiuku/api/database"
)

func HeaderTests(w http.ResponseWriter, r *http.Request) string {
	h := r.Header.Get("X-Init")
	if h == "db" {
		database.InitDB()
		return "InitDB"
	}
	if h == "populate" {
		database.PopulateSchool()
		database.PopulateStudents()
		database.PopulateCourses()
		return "Populating"
	}
	if h == "Hello" {
		return "Hello"
	}
	if h == "anonId" {
		user := r.Header.Get("X-User")
		return database.GetAnonId(user)
	}
	if h == "countusers" {
		user := r.Header.Get("X-User")
		returnNum := database.CheckIfUserExists(user)
		s := strconv.Itoa(int(returnNum))
		return s
		//log.Println(returnNum)
	}
	if h == "courses" {
		database.GetCourses(w, r)
	}
	if h == "studentsegment" {
		user := r.Header.Get("X-User")
		return database.CreateStudentSegmentTable(user)
	}
	if h == "createcourse" {
		return database.CreateCourse(r)
	}

	return "nothing"
}
