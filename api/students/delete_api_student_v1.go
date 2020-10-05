package students

/*
// delete_api_student_v1.go
// Description: DELETE request functions for Student users
// CRUD Delete Colelction  405 (Method Not Allowed) Single: 200 (Ok), 404 (Not Found)
*/

import "net/http"

// Leave segment of the course
func DeleteCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Remove {session} from {segment}
func DeleteSegmentsSegmentSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
