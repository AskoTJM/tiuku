package students

/*
// delete_api_student_v1.go
// Description: DELETE request functions for Student users
// CRUD Delete Colelction  405 (Method Not Allowed) Single: 200 (Ok), 404 (Not Found)
*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/gorilla/mux"
)

// Leave {segment} of the {course}
// W0rks need to add check if actually enrolled for that.
func DeleteCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	segCode := vars["segment"]
	var resCode int
	var resString string
	var response string
	//resSeg := database.GetSegmentDataById(scripts.StringToUint(segCode))

	user := r.Header.Get("X-User")
	resString, resCode = database.CheckIfUserExists(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if resCode != http.StatusOK {
		w.WriteHeader(resCode)
		response = resString
	} else {
		//studentToJoin := database.GetStudentUser(user)
		resCheck := database.CheckSegmentParticipation(user, scripts.StringToUint(segCode))
		if resCheck == 0 {
			response = "Not participating to this Segment"
		} else if resCheck == 1 {
			res := database.DeleteStudentFromSegment(user, scripts.StringToUint(segCode))
			if res {
				w.WriteHeader(http.StatusInternalServerError)
				response = "Error with removing student user from Segment. "
			} else {
				w.WriteHeader(http.StatusOK)
				response = "Succesfully removed student user from Segment"
			}
		} else {
			log.Printf("Error: In <delete_api_student_v1.go->DeleteCoursesCourseSegmentsSegment>")
			w.WriteHeader(http.StatusInternalServerError)
			response = "Error. Incoherent data on server."
		}

	}
	fmt.Fprintf(w, "%s", response)
}

// Remove {session} from {segment}, SoftDelete
// W0rks
func DeleteSegmentsSegmentSessionsSession(w http.ResponseWriter, r *http.Request) {
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
		resSes := vars["session"]
		res := database.DeleteSessionFromStudent(user, scripts.StringToUint(resSes))
		if res {
			w.WriteHeader(http.StatusInternalServerError)
			response = "Error: Deleting Session failed."
		} else {
			w.WriteHeader(http.StatusOK)
			response = "Session Deleted."
		}
	}
	fmt.Fprintf(w, "%s", response)
}
