/*
//
//
*/

package faculty

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/gorilla/mux"
)

// Remove {Student} identifying data.
// W0rks
func DeleteStudentsStudent(w http.ResponseWriter, r *http.Request) {

	var resString string
	var resCode int
	var response string

	user := r.Header.Get("X-User")
	resF := database.GetFacultyUser(user)
	if resF.ID == 0 {
		response = "Access denied."
		w.WriteHeader(http.StatusBadRequest)
	} else if !resF.Admin {
		response = "Insuffient rights."
		w.WriteHeader(http.StatusBadRequest)
	} else {

		vars := mux.Vars(r)
		seg := vars["student"]
		tempStudent := database.GetStudentUserWithID(scripts.StringToUint(seg))
		resString, resCode = database.CheckIfUserExists(tempStudent.StudentID)
		//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if resCode != http.StatusOK {
			response = resString
		} else {
			result := database.DeleteStudentFromAllSegments(tempStudent.StudentID)
			if result {
				resCode = http.StatusInternalServerError
				response = "Error in removing user."
			} else {
				result2 := database.DeleteStudentUser(tempStudent.StudentID)
				//result2 := database.DeleteStudentFromAllSegments(user.StudentID)
				if result2 {
					log.Printf("Error in <database/delete_api_v1.go->DeleteStudentsStudent>")
					resCode = http.StatusInternalServerError
					response = "Error in removing user."
				} else {
					resCode = http.StatusOK
					response = "Student User succesfully removed."
				}
			}
		}
		w.WriteHeader(resCode)
	}
	fmt.Fprintf(w, "%s", response)
}
