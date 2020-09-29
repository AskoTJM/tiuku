package tiuku

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

	if h == "populate" {
		database.PopulateSchool()
		return "Populated School"
	}

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
		database.GetCourses(r)
	}
	if h == "studentsegment" {
		user := r.Header.Get("X-User")
		return database.CreateStudentSegmentTable(user)

	}
	if h == "createcourse" {
		return database.CreateCourse(r)
	}
	if h == "getstudentdata" {
		user := r.Header.Get("X-User")
		result := database.GetStudent(user)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		/*
			fmt.Fprintf(w, "%s", s)
			return "yep"
		*/

		//tempJSON := gjson.Get(s, "Value.AnonID")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		return s //tempJSON.String()
	}
	if h == "findasso" {
		database.CheckAssociation(w, r)

	}
	if h == "maincategories" {

	}

	return "nothing"
}
