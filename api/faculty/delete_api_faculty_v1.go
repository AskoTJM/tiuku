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

	vars := mux.Vars(r)
	seg := vars["student"]
	//user := r.Header.Get("X-User")
	user := database.GetStudentUserWithID(scripts.StringToUint(seg))
	resString, resCode = database.CheckIfUserExists(user.StudentID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if resCode != http.StatusOK {
		response = resString
	} else {
		//log.Printf("http.StatusOk resCode is %v", resCode)
		//result := database.DeleteStudentUser(user.StudentID)
		result := database.DeleteStudentFromAllSegments(user.StudentID)
		if result {
			log.Printf("Error in <database/delete_api_v1.go->DeleteStudentsStudent>")
			resCode = http.StatusInternalServerError
			response = "Error in removing user."
		} else {
			result2 := database.DeleteStudentUser(user.StudentID)
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
	fmt.Fprintf(w, "%s", response)
}
