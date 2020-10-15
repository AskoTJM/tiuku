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
// W0rks
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
				dec := json.NewDecoder(r.Body)
				dec.DisallowUnknownFields()
				err := dec.Decode(&session)
				if err != nil {
					log.Println(err)
				}

				vars := mux.Vars(r)
				seg := vars["segment"]
				ses := vars["session"]
				session.SegmentID = scripts.StringToUint(seg)
				session.Created = time.Now().Format(time.RFC3339)
				session.Updated = time.Now().Format(time.RFC3339)
				session.ResourceID = scripts.StringToUint(ses)
				if session.Deleted == "" {
					session.Deleted = database.StringForEmpy
				}

				resBool, resString := database.ValidateNewSessionStruct(session)

				if !resBool {
					log.Printf("Result from Validity test %v", resString)
					response = response + " " + resString
				} else {
					/*
						test, result, _ := database.CheckIfSessionMatchesCategory(session)
						if test {
					*/

					responseBool, resString := database.ReplaceSession(user, scripts.StringToUint(ses), session)
					if responseBool {
						w.Header().Set("Content-Type", "application/json; charset=UTF-8")
						w.WriteHeader(http.StatusOK)
					} else {
						w.Header().Set("Content-Type", "application/json; charset=UTF-8")
						w.WriteHeader(http.StatusInternalServerError)
					}
					response = res + " & " + resString

				}

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
