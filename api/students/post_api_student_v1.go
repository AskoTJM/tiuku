package students

/*
// post_api_student_v1.go
// Description: POST request functions for Student users
// CRUD: Create, should return 201 on collection with Id, on Specific 404 (Not Found) or 409 (Conflict)
*/

import (
	"fmt"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/gorilla/mux"
)

// Join {segment} of the {course}
// status:
func PostCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	segCode := vars["segment"]
	resSeg := database.GetSegmentDataById(scripts.StringToUint(segCode))

	user := r.Header.Get("X-User")
	returnNum := database.CheckIfUserExists(user)
	if returnNum == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "Incorrect request")
	} else if returnNum > 1 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Problems with the server, please try again later.")
	} else {
		studentToJoin := database.GetStudentUser(user)
		res := database.AddStudentToSegment(studentToJoin, resSeg)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", res)
	}

}

// Add or start new session to {segment} table
// status:
func PostSegmentsSegmentSessions(w http.ResponseWriter, r *http.Request) {

	// := mux.Vars(r)
	//segCode := vars["segment"]
	//resSeg := database.GetSegmentDataById(scripts.StringToUint(segCode))

	user := r.Header.Get("X-User")
	returnNum := database.CheckIfUserExists(user)
	if returnNum == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "Incorrect request")
	} else if returnNum > 1 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Problems with the server, please try again later.")
	} else {
		// Check if content is valid
		if r.Header.Get("Content-Type") == "" {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusNoContent)
			response := "Problems creating new Course, no body in request information. <database/create_database->CreateCourse> Error: No body information available."
			fmt.Fprintf(w, "%s", response)
		} else {
			//rbody, _ := header.ParseValueAndParams(r.Header, "Content-Type")
			rbody := r.Header.Get("Content-Type")
			// Check if content type is correct one.
			if rbody != "application/json" {
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusNotAcceptable)
				response := "Error: Content-Type is not application/json."
				fmt.Fprintf(w, "%s", response)
			}

		}

		studentNow := database.GetStudentUser(user)
		var session database.StudentSegmentSession

		res := database.AddSessionToSegment(studentNow, session)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", res)
	}
}

// Add Segment to Students CourseList
// comment: Not sure this is needed. Segment should be automatically added on Students segments anyway.
func PostUserSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
