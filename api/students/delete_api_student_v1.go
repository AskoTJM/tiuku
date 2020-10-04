package students

/*
// delete_api_student_v1.go
// Description: DELETE request functions for Student users
*/

import "net/http"

// desc: Leave segment of the course
func DeleteCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// desc: Remove {session} from {segment}
func DeleteSegmentsSegmentSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
