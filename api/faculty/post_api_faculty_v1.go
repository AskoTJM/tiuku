package faculty

/*
// post_api_faculty_v1.go
// Description: POST request functions for Faculty users
*/
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/AskoTJM/tiuku/api/scripts"
	"github.com/gorilla/mux"
)

// Create new Course in course table
// W0rks
func PostCourses(w http.ResponseWriter, r *http.Request) {

	res := database.CheckJSONContent(w, r)
	if res != "PASS" {
		fmt.Fprintf(w, "%s", res)
	} else {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var newCourse database.Course
		err := dec.Decode(&newCourse)
		if err != nil {
			log.Println(err)
		}
		result := database.CreateCourse(newCourse, database.CourseTableToEdit)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", result)
	}

}

// New segment for the course
// W0rks
func PostCoursesCourseSegments(w http.ResponseWriter, r *http.Request) {

	res := database.CheckJSONContent(w, r)
	if res != "PASS" {
		fmt.Fprintf(w, "%s", res)
	} else {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var newSegment database.Segment
		err := dec.Decode(&newSegment)
		if err != nil {
			log.Println(err)
		}
		vars := mux.Vars(r)
		courseCode := vars["course"]
		newSegment.CourseID = scripts.StringToUint(courseCode)
		result := database.CreateSegment(newSegment, database.SegmentTableToEdit)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", result)
	}

}

// Add New category for segment
// W0rks
func PostCoursesCourseSegmentsSegmentCategories(w http.ResponseWriter, r *http.Request) {

	res := database.CheckJSONContent(w, r)
	if res != "PASS" {
		fmt.Fprintf(w, "%s", res)
	} else {
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		var newCategory database.SegmentCategory
		err := dec.Decode(&newCategory)
		if err != nil {
			log.Println(err)
		}
		vars := mux.Vars(r)
		segmentCode := vars["segment"]
		newCategory.SegmentID = scripts.StringToUint(segmentCode)
		result := database.CreateCategory(newCategory, database.CategoriesTableToEdit)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "%s", result)
	}

}

// Add New School
// T0D0

// Add New Campus
// T0D0

// Add New Apartment
// T0D0

// Add New Degree
// T0D0
