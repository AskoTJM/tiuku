package students

/*
// patch_api_student_v1.go
// Description: PATCH request functions for Student users
*/

import "net/http"

// desc: Change settings for the {segment}
func PatchSegmentSegmentSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// desc: Stop session by inserting Stop time
// status:
func PatchSegmentsSegmentSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

//desc: Update {setting} for {segment}
func PatchSegmentsSegmentSettingsSetting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
