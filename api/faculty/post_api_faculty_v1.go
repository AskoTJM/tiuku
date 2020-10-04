package faculty

/*
// post_api_faculty_v1.go
// Description: POST request functions for Faculty users
*/
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/AskoTJM/tiuku/api/database"
)

// desc: Create new Course in course table
// Status: Need to clean and re-think, but works.
// Probably doesn't need so much to response stuff
func PostCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//w.WriteHeader(http.StatusOK)

	rbody := database.CreateCourse(r)
	if rbody != "Error" {
		t := &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			//Header:     //map[string][]string{},
			Body: ioutil.NopCloser(bytes.NewBufferString(rbody)),
			//ContentLength:    0,
			//TransferEncoding: []string{},
			///Request: r,
			//TLS:              &tls.ConnectionState{},
		}
		buff := bytes.NewBuffer(nil)
		err := t.Write(buff)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, "%s", buff)
	}

}

// desc: New segment for the course
// status: Works
func PostCoursesCourseSegments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	result := database.CreateSegment(r)
	anon, _ := json.Marshal(result)
	n := len(anon)
	s := string(anon[:n])

	fmt.Fprintf(w, "%s", s)

}

// desc: Add New categories for segment
// status:
func PostCoursesCourseSegmentsSegmentCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	database.CreateCategoriesForSegment(r)
}
