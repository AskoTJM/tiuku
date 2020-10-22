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

// Update Faculty Users information
// T35T Maybe add check to see if data is actually changing?
func PatchFacultyFaculty(w http.ResponseWriter, r *http.Request) {

	var response string

	resString, resCode := database.CheckJSONContent(w, r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if resCode != http.StatusOK {
		w.WriteHeader(resCode)
		response = "Error: " + resString
	} else {
		var tempFacultyNewData database.FacultyUser
		vars := mux.Vars(r)
		facultyToEdit := vars["faculty"]
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		err := dec.Decode(&tempFacultyNewData)
		if err != nil {
			log.Printf("Error. Issue with decoding JSON. %v", err)
		} else {
			tempFacultyNewData.ID = scripts.StringToUint(facultyToEdit)
			log.Printf("Updating data for %v \n", tempFacultyNewData.FacultyName)
			res := database.UpdateFacultyUser(tempFacultyNewData)
			if res {
				log.Printf("Error. Could not update faculty user data.")
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				response = "Faculty user data updated succesfully." // resString + studentToEdit
				w.WriteHeader(http.StatusOK)
			}
		}
	}
	fmt.Fprintf(w, "%s", response)
}

// Change Course Setting, also for archiving
// W0rks
func PatchCoursesCourse(w http.ResponseWriter, r *http.Request) {

	var response string
	resString, resCode := database.CheckJSONContent(w, r)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if resCode != http.StatusOK {
		w.WriteHeader(resCode)
		response = "Error: " + resString
	} else {
		var tempCourse database.Course
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		err := dec.Decode(&tempCourse)
		if err != nil {
			log.Printf("Error. Issue with decoding JSON. %v", err)
		} else {
			vars := mux.Vars(r)
			courseToEdit := vars["course"]
			tempCourse.ID = scripts.StringToUint(courseToEdit)
			resString, res := database.UpdateCourse(tempCourse)
			if res {
				w.WriteHeader(http.StatusInternalServerError)
				response = resString
			} else {
				response = resString
				w.WriteHeader(http.StatusOK)
			}
		}
	}
	fmt.Fprintf(w, "%s", response)

}
