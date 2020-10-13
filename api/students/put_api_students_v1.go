package students

/*
// put_api_student_v1.go
// Description: PUT request functions for Student users
// CRUD: Update/Replace Collection: 405(Method Not Allowed) if not replacing all Single: 200(OK) or 204(No Content), 404 (Not Found)
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

// Replace {session} of the {segment} with new data
// T0D0
func PutSegmentsSegmentSessionsSession(w http.ResponseWriter, r *http.Request) {

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
				//log.Println("JSON PASS")
				dec := json.NewDecoder(r.Body)
				dec.DisallowUnknownFields()
				err := dec.Decode(&session)
				if err != nil {
					log.Println(err)
				}

				vars := mux.Vars(r)
				seg := vars["segment"]
				ses := vars["session"]
				if database.DebugMode {
					log.Printf("%s", ses)
				}
				// Category check here? Is it set and it's approved one?
				if session.Category == 0 {
					w.Header().Set("Content-Type", "application/json; charset=UTF-8")
					w.WriteHeader(http.StatusBadRequest)
					response = "Empty JSON Body: Category not provided or incorrect one."
				} else {

					// Code for handling of the body
					if session.StartTime == "" {

						//session.StartTime = time.Now().Format(time.RFC3339)
						// If there is no StartTime there should not be EndTime or Deleted
						//session.EndTime = database.StringForEmpy
						//session.Deleted = database.StringForEmpy
					}
					if session.EndTime == "" {
						session.EndTime = database.StringForEmpy
						//If there is no EndTime, should not be Deleted
						//session.Deleted = database.StringForEmpy
					}
					if session.Deleted == "" {
						session.Deleted = database.StringForEmpy
					}
					// Set/OverWrite values set by the system

					session.SegmentID = scripts.StringToUint(seg)
					session.Version = 2
					session.Created = time.Now().Format(time.RFC3339)
					session.Updated = time.Now().Format(time.RFC3339)

					response2 := database.ReplaceSession(scripts.StringToUint(seg), session)
					if response2 {
						w.Header().Set("Content-Type", "application/json; charset=UTF-8")
						w.WriteHeader(http.StatusOK)
					} else {
						w.Header().Set("Content-Type", "application/json; charset=UTF-8")
						w.WriteHeader(http.StatusInternalServerError)
					}
				}
				// If there is no content
			} else if res == "EMPTY" {
				log.Printf("Empty JSON Body: Minimum of required data not provided. %v", r)
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusBadRequest)
				response = "Empty JSON Body: Minimum of required data not provided."
			}

		}
		fmt.Fprintf(w, "%s", response)
	}
}
