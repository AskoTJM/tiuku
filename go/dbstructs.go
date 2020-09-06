package swagger

type User struct {
	ID       int
	Username string
}

type UserModel struct {
	ID      int `gorm:"primary_key"`
	Name    string
	Address string
}

// Table for possible schools
type Schools struct {
	ID      uint `gorm:"primary_key"`
	Finnish string
	English string
}

// Table for different apartmens in t
type Apartments struct {
	ID      uint `gorm:"primary_key"`
	Finnish string
	English string
}

type Campuses struct {
	ID      uint `gorm:"primary_key"`
	Finnish string
	English string
}

type Degrees struct {
	ID      uint `gorm:"primary_key"`
	Finnish string
	English string
}

type Segment struct {
	SegmentNumber          int
	SegmentName            string
	TeacherID              string
	Scope                  uint
	Categories             string
	ExpectedAttendance     uint
	SchoolSegmentsSessions string
}

// Table for course
type Course struct {
	ID              uint `gorm:"primary_key"`
	RID             uint
	School          Schools
	Aparment        Apartments
	Campus          Campuses
	Degree          Degrees
	CourseCode      string
	CourseName      string
	CourseStartDate string
	CourseEndDate   string
	Archived        bool
	Segment         Segment
}

// Where to find students Session listing for this segment.
type SchoolSegmentsSessions struct {
	ID                      uint `gorm:"primary_key"`
	StudentID               string
	StudentSegmentsSessions string
	Privacy                 string
}

type SegmentCategories struct {
	Id                 uint `gorm:"primary_key"`
	MainCategory       string
	SubCategory        string
	MandatoryToTrack   bool
	MandatoryToComment bool
	Tickable           bool
	LocationNeeded     bool
	Active             bool
}
