/*
 * facultyAPI
 *
 * API for faculty members to access Tiuku
 *
 * API version: 1.0
 * Contact: asko.mattila@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package faculty

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/AskoTJM/tiuku/api/database"
	"github.com/gorilla/mux"
)

// desc: List of active courses
// status: works
func GetCourses(w http.ResponseWriter, r *http.Request) {

	result := database.GetCourses(r)

	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)
}

// desc: Get {course} information
// status: Works. Doesn't give any information about segments
func GetCoursesCourse(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	courseCode := vars["course"]
	result := database.FindCourseTableById(courseCode)

	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)
}

// desc: Get list of segments for {course}
// status: Works
func GetCoursesCourseSegments(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	courseCode := vars["course"]
	// Get course information
	result := database.FindCourseTableById(courseCode)
	// Get segment data
	result2 := database.FindSegmentTableByCourseId(result.ID)
	//Transform results to json
	anon, _ := json.Marshal(result2)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)
}

// desc: Get data/sessions of the {segment} in the {course}
// status: Works
func GetCoursesCourseSegmentsSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	//courseCode := vars["course"]
	segCode := vars["segment"]
	// Get course information
	//courseRes := database.FindCourseTableById(courseCode)
	// Get segment data
	segRes := database.FindSegmentDataById(segCode)
	//Transform results to json
	anon, _ := json.Marshal(segRes)
	n := len(anon)
	s := string(anon[:n])

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", s)
}

// desc: Get categories for the {Segment}
// status:
func GetCoursesCourseSegmentsSegmentCategoriesCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// desc: Get settings for {category}
// status:
func GetCoursesCourseSegmentsSegmentCategoriesCategorySettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// desc: Change(GET) {setting} in {category}
// status: Unnecessary? Better way to do this. Or least should be PUT/PATCH
// ToDo: Remove or repurpose
func GetCoursesCourseSegmentsSegmentCategoriesCategorySettingsSetting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// desc: Get settings for {segment} of the {course}
// status: Not sure about this one either
// todo: Think about this.
func GetCoursesCourseSegmentsSegmentSettings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

// desc: Get segments table for Faculty User
// status: WIP
func GetUserSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	url, _ := url.Parse(r.RequestURI)
	path := url.Path
	uriParts := strings.Split(path, "/")
	log.Printf("%s", uriParts)

}
