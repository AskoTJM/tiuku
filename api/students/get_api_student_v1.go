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

// Get school, campus, apartments and degrees...
func GetSchools(w http.ResponseWriter, r *http.Request) {

	result := database.GetSchool(0)
	//log.Println(result)
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)

}

// Get school, campus, apartments and degrees...
func GetCampuses(w http.ResponseWriter, r *http.Request) {

	result := database.GetCampus(0)
	//log.Println(result)
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)

}

// Get school, campus, apartments and degrees...
func GetApartments(w http.ResponseWriter, r *http.Request) {

	result := database.GetApartment(0)
	//log.Println(result)
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)

}

// Get school, campus, apartments and degrees...
func GetDegrees(w http.ResponseWriter, r *http.Request) {

	result := database.GetDegree(0)
	//log.Println(result)
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)

}

// Get list of courses available
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
// W0rks
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

// Get list of segments on the {course}
// W0rks
func GetCoursesCourseSegments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseCode := vars["course"]
	result2 := database.GetSegmentTableByCourseId(scripts.StringToUint(courseCode))
	// Transform results to json
	anon, _ := json.Marshal(result2)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)
}

// Get  data for specific Segment
// W0rks
// comment: Maybe add information if enrolled to it already?
func GetCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	segCode := vars["segment"]

	segRes := database.GetSegmentDataById(scripts.StringToUint(segCode))
	//Transform results to json
	anon, _ := json.Marshal(segRes)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)
}

// Get categories for {segment}
// W0rks
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
// W0rks
func GetSegmentsSegmentSessions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	segId := vars["segment"]
	user := r.Header.Get("X-User")

	// Migrate UserChecking to next func ?
	// would be more consistent with dividing of work
	returnNum := database.CheckIfUserExists(user)
	if returnNum == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "Incorrect request")
	} else if returnNum > 1 {
		log.Printf("Error: Found %d with userId %s", returnNum, user)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Problems with user identification, please try again later.")
	} else {
		result := database.GetStudentsSessionsForSegment(user, scripts.StringToUint(segId))
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", s)
	}
}

// desc: Get particular {session} for {segment}
// W0rks
func GetSegmentsSegmentSessionsSession(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	//segId := vars["segment"]
	sesId := vars["session"]
	user := r.Header.Get("X-User")
	result := database.GetSession(user, scripts.StringToUint(sesId))
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)
}

// Get value of {setting} for {segment}
// 'GET particular Setting of the Segment'
// Can't remember what this was about, or is it actually needed
func GetSegmentsSegmentSettingsSetting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

// Get settings for segment, ? Wut?
func GetUserSegmentsSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Get list of active segments for student user
// W0rks
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
