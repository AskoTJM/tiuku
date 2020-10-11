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
// Not needed until we implement personal categories
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
// status:
func PatchSegmentsSegmentSessionsSession(w http.ResponseWriter, r *http.Request) {
	user := r.Header.Get("X-User")

	returnNum := database.CheckIfUserExists(user)
	if returnNum == 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "Incorrect request")
	} else if returnNum > 1 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Problems with the server, please try again later.")
	} else {

		vars := mux.Vars(r)
		seg := vars["session"]
		response := database.StopActiveSession(user, scripts.StringToUint(seg))
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", response)
	}

}
