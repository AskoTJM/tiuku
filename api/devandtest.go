package tiuku

/*
// devandtest.go
// Description: Code for testing features before proper implementing
// also for running population scripts.
*/

import (
	"encoding/json"
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
	/*
		if h == "populate" {
			//database.PopulateSchool()
			//return "Populated School"
		}
	*/
	if h == "populatecourses" {
		num := r.Header.Get("X-Number")
		i, _ := strconv.Atoi(num)
		database.PopulateCourses(i)
		return "Populated courses"
	}

	if h == "populatestudents" {
		num := r.Header.Get("X-Number")
		i, _ := strconv.Atoi(num)
		database.PopulateStudents(i)
		return "Populated students"
	}

	if h == "populatefaculty" {
		num := r.Header.Get("X-Number")
		i, _ := strconv.Atoi(num)
		database.PopulateFaculty(i)
		return "Populated faculty"
	}
	if h == "Hello" {
		return "Hello"
	}
	if h == "anonId" {
		user := r.Header.Get("X-User")
		return database.GetStudentUser(user).AnonID
	}
	if h == "countusers" {
		user := r.Header.Get("X-User")
		returnNum := database.CheckIfUserExists(user)
		s := strconv.Itoa(int(returnNum))
		return s
		//log.Println(returnNum)
	}
	if h == "courses" {
		database.GetCourses(r)
	}
	if h == "studentsegment" {
		user := r.Header.Get("X-User")
		studentTemp := database.GetStudentUser(user)
		return database.CreateStudentSegmentTable(studentTemp)

	}
	if h == "createcourse" {
		return database.CreateCourse(w, r)
	}
	if h == "getstudentdata" {
		user := r.Header.Get("X-User")
		result := database.GetStudentUser(user)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return s
	}
	if h == "findasso" {
		database.CheckAssociation(w, r)

	}
	if h == "populatesegments" {
		database.AutoCreateSegments()
		//database.AutoCreateStudentUserTables()
	}
	if h == "populatesegments2" {
		//database.AutoCreateSegments()
		database.AutoCreateStudentUserTables()
		// Do faculty actually need their own lists as we can search segments table and filter that.
		//database.AutoCreateFacultyUserTables()
	}
	if h == "populatesegments3" {
		database.PopulateCategories()
	}
	if h == "archivecourse" {
		courseNum := r.Header.Get("X-Course")
		//log.Println(courseNum)
		tempCourse := database.GetCourseTableById(courseNum)
		//log.Println(tempCourse)
		database.ArchiveCourse(tempCourse, true)
	}

	return "nothing"
}
