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
	"github.com/AskoTJM/tiuku/api/stats"
	"github.com/gorilla/mux"
)

// Get Sessions for Student User
// W0rks but a bit ugly.
func GetSessions(w http.ResponseWriter, r *http.Request) {

	user := r.Header.Get("X-User")

	var resCode int
	var resString string
	var response string

	resString, resCode = database.CheckIfUserExists(user)
	if resCode != http.StatusOK {
		w.WriteHeader(resCode)
		response = resString
	} else {
		var result []database.StudentSegmentSession
		paramTest := r.URL.Query()
		filter, params := paramTest["archived"]

		if !params || len(filter) == 0 {
			result = database.GetStudentsSessionsForSegment(user, 0)
			anon, _ := json.Marshal(result)
			n := len(anon)
			s := string(anon[:n])
			response = s
			w.WriteHeader(http.StatusOK)

		} else if paramTest.Get("archived") == "yes" {
			result = database.GetStudentsSessionsForSegment(user, 0)
			result2 := database.GetStudentsArchivedSessions(user, 0)
			result = append(result, result2...)
			anon, _ := json.Marshal(result)
			n := len(anon)
			s := string(anon[:n])
			response = s
			w.WriteHeader(http.StatusOK)
		} else if paramTest.Get("archived") == "only" {
			result = database.GetStudentsArchivedSessions(user, 0)
			anon, _ := json.Marshal(result)
			n := len(anon)
			s := string(anon[:n])
			response = s
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
		} else {
			log.Println("Error: Invalid parameters. <get_api_student_v1.go->GetSessions")
			response = "Error: Invalid parameters."
			w.WriteHeader(http.StatusBadRequest)
		}
		/*
			anon, _ := json.Marshal(result)
			n := len(anon)
			s := string(anon[:n])
			response = s
			w.WriteHeader(http.StatusOK)
		*/
	}

	fmt.Fprintf(w, "%s", response)

}

// Get Last Sessions for Student User
// W0rks
func GetSessionsLast(w http.ResponseWriter, r *http.Request) {

	user := r.Header.Get("X-User")
	var resCode int
	var resString string
	var response string

	resString, resCode = database.CheckIfUserExists(user)
	if resCode != http.StatusOK {
		w.WriteHeader(resCode)
		response = resString
	} else {
		//var result database.StudentSegmentSession
		result := database.GetLastSession(user)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		response = s
	}

	fmt.Fprintf(w, "%s", response)

}

// Get school...
// W0rks
func GetSchools(w http.ResponseWriter, r *http.Request) {

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
		result := database.GetSchool(0)
		//log.Println(result)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get campus...
// W0rks
func GetCampuses(w http.ResponseWriter, r *http.Request) {
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
		result := database.GetCampus(0)
		//log.Println(result)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get apartments...
// W0rks
func GetApartments(w http.ResponseWriter, r *http.Request) {
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
		result := database.GetApartment(0)
		//log.Println(result)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get degrees...
// W0rks
func GetDegrees(w http.ResponseWriter, r *http.Request) {
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
		result := database.GetDegree(0)
		//log.Println(result)
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get list of courses available
// W0rks
func GetCourses(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)
}

// Get {course} information
// W0rks
func GetCoursesCourse(w http.ResponseWriter, r *http.Request) {
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
		vars := mux.Vars(r)
		courseCode := vars["course"]
		result := database.GetCourseTableById(scripts.StringToUint(courseCode))

		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)
}

// Get list of segments on the {course}
// W0rks
func GetCoursesCourseSegments(w http.ResponseWriter, r *http.Request) {
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
		vars := mux.Vars(r)
		courseCode := vars["course"]
		result2 := database.GetSegmentTableByCourseId(scripts.StringToUint(courseCode))
		// Transform results to json
		anon, _ := json.Marshal(result2)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)
}

// Get  data for specific Segment
// W0rks
// comment: Maybe add information if enrolled to it already?
func GetCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {

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
		vars := mux.Vars(r)
		segCode := vars["segment"]
		segRes := database.GetSegmentDataById(scripts.StringToUint(segCode))
		//Transform results to json
		anon, _ := json.Marshal(segRes)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)
}

// Get categories for {segment}
// W0rks
func GetCoursesCourseSegmentsSegmentCategories(w http.ResponseWriter, r *http.Request) {
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

		vars := mux.Vars(r)
		segId := vars["segment"]
		res := scripts.StringToUint(segId)
		// Get categories for segment, filter out InActive ones.
		result2 := database.GetCategoriesBySegmentId(res, true, false)
		// Transform results to json
		anon, _ := json.Marshal(result2)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// desc:Get sessions for {segment}
// W0rks for active, need to work for archived
func GetSegmentsSegmentSessions(w http.ResponseWriter, r *http.Request) {

	user := r.Header.Get("X-User")
	//var resCode int
	//var resString string
	var response string

	resString, resCode := database.CheckIfUserExists(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if resCode != http.StatusOK {
		w.WriteHeader(resCode)
		response = resString
	} else {
		var result []database.StudentSegmentSession
		vars := mux.Vars(r)
		segId := vars["segment"]
		paramTest := r.URL.Query()
		filter, params := paramTest["stats"]
		// Check if the requested segment is in active or achived
		testString, testCode := database.CheckStudentsSegmentStatus(user, scripts.StringToUint(segId))
		log.Printf("TestString is: %s and testCode is: %d", testString, testCode)
		// End Check
		if testString == "active" {
			result = database.GetStudentsSessionsForSegment(user, scripts.StringToUint(segId))
		} else if testString == "archived" {
			result = database.GetStudentsArchivedSessions(user, scripts.StringToUint(segId))
		}
		if !params || len(filter) == 0 {
			anon, _ := json.Marshal(result)
			n := len(anon)
			s := string(anon[:n])
			response = s
			w.WriteHeader(http.StatusOK)
		} else if paramTest.Get("stats") == "overall" {
			responseTime, _ := stats.CalculateOverAllTime(result)
			response = responseTime.String()
		} else if paramTest.Get("stats") == "week" {
			response = "week"
		} else {
			log.Println("Error: Invalid parameters.")
		}
		/*
			result := database.GetStudentsSessionsForSegment(user, scripts.StringToUint(segId))
			anon, _ := json.Marshal(result)
			n := len(anon)
			s := string(anon[:n])
			response = s
		*/
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// desc: Get particular {session} for {segment}
// W0rks
func GetSegmentsSegmentSessionsSession(w http.ResponseWriter, r *http.Request) {

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
		vars := mux.Vars(r)
		sesId := vars["session"]
		result := database.GetSession(user, scripts.StringToUint(sesId))
		anon, _ := json.Marshal(result)
		n := len(anon)
		s := string(anon[:n])
		response = s
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)
}

// Get list of active segments for student user
// W0rks
func GetUserSegments(w http.ResponseWriter, r *http.Request) {

	user := r.Header.Get("X-User")
	var resCode int
	var resString string
	var response string

	resString, resCode = database.CheckIfUserExists(user)
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
	}
	fmt.Fprintf(w, "%s", response)

}

// Get Session data for {segment}, filtered with Resource ID to filter out deleted data.
// Not sure if this is needed, probably already have different endpoint to handle this
// T0D0
func GetUserSegmentsResourceID(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusNotImplemented)
	}
	fmt.Fprintf(w, "%s", response)
}

// Get Faculty users.
// W0rks
func GetFaculty(w http.ResponseWriter, r *http.Request) {

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
// T0D0
func GetFacultyFaculty(w http.ResponseWriter, r *http.Request) {
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
