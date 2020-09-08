package swagger

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

// Table for FacultyUsers f the API
type FacultyUsers struct {
	ID              uint `gorm:"primary_key"`
	FacultyID       string
	FacultyName     string
	FacultyEmail    string
	FacultySegments FacultySegments
}

// Table of Faculty(in this cse Teachers), to save their Segments
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
	Schools         Schools
	CourseCode      string
	CourseName      string
	CourseStartDate string
	CourseEndDate   string
	Archived        bool
	Segment         []Segment
}

// Course can have on or more Segments
type Segment struct {
	ID                     uint `gorm:"primary_key"`
	SegmentName            string
	TeacherID              string
	Scope                  uint
	SegmentCategories      SegmentCategories
	ExpectedAttendance     uint
	SchoolSegmentsSessions SchoolSegmentsSessions
}

// Schools Segment table has data fo students and where to find their Session for the Segement.
type SchoolSegmentsSessions struct {
	ID                      uint `gorm:"primary_key"`
	StudentID               string
	StudentSegmentsSessions string
	Privacy                 string
}

// Segment has different Categories for tracking and settings fr them.
// All SubCategories belong in to one of the three MainCategoy
// Maybe should be belongs to o one-to-one with two structs?
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

// What Segments of Courses sudent is tracking.
type StudentSegments struct {
	ID                     uint `gorm:"primary_key"`
	ResourceID             string
	StudentID              string
	Course                 Course
	SegmentNumber          uint
	StudentSegmentSessions StudentSegmentSessions
	SegmentCategories      SegmentCategories
	Archived               bool
}

// Students Sessions for Segment
type StudentSegmentSessions struct {
	// Maybe use gorm.Model that automatically give ID, CreatedAt,UpdatedAt and DeletedAt fields. ?
	//gorm.Model
	ID                uint `gorm:"primary_key"`
	ResourceID        uint
	StartTime         string
	EndTime           string
	CreatedAt         string
	UpdateAt          string
	DeletedAt         string
	SegmentCategories SegmentCategories
	Comment           string
	Version           uint
	Locations         string
}

// Table for Schools. School can have multiple campuses
type Schools struct {
	ID       uint `gorm:"primary_key"`
	Finnish  string
	English  string
	Campuses []Campuses `gorm:"association_foreignkey:ID;AssociationForeignKey:ID"`
}

// Campus of the School, Campus can have multiple Apartments
type Campuses struct {
	ID         uint `gorm:"primary_key"`
	Finnish    string
	English    string
	Apartments []Apartments `gorm:"association_foreignkey:ID;AssociationForeignKey:ID"`
}

// Table for different Apartments in Campus, Apartment can have multiple Degrees
type Apartments struct {
	ID      uint `gorm:"primary_key"`
	Finnish string
	English string
	Degrees []Degrees `gorm:"association_foreignkey:ID;AssociationForeignKey:ID"`
}

// Degrees in the Apartment.
type Degrees struct {
	ID      uint `gorm:"primary_key"`
	Finnish string
	English string
}
