package students

/*
// put_api_student_v1.go
// Description: PUT request functions for Student users
// CRUD: Update/Replace Collection: 405(Method Not Allowed) if not replacing all Single: 200(OK) or 204(No Content), 404 (Not Found)
*/
import "net/http"

// Replace {session} of the {segment} with new data
// status:
func PutSegmentsSegmentSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
