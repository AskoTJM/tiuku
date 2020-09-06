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
	StudentSegments StudentSegments
	StudentEmail    string
	StudentClass    string
}

// Table for FacultyUsers
type FacultyUsers struct {
	ID           uint `gorm:"primary_key"`
	FacultyID    string
	FacultyEmail string
}

// Table for possible schools. School can have multiple campuses
type Schools struct {
	ID         uint `gorm:"primary_key"`
	Finnish    string
	English    string
	Apartments []Apartments
}

// Campus of the School, Campus can have multiple Apartments
type Campuses struct {
	ID      uint `gorm:"primary_key"`
	Finnish string
	English string
	Degrees []Degrees
}

// Table for different Apartments in Campus, Apartment can have multiple Degrees
type Apartments struct {
	ID       uint `gorm:"primary_key"`
	Finnish  string
	English  string
	Campuses []Campuses
}

// Degrees in the Apartment.
type Degrees struct {
	ID      uint `gorm:"primary_key"`
	Finnish string
	English string
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
	Categories             string
	ExpectedAttendance     uint
	SchoolSegmentsSessions string
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
