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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)
}

// Get {course} information
// W0rks.
func GetCoursesCourse(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	courseCode := vars["course"]
	result := database.GetCourseTableById(scripts.StringToUint(courseCode))

	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)
}

// Get list of segments for {course}
// W0rks
func GetCoursesCourseSegments(w http.ResponseWriter, r *http.Request) {

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

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)
}

// Get data of the {segment} in the {course}
// W0rks
func GetCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//courseCode := vars["course"]
	segCode := vars["segment"]
	// Get course information
	//courseRes := database.FindCourseTableById(courseCode)
	// Get segment data
	segRes := database.GetSegmentDataById(scripts.StringToUint(segCode))
	//Transform results to json
	anon, _ := json.Marshal(segRes)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)
}

// Get Students for the {segment}
// W0rks
func GetCoursesCourseSegmentsSegmentStudents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	segCode := vars["segment"]

	studentResult := database.GetStudentsJoinedOnSegment(scripts.StringToUint(segCode))
	anon, _ := json.Marshal(studentResult)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)

}

// Get {category} for the {Segment}
// W0rks
func GetCoursesCourseSegmentsSegmentCategoriesCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	catCode := vars["category"]
	catRes := database.GetCategoryById(scripts.StringToUint(catCode))

	anon, _ := json.Marshal(catRes)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)

}

// Get specific settings of {category}
// Not sure if this is needed. Better to serve all settings at once.
func GetCoursesCourseSegmentsSegmentCategoriesCategorySettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Change(GET) {setting} in {category}
// status: Unnecessary? Better way to do this by sending all the new settings. Or least should be PUT/PATCH
// T0D0 : Remove or repurpose
func GetCoursesCourseSegmentsSegmentCategoriesCategorySettingsSetting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Get Categories for {segment} of the {course}
// status: works
// T0D0 : Think about this.
func GetCoursesCourseSegmentsSegmentSettings(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	segId := vars["segment"]

	res := scripts.StringToUint(segId)
	result2 := database.GetCategoriesBySegmentId(res, true, true)

	anon, _ := json.Marshal(result2)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)

}

// Get segments table for Faculty User
// W0rks , I think
func GetUserSegments(w http.ResponseWriter, r *http.Request) {

	user := r.Header.Get("X-User")
	returnNum := database.CheckIfFacultyUserExists(user)
	var choice string
	if returnNum == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "Incorrect request")
	} else if returnNum > 1 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Problems with the server, please try again later.")
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
			fmt.Println("Error: Invalid parameters.")
		}
		result := database.GetFacultyUserSegments(user, choice)

		anon, _ := json.Marshal(result)

		n := len(anon)
		s := string(anon[:n])
		log.Println(s)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", s)
	}

}
