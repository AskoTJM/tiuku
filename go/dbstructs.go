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

// Table for StudentUsers
type StudentUsers struct {
	ID              uint `gorm:"primary_key"`
	StudentID       string
	AnonID          string
	StudentName     string
	StudentSegments StudentSegments `gorm:"foreignkey:ID"`
	StudentEmail    string
	StudentClass    string
}

// Table for FacultyUsers of the API
type FacultyUsers struct {
	ID              uint `gorm:"primary_key"`
	FacultyID       string
	FacultyName     string
	FacultyEmail    string
	FacultySegments FacultySegments
}

// Table of Faculty(in this case Teachers), to save their Segments
type FacultySegments struct {
	ID                     uint `gorm:"primary_key"`
	ResourceID             string
	Course                 Course
	SegmentNumber          uint
	SchoolSegmentsSessions SchoolSegmentsSessions
	SegmentCategories      SegmentCategories
	Archived               bool
}

// Table for course
type Course struct {
	ID              uint `gorm:"primary_key"`
	ResourceID      uint
	School          Schools
	CourseCode      string
	CourseName      string
	CourseStartDate string
	CourseEndDate   string
	Archived        bool
	Segment         []Segment
}

// Course can have one or more Segments
type Segment struct {
	ID                     uint `gorm:"primary_key"`
	SegmentName            string
	TeacherID              string
	Scope                  uint
	SegmentCategories      SegmentCategories
	ExpectedAttendance     uint
	SchoolSegmentsSessions SchoolSegmentsSessions
}

// Schools Segment table has data for students and where to find their Session for the Segement.
type SchoolSegmentsSessions struct {
	ID                      uint `gorm:"primary_key"`
	StudentID               string
	StudentSegmentsSessions string
	Privacy                 string
}

// Segment has different Categories for tracking and settings for them.
// All SubCategories belong in to one of the three MainCategory
// Maybe should be belongs to or one-to-one with two structs?
type SegmentCategories struct {
	ID                 uint `gorm:"primary_key"`
	MainCategory       string
	SubCategory        string
	MandatoryToTrack   bool
	MandatoryToComment bool
	Tickable           bool
	LocationNeeded     bool
	Active             bool
}

// What Segments of Courses student is tracking.
type StudentSegments struct {
	ID                     uint `gorm:"primary_key"`
	ResourceID             string
	Course                 Course
	SegmentNumber          uint
	StudentSegmentSessions StudentSegmentSessions
	SegmentCategories      SegmentCategories
	Archived               bool
}

// Students Sessions for Segment
type StudentSegmentSessions struct {
	ID         uint `gorm:"primary_key"`
	ResourceID uint
	StartTime  string
	EndTime    string
	CreatedAt  string
	UpdateAt   string
	DeletedAt  string
	Categories SegmentCategories
	Comment    string
	Version    uint
	Locations  string
}

// Table for Schools. School can have multiple campuses
type Schools struct {
	ID       uint `gorm:"primary_key"`
	Finnish  string
	English  string
	Campuses []Campuses
}

// Campus of the School, Campus can have multiple Apartments
type Campuses struct {
	ID         uint `gorm:"primary_key"`
	Finnish    string
	English    string
	Apartments []Apartments
}

// Table for different Apartments in Campus, Apartment can have multiple Degrees
type Apartments struct {
	ID      uint `gorm:"primary_key"`
	Finnish string
	English string
	Degrees []Degrees
}

// Degrees in the Apartment.
type Degrees struct {
	ID      uint `gorm:"primary_key"`
	Finnish string
	English string
}
