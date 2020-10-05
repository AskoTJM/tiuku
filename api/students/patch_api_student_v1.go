package students

/*
// patch_api_student_v1.go
// Description: PATCH request functions for Student users
// CRUD: Update/Modify Collection: 405 (Method Not Allowed) Single: 200 (OK), 204( No Content), 404 (Not Found)
*/

import "net/http"

// Change settings for the {segment}
func PatchSegmentSegmentSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Stop session by inserting Stop time
// status:
func PatchSegmentsSegmentSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

//Update {setting} for {segment}
func PatchSegmentsSegmentSettingsSetting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
