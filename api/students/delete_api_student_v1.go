package students

/*
// delete_api_student_v1.go
// Description: DELETE request functions for Student users
// CRUD Delete Colelction  405 (Method Not Allowed) Single: 200 (Ok), 404 (Not Found)
*/

import (
	"fmt"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/gorilla/mux"
)

// Leave {segment} of the {course}
// W0rks
func DeleteCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	segCode := vars["segment"]
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
		//studentToJoin := database.GetStudentUser(user)
		res := database.DeleteStudentFromSegment(user, scripts.StringToUint(segCode))
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", res)
	}
}

// Remove {session} from {segment}, SoftDelete
// W0rks
func DeleteSegmentsSegmentSessionsSession(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("X-User")

	returnNum := database.CheckIfUserExists(user)
	if returnNum == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "Incorrect request")
	} else if returnNum > 1 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Problems with the server, please try again later.")
	} else {

		vars := mux.Vars(r)
		resSes := vars["session"]
		res := database.DeleteSessionFromStudent(user, scripts.StringToUint(resSes))
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		//res := "NotReadyYet"
		fmt.Fprintf(w, "%s", res)
	}
}
