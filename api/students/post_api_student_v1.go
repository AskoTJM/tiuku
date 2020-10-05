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
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
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
func PostSegmentsSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Add Segment to Students CourseList
// comment: Not sure this is needed. Segment should be automatically added on Students segments anyway.
func PostUserSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
