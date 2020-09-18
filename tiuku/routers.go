/*
 * StudentAPI
 *
 * API for Students to acccess the Tiuku system.
 *
 * API version: 1.0
 * Contact: asko.mattila@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tiuku

import (
	"fmt"
	"net/http"
	"strings"

	fac "github.com/AskoTJM/tiuku/tiuku/faculty"
	students "github.com/AskoTJM/tiuku/tiuku/students"
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
	fmt.Fprintf(w, "Welcome to tiuku API")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"DeleteCoursesCourseSegmentsSegment",
		strings.ToUpper("Delete"),
		"/students/v1/courses/{course}/segments/{segment}",
		students.DeleteCoursesCourseSegmentsSegment,
	},

	Route{
		"DeleteSegmentsSegmentSession",
		strings.ToUpper("Delete"),
		"/students/v1/segments/{segment}/{session}",
		students.DeleteSegmentsSegmentSession,
	},

	Route{
		"GetCourses",
		strings.ToUpper("Get"),
		"/students/v1/courses",
		students.GetCourses,
	},

	Route{
		"GetCoursesCourseSegments",
		strings.ToUpper("Get"),
		"/students/v1/courses/{course}/segments",
		students.GetCoursesCourseSegments,
	},

	Route{
		"GetCoursesCourseSegmentsSegment",
		strings.ToUpper("Get"),
		"/students/v1/courses/{course}/segments/{segment}",
		students.GetCoursesCourseSegmentsSegment,
	},

	Route{
		"GetCoursesCourseSegmentsSegmentCategories",
		strings.ToUpper("Get"),
		"/students/v1/courses/{course}/segments/{segment}/categories",
		students.GetCoursesCourseSegmentsSegmentCategories,
	},

	Route{
		"GetSegmentsSegmentSession",
		strings.ToUpper("Get"),
		"/students/v1/segments/{segment}/{session}",
		students.GetSegmentsSegmentSession,
	},

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

	// Faculty routes

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