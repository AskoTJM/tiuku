/*
 * Routing information for tiuku API
 */
package tiuku

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	database "github.com/AskoTJM/tiuku/api/database"
	fac "github.com/AskoTJM/tiuku/api/faculty"
	students "github.com/AskoTJM/tiuku/api/students"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	gotcmd := "Nothing."
	h := r.Header.Get("X-Init")
	if h == "db" {
		database.InitDB()
		gotcmd = "InitDB"
	}
	if h == "populate" {
		database.PopulateSchool()
		database.PopulateStudents()
		gotcmd = "Populating"
	}
	if h == "Hello" {
		gotcmd = "Hello"
	}
	if h == "anonId" {
		gotcmd = database.GetAnonId("oppi1")
	}

	fmt.Fprintf(w, "Welcome to tiuku API %s \nDone %s", now, gotcmd)
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	// Student v1 RAW Routes
	// DELETE yoursef from Segment participation
	Route{
		"DeleteCoursesCourseSegmentsSegment",
		strings.ToUpper("Delete"),
		"/students/v1/courses/{course}/segments/{segment}",
		students.DeleteCoursesCourseSegmentsSegment,
	},

	// DELETE Session from Segment
	Route{
		"DeleteSegmentsSegmentSession",
		strings.ToUpper("Delete"),
		"/students/v1/segments/{segment}/{session}",
		students.DeleteSegmentsSegmentSession,
	},

	// GET available Courses
	Route{
		"GetCourses",
		strings.ToUpper("Get"),
		"/students/v1/courses",
		students.GetCourses,
	},

	// GET Segments for the Course
	Route{
		"GetCoursesCourseSegments",
		strings.ToUpper("Get"),
		"/students/v1/courses/{course}/segments",
		students.GetCoursesCourseSegments,
	},

	// GET data for Segment
	Route{
		"GetCoursesCourseSegmentsSegment",
		strings.ToUpper("Get"),
		"/students/v1/courses/{course}/segments/{segment}",
		students.GetCoursesCourseSegmentsSegment,
	},

	// GET Categories for the Segment
	Route{
		"GetCoursesCourseSegmentsSegmentCategories",
		strings.ToUpper("Get"),
		"/students/v1/courses/{course}/segments/{segment}/categories",
		students.GetCoursesCourseSegmentsSegmentCategories,
	},

	// GET particular Session data from Segment
	Route{
		"GetSegmentsSegmentSession",
		strings.ToUpper("Get"),
		"/students/v1/segments/{segment}/{session}",
		students.GetSegmentsSegmentSession,
	},

	// GET particular Setting of the Segment
	Route{
		"GetSegmentsSegmentSettingsSetting",
		strings.ToUpper("Get"),
		"/students/v1/segments/{segment}/settings/{setting}",
		students.GetSegmentsSegmentSettingsSetting,
	},

	Route{
		"GetUserSegmentSettings",
		strings.ToUpper("Get"),
		"/students/v1/segment/{segment}/settings",
		students.GetUserSegmentSettings,
	},

	Route{
		"GetUserSegments",
		strings.ToUpper("Get"),
		"/students/v1/segments",
		students.GetUserSegments,
	},

	Route{
		"GetUserSegmentsResourceID",
		strings.ToUpper("Get"),
		"/students/v1/segments/{segment}",
		students.GetUserSegmentsResourceID,
	},

	Route{
		"PatchSegmentSegmentSettings",
		strings.ToUpper("Patch"),
		"/students/v1/segment/{segment}/settings",
		students.PatchSegmentSegmentSettings,
	},

	Route{
		"PatchSegmentsSegmentSession",
		strings.ToUpper("Patch"),
		"/students/v1/segments/{segment}/{session}",
		students.PatchSegmentsSegmentSession,
	},

	Route{
		"PatchSegmentsSegmentSettingsSetting",
		strings.ToUpper("Patch"),
		"/students/v1/segments/{segment}/settings/{setting}",
		students.PatchSegmentsSegmentSettingsSetting,
	},

	Route{
		"PostCoursesCourseSegmentsSegment",
		strings.ToUpper("Post"),
		"/students/v1/courses/{course}/segments/{segment}",
		students.PostCoursesCourseSegmentsSegment,
	},

	Route{
		"PostSegmentsSegment",
		strings.ToUpper("Post"),
		"/students/v1/segments/{segment}",
		students.PostSegmentsSegment,
	},

	Route{
		"PostUserSegments",
		strings.ToUpper("Post"),
		"/students/v1/segments",
		students.PostUserSegments,
	},

	Route{
		"PutSegmentsSegmentSession",
		strings.ToUpper("Put"),
		"/students/v1/segments/{segment}/{session}",
		students.PutSegmentsSegmentSession,
	},

	// Faculty v1 RAW Routes

	Route{
		"GetCourses",
		strings.ToUpper("Get"),
		"/faculty/v1/courses",
		fac.GetCourses,
	},

	Route{
		"GetCoursesCourse",
		strings.ToUpper("Get"),
		"/faculty/v1//courses/{course}",
		fac.GetCoursesCourse,
	},

	Route{
		"GetCoursesCourseSegments",
		strings.ToUpper("Get"),
		"/faculty/v1//courses/{course}/segments",
		fac.GetCoursesCourseSegments,
	},

	Route{
		"GetCoursesCourseSegmentsSegment",
		strings.ToUpper("Get"),
		"/faculty/v1//courses/{course}/segments/{segment}",
		fac.GetCoursesCourseSegmentsSegment,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentCategoriesCategory",
		strings.ToUpper("Get"),
		"/faculty/v1//courses/{course}/segments/{segment}/categories/{category}",
		fac.GetCoursesCourseSegmentsSegmentCategoriesCategory,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentCategoriesCategorySettings",
		strings.ToUpper("Get"),
		"/faculty/v1//courses/{course}/segments/{segment}/categories/{category}/settings",
		fac.GetCoursesCourseSegmentsSegmentCategoriesCategorySettings,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentCategoriesCategorySettingsSetting",
		strings.ToUpper("Get"),
		"/faculty/v1//courses/{course}/segments/{segment}/categories/{category}/settings/{setting}",
		fac.GetCoursesCourseSegmentsSegmentCategoriesCategorySettingsSetting,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentSettings",
		strings.ToUpper("Get"),
		"/faculty/v1//courses/{course}/segments/{segment}/categories",
		fac.GetCoursesCourseSegmentsSegmentSettings,
	},

	Route{
		"PostCourses",
		strings.ToUpper("Post"),
		"/faculty/v1/courses",
		fac.PostCourses,
	},

	Route{
		"PostCoursesCourseSegments",
		strings.ToUpper("Post"),
		"/faculty/v1//courses/{course}/segments",
		fac.PostCoursesCourseSegments,
	},

	Route{
		"PostCoursesCourseSegmentsSegmentCategories",
		strings.ToUpper("Post"),
		"/faculty/v1//courses/{course}/segments/{segment}/categories",
		fac.PostCoursesCourseSegmentsSegmentCategories,
	},
}
