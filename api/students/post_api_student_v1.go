package students

/*
// post_api_student_v1.go
// Description: POST request functions for Student users
*/

import "net/http"

// Join {segment} of the {course}
// status:
func PostCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Add or start new session to {segment} tab
// status:
func PostSegmentsSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Add new Segment to CourseList
// comment: Not sure if this is needed, should be happen automatically when we join segment
func PostUserSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
