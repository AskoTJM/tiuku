/*
// get_api_faculty_v1.go
// Description: GET request functions for Faculty users
//
*/
package faculty

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/gorilla/mux"
)

// List of active courses
// W0rks
func GetCourses(w http.ResponseWriter, r *http.Request) {

	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		paramTest := r.URL.Query()
		filter, params := paramTest["archived"]
		var choice string
		if !params || len(filter) == 0 {
			choice = "no"
		} else if paramTest.Get("archived") == "yes" {
			choice = "yes"
		} else if paramTest.Get("archived") == "only" {
			choice = "only"
		} else {
			log.Println("Error: Invalid parameters.")
		}

		result := database.GetCourses(choice)

		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)
}

// Get {course} information
// W0rks.
func GetCoursesCourse(w http.ResponseWriter, r *http.Request) {

	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {

		vars := mux.Vars(r)
		courseCode := vars["course"]
		result := database.GetCourseTableById(scripts.StringToUint(courseCode))

		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)
}

// Get list of segments for {course}
// W0rks
func GetCoursesCourseSegments(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		vars := mux.Vars(r)
		courseCode := vars["course"]
		// Get course information
		result := database.GetCourseTableById(scripts.StringToUint(courseCode))
		// Get segment data
		result2 := database.GetSegmentTableByCourseId(result.ID)
		//Transform results to json
		anon, _ := json.Marshal(result2)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}

	fmt.Fprintf(w, "%s", response)
}

// Get data of the {segment} in the {course}
// W0rks
func GetCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {

	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)

	} else {

		vars := mux.Vars(r)
		segCode := vars["segment"]
		segRes := database.GetSegmentDataById(scripts.StringToUint(segCode))
		anon, _ := json.Marshal(segRes)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)
}

// Get Students for the {segment}
// W0rks
func GetCoursesCourseSegmentsSegmentStudents(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)

	} else {

		vars := mux.Vars(r)
		segCode := vars["segment"]

		studentResult := database.GetStudentsJoinedOnSegment(scripts.StringToUint(segCode))
		anon, _ := json.Marshal(studentResult)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get Sessions for the {segment}
// W0rks
func GetCoursesCourseSegmentsSegmentSessions(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)

	} else {
		vars := mux.Vars(r)
		segCode := vars["segment"]

		studentResult := database.GetAllSessionsForSegment(scripts.StringToUint(segCode))
		anon, _ := json.Marshal(studentResult)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get {category} for the {Segment}
// W0rks, maybe needs filtering so not to work for categories not beloning to the segment.
func GetCoursesCourseSegmentsSegmentCategoriesCategory(w http.ResponseWriter, r *http.Request) {

	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {

		vars := mux.Vars(r)
		catCode := vars["category"]
		catRes := database.GetCategoryById(scripts.StringToUint(catCode))

		anon, _ := json.Marshal(catRes)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get Categories for {segment} of the {course}
// W0rks
// But, Think about this. Should rename to SegmentCategories
func GetCoursesCourseSegmentsSegmentSettings(w http.ResponseWriter, r *http.Request) {

	var response string

	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		vars := mux.Vars(r)
		segId := vars["segment"]
		res := scripts.StringToUint(segId)
		result2 := database.GetCategoriesBySegmentId(res, true, true)
		anon, _ := json.Marshal(result2)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get segments table for Faculty User
// W0rks
func GetUserSegments(w http.ResponseWriter, r *http.Request) {

	var response string
	var choice string

	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {

		paramTest := r.URL.Query()
		filter, params := paramTest["archived"]
		if !params || len(filter) == 0 {
			choice = "no"
		} else if paramTest.Get("archived") == "yes" {
			choice = "yes"

		} else if paramTest.Get("archived") == "only" {
			choice = "only"

		} else {
			choice = "error"
			w.WriteHeader(http.StatusBadRequest)
		}

		if choice != "error" {
			result := database.GetFacultyUserSegments(user, choice)
			anon, _ := json.Marshal(result)
			n := len(anon)
			s := string(anon[:n])
			response = s
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
		}
	}
	fmt.Fprintf(w, "%s", response)

}

// Get Students.
// W0rks
func GetStudents(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		result := database.GetStudents(0)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get Student.
// W0rks
func GetStudentsStudent(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {

		vars := mux.Vars(r)
		stuId := vars["student"]
		result := database.GetStudents(scripts.StringToUint(stuId))
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get Faculty users.
// W0rks
func GetFaculty(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		result := database.GetFaculty(0)
		//log.Println(result)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get Faculty by id.
// W0rks
func GetFacultyFaculty(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {

		vars := mux.Vars(r)
		stuId := vars["faculty"]
		result := database.GetFaculty(scripts.StringToUint(stuId))
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get school,
// W0rks
func GetSchools(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		result := database.GetSchool(0)
		//log.Println(result)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get campus,
// W0rks
func GetCampuses(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		result := database.GetCampus(0)
		//log.Println(result)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get apartments
// W0rks
func GetApartments(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		result := database.GetApartment(0)
		//log.Println(result)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get degrees
// W0rks
func GetDegrees(w http.ResponseWriter, r *http.Request) {
	var response string
	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else {
		result := database.GetDegree(0)
		//log.Println(result)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}
