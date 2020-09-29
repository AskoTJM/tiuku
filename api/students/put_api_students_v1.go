package students

import "net/http"

// desc: Replace {session} of the {segment} with new data
// status:
func PutSegmentsSegmentSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
