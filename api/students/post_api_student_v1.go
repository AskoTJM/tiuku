package students

/*
// post_api_student_v1.go
// Description: POST request functions for Student users
// CRUD: Create, should return 201 on collection with Id, on Specific 404 (Not Found) or 409 (Conflict)
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/gorilla/mux"
)

// Join {segment} of the {course}
// W0rks
func PostCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {

	user := r.Header.Get("X-User")
	var res string
	var response string

	resString, resCode := database.CheckIfUserExists(user)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if resCode != http.StatusOK {
		w.WriteHeader(resCode)
		res = resString
	} else {
		vars := mux.Vars(r)
		segCode := vars["segment"]
		resCheck := database.CheckSegmentParticipation(user, scripts.StringToUint(segCode))
		if resCheck == 0 {
			response = database.AddStudentToSegment(user, scripts.StringToUint(segCode))
			w.WriteHeader(http.StatusOK)
		} else if resCheck == 1 {
			response = "Error. Already participating to this Segment. \n"
			w.WriteHeader(http.StatusBadRequest)
		} else {
			log.Printf("Error. Incorrect value in <database/post_api_student_v1.go->PostCoursesCourseSegmentsSegment>.\n")
			response = "Error. Server data incoherent."
			w.WriteHeader(http.StatusInternalServerError)
		}
		res = response
	}
	fmt.Fprintf(w, "%s", res)
}

// Add or start NEW session to {segment} table
// W0rks
func PostSegmentsSegmentSessions(w http.ResponseWriter, r *http.Request) {

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
		resJsonString, resJsonCode := database.CheckJSONContent(w, r)
		if resJsonCode != http.StatusOK {
			w.WriteHeader(resCode)
			response = resJsonString
		} else {
			var session database.StudentSegmentSession
			// If body has content and is JSON then...
			//var tempOldSession database.StudentSegmentSession
			// Check if there is open Session
			tempOldSession := database.GetLastSession(user)
			if tempOldSession.EndTime == database.StringForEmpy {
				database.StopActiveSession(user, tempOldSession.ID)
			}
			dec := json.NewDecoder(r.Body)
			dec.DisallowUnknownFields()
			err := dec.Decode(&session)
			if err != nil {
				log.Println(err)
			}
			vars := mux.Vars(r)
			seg := vars["segment"]
			// Default settings for new Session
			session.SegmentID = scripts.StringToUint(seg)
			session.Version = 1
			session.Created = time.Now().Format(time.RFC3339)
			session.Updated = time.Now().Format(time.RFC3339)
			if session.StartTime == "" {
				session.StartTime = time.Now().Format(time.RFC3339)
				// If there is no StartTime there should not be EndTime or Deleted
				session.EndTime = database.StringForEmpy
				session.Deleted = database.StringForEmpy
			}
			if session.EndTime == "" {
				session.EndTime = database.StringForEmpy
				//If there is no EndTime, should not be Deleted
				session.Deleted = database.StringForEmpy
			}
			if session.Deleted == "" {
				session.Deleted = database.StringForEmpy
			}
			if session.Privacy == (database.StudentSegmentSession{}.Privacy) {
				session.Privacy = true
			}

			resString, resBool := database.ValidateNewSessionStruct(session)

			if resBool {
				log.Printf("Result from Validity test %v", resString)
				response = response + " " + resString
			} else {
				response2 := database.CreateNewSessionOnSegment(user, session)
				if response2 {
					w.WriteHeader(http.StatusInternalServerError)
					response = resString + " Server error."
				} else {
					w.WriteHeader(http.StatusCreated)
					response = resString + " New Session Started"
				}
				response = response + " " + resString
			}
			response = response + " " + resString
		}

	}
	fmt.Fprintf(w, "%s", response)
}

// Add Segment to Students CourseList
// comment: This isn't needed unless implemented personal segments.
// Segments on School should be automatically added on Students segments when joining them anyway.
func PostUserSegments(w http.ResponseWriter, r *http.Request) {
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
