package faculty

/*
// post_api_faculty_v1.go
// Description: POST request functions for Faculty users
*/
import (
	"fmt"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
)

// Create new Course in course table
// Status: Works
func PostCourses(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)

	result := database.CreateCourse(w, r)
	//anon, _ := json.Marshal(result)
	//n := len(anon)
	//s := string(anon[:n])
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", result)

}

// New segment for the course
// status: Works
func PostCoursesCourseSegments(w http.ResponseWriter, r *http.Request) {

	result := database.CreateSegment(w, r)
	//anon, _ := json.Marshal(result)
	//n := len(anon)
	//s := string(anon[:n])
	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", result)

}

// Add New category for segment
// status:
func PostCoursesCourseSegmentsSegmentCategories(w http.ResponseWriter, r *http.Request) {

	result := database.CreateCategory(w, r)
	fmt.Fprintf(w, "%s", result)
}
