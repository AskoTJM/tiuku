package students

/*
// put_api_student_v1.go
// Description: PUT request functions for Student users
*/
import "net/http"

// Replace {session} of the {segment} with new data
// status:
func PutSegmentsSegmentSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
