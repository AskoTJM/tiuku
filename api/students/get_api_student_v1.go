/*
// get_api_student_v1.go
// Description: GET request functions for Student users
// CRUD: Read, Collection: 200(OK) Item: 200(OK), 404 (Not Found)
*/
package students

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/gorilla/mux"
)

// Get list of courses available
// status:
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

// Get list of segments on the {course}
// status:
func GetCoursesCourseSegments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseCode := vars["course"]
	// Get course information
	result := database.GetCourseTableById(scripts.StringToUint(courseCode))
	// Get segment data
	result2 := database.GetSegmentTableByCourseId(result.ID)
	// Transform results to json
	anon, _ := json.Marshal(result2)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)
}

// Get your data for specific Segment
// status:
// comment: Maybe add information if enrolled to it already?
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

// Get categories used for {segment}
// status:
func GetCoursesCourseSegmentsSegmentCategories(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	segId := vars["segment"]

	res := scripts.StringToUint(segId)
	// Get categories for segment, filter out InActive ones.
	result2 := database.GetCategoriesBySegmentId(res, true, false)
	// Transform results to json
	anon, _ := json.Marshal(result2)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)

}

// desc:Get sessions for {segment}
// status: ToDo
func GetSegmentsSegmentSessions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	segId := vars["segment"]
	user := r.Header.Get("X-User")
	returnNum := database.CheckIfUserExists(user)
	if returnNum == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "Incorrect request")
	} else if returnNum > 1 {
		log.Printf("Error: Found %d with userId %s", returnNum, user)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Problems with the server, please try again later.")
	} else {
		result := database.GetAllSessionsForSegment(user, scripts.StringToUint(segId))
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", s)
	}
}

// desc:Get particular {session} for {segment}
// status:
func GetSegmentsSegmentSessionsSession(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//segId := vars["segment"]

	//anon, _ := json.Marshal(result)
	//n := len(anon)
	//s := string(anon[:n])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", "SessionsSession")
}

// Get value of {setting} for {segment}
// 'GET particular Setting of the Segment'
// Can't remember what this was about, or is it actually needed
func GetSegmentsSegmentSettingsSetting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

// Get settings for segment, ? Same as getting categories?
func GetUserSegmentsSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Get list of active segments for student user
// status:
func GetUserSegments(w http.ResponseWriter, r *http.Request) {

	user := r.Header.Get("X-User")
	returnNum := database.CheckIfUserExists(user)
	if returnNum == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "Incorrect request")
	} else if returnNum > 1 {
		log.Printf("Error: Found %d with userId %s", returnNum, user)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Problems with the server, please try again later.")
	} else {
		var choice string
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
		usedStudent := database.GetStudentUser(user)
		result := database.GetUserSegments(usedStudent, choice)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", s)

	}
}

// Get Session data for {segment}, filtered with Resource ID to filter out deleted data.
// Not sure if this is needed, probably already have different endpoint to handle this
func GetUserSegmentsResourceID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
