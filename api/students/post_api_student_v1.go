package students

import "net/http"

// desc: Join {segment} of the {course}
// status:
func PostCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// desc: Add or start new session to {segment} tab
// status:
func PostSegmentsSegment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// desc: Add new Segment to CourseList
// comment: Not sure if this is needed, should be happen automatically when we join segment
func PostUserSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
