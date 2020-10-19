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

// Get Sessions for Student User
// W0rks
func GetSessions(w http.ResponseWriter, r *http.Request) {

	user := r.Header.Get("X-User")
	var result []database.StudentSegmentSession
	var resCode int
	var resString string
	var response string

	resString, resCode = database.CheckIfUserExists(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if resCode != http.StatusOK {
		w.WriteHeader(resCode)
		response = resString
	} else {

		paramTest := r.URL.Query()
		filter, params := paramTest["archived"]

		if !params || len(filter) == 0 {
			result = database.GetStudentsSessionsForSegment(user, 0)

		} else if paramTest.Get("archived") == "yes" {
			result = database.GetStudentsSessionsForSegment(user, 0)
			result2 := database.GetStudentsArchivedSessions(user, 0)
			result = append(result, result2...)

		} else if paramTest.Get("archived") == "only" {
			result = database.GetStudentsArchivedSessions(user, 0)

		} else {
			log.Println("Error: Invalid parameters.")
		}

		//result = database.GetStudentsSessionsForSegment(user, 0) //, choice)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])

		w.WriteHeader(http.StatusOK)
		response = s
	}

	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", response)

}

// Get Last Sessions for Student User
// W0rks
func GetSessionsLast(w http.ResponseWriter, r *http.Request) {

	user := r.Header.Get("X-User")
	var result database.StudentSegmentSession
	var resCode int
	var resString string
	var response string

	resString, resCode = database.CheckIfUserExists(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if resCode != http.StatusOK {
		w.WriteHeader(resCode)
		response = resString
	} else {

		result = database.GetLastSession(user)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])

		w.WriteHeader(http.StatusOK)
		response = s
	}

	fmt.Fprintf(w, "%s", response)

}

// Get school...
// W0rks
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

// Get campus...
// W0rks
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

// Get apartments...
// W0rks
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

// Get degrees...
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

	var resCode int
	var resString string
	var response string

	resString, resCode = database.CheckIfUserExists(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if resCode != http.StatusOK {
		w.WriteHeader(resCode)
		response = resString
	} else {
		result := database.GetStudentsSessionsForSegment(user, scripts.StringToUint(segId))
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		w.WriteHeader(http.StatusOK)
		response = s
	}
	fmt.Fprintf(w, "%s", response)

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
// T0D0
func GetSegmentsSegmentSettingsSetting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

// Get settings for segment, ? Wut?
// T0D0
func GetUserSegmentsSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Get list of active segments for student user
// W0rks
func GetUserSegments(w http.ResponseWriter, r *http.Request) {

	user := r.Header.Get("X-User")
	var resCode int
	var resString string
	var response string

	resString, resCode = database.CheckIfUserExists(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if resCode != http.StatusOK {
		w.WriteHeader(resCode)
		response = resString
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
		usedStudent := database.GetStudentUserWithStudentID(user)
		result := database.GetUserSegments(usedStudent, choice)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get Session data for {segment}, filtered with Resource ID to filter out deleted data.
// Not sure if this is needed, probably already have different endpoint to handle this
// T0D0
func GetUserSegmentsResourceID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
