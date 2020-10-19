package students

/*
// patch_api_student_v1.go
// Description: PATCH request functions for Student users
// CRUD: Update/Modify Collection: 405 (Method Not Allowed) Single: 200 (OK), 204( No Content), 404 (Not Found)
*/

import (
	"fmt"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/gorilla/mux"
)

// Change settings for the {segment}
// Not needed until we implement personal categories?
func PatchSegmentSegmentSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Update {setting} for {segment}
//
func PatchSegmentsSegmentSettingsSetting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Stop session by inserting Stop time
// W0rks
func PatchSegmentsSegmentSessionsSession(w http.ResponseWriter, r *http.Request) {
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
		seg := vars["session"]
		//var response string
		result := database.StopActiveSession(user, scripts.StringToUint(seg))
		if result {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			response = "Session stopped"
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			response = "Error with stopping Session."
		}

	}
	fmt.Fprintf(w, "%s", response)

}
