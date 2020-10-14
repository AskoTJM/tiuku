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

	vars := mux.Vars(r)
	segCode := vars["segment"]

	user := r.Header.Get("X-User")
	returnNum := database.CheckIfUserExists(user)
	if returnNum == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "Incorrect request")
	} else if returnNum > 1 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Problems with the server, please try again later.")
	} else {

		res := database.AddStudentToSegment(user, scripts.StringToUint(segCode))
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", res)
	}

}

// Add or start NEW session to {segment} table with empty body the start time is inserted automatically
// W0rks
func PostSegmentsSegmentSessions(w http.ResponseWriter, r *http.Request) {

	user := r.Header.Get("X-User")
	var response string

	returnNum := database.CheckIfUserExists(user)
	if returnNum == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "Incorrect request")
	} else if returnNum > 1 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Problems with the server, please try again later.")
	} else {

		res := database.CheckJSONContent(w, r)
		if res == "TYPE_ERROR" {
			log.Printf("Type Error with body.")
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			response = "Incorrect body type."
		} else {
			var session database.StudentSegmentSession
			// If body has content and is JSON then...
			if res == "PASS" {
				dec := json.NewDecoder(r.Body)
				dec.DisallowUnknownFields()
				err := dec.Decode(&session)
				if err != nil {
					log.Println(err)
				}
				vars := mux.Vars(r)
				seg := vars["segment"]
				session.SegmentID = scripts.StringToUint(seg)
				// Category check

				//test, result := database.CheckIfCategoryMatchSegment(session.Category, scripts.StringToUint(seg))
				test, result := database.CheckIfSessionMatchesCategory(session)

				if test {

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
					// Set/OverWrite values set by the system

					session.SegmentID = scripts.StringToUint(seg)
					session.Version = 1
					session.Created = time.Now().Format(time.RFC3339)
					session.Updated = time.Now().Format(time.RFC3339)

					response2 := database.StartSessionOnSegment(user, session)
					if response2 {
						w.Header().Set("Content-Type", "application/json; charset=UTF-8")
						w.WriteHeader(http.StatusOK)
						response = "Session added to database."
					} else {
						w.Header().Set("Content-Type", "application/json; charset=UTF-8")
						w.WriteHeader(http.StatusInternalServerError)
						response = "Server error."
					}
				} else {
					log.Printf("Incorrect category. %v", r)
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(http.StatusBadRequest)
					//response = "Incorrect category for Segment."
				}
				response = result + "  " + response

			} else if res == "EMPTY" {
				log.Printf("Empty JSON Body: Minimum required data not provided. %v", r)
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusBadRequest)
				response = "Empty JSON Body: Minimum required data not provided."

			}

		}
		fmt.Fprintf(w, "%s", response)
	}
}

// Add Segment to Students CourseList
// comment: This is needed unless implemented personal segments.
// Segments on School should be automatically added on Students segments when joining them anyway.
func PostUserSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
