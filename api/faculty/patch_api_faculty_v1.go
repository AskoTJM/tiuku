/*
// patch_api_faculty_v1.go
// Description: PATCH request functions for Faculty users
//
*/

package faculty

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/gorilla/mux"
)

// Update Student Users information
// W0rks Maybe add check to see if data is actually changing?
func PatchStudentsStudent(w http.ResponseWriter, r *http.Request) {

	var response string

	resString, resCode := database.CheckJSONContent(w, r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if resCode != http.StatusOK {
		w.WriteHeader(resCode)
		response = "Error: " + resString
	} else {
		var tempStudentNewData database.StudentUser
		vars := mux.Vars(r)
		studentToEdit := vars["student"]
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		err := dec.Decode(&tempStudentNewData)
		if err != nil {
			log.Printf("Error. Issue with decoding JSON. %v", err)
		} else {
			tempStudentNewData.ID = scripts.StringToUint(studentToEdit)
			res := database.UpdateStudentUser(tempStudentNewData)
			if res {
				log.Printf("Error. Could not update student user data.")
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				response = "Student user data updated succesfully." // resString + studentToEdit
				w.WriteHeader(http.StatusOK)
			}
		}
	}
	fmt.Fprintf(w, "%s", response)
}

// Change Course Setting, also for archiving
// W1P
func PatchCoursesCourse(w http.ResponseWriter, r *http.Request) {
	var response string = "test"

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", response)

}
