/*
// get_api_student_v1.go
// Description: GET request functions for Student users
// CRUD: Read, Collection: 200(OK) Item: 200(OK), 404 (Not Found)
*/
package students

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/gorilla/mux"
)

// Get list of courses available
// status:
func GetCourses(w http.ResponseWriter, r *http.Request) {

	result := database.GetCourses(r)
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
	result := database.GetCourseTableById(courseCode)
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
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	segId := vars["segment"]
	// Get course information
	//result := database.GetCourseTableById(segId)
	// Get segment data
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

// desc:Get 'Your' Sessions for {segment}
// status:
func GetSegmentsSegmentSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Get value of {setting} for {segment}
// Can't remember what this was about
func GetSegmentsSegmentSettingsSetting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

// What is this
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
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Problems with the server, please try again later.")
	} else {
		result := database.GetUserSegments(r)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", s)

	}
}

// Get Session data for {segment}
// Not sure if this is needed, probably already have different endpoint to handle this
func GetUserSegmentsResourceID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
