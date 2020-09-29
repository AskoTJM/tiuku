package database

// Table for StudentUsers
/*
type StudentUser struct {
	ID              uint `gorm:"primary_key"`
	StudentID       string
	AnonID          string
	StudentName     string
	StudentSegments StudentSegment
	StudentEmail    string
	StudentClass    string
}
*/
// Restructed
type StudentUser struct {
	ID              uint `gorm:"primary_key"`
	StudentID       string
	AnonID          string
	StudentName     string
	StudentSegments string
	StudentEmail    string
	StudentClass    string
}

// Table for FacultyUsers f the API
type FacultyUser struct {
	ID           uint `gorm:"primary_key"`
	FacultyID    string
	FacultyName  string
	FacultyEmail string
	School       School
	Apartment    Apartment
	//FacultySegment FacultySegment
	FacultySegment string
}

type StudentClass struct {
	School    School
	Campus    Campus
	Apartment Apartment
	Degree    Degree
}

// Table of Faculty(in this case Teachers), to save their Segments
type FacultySegment struct {
	ID                    uint `gorm:"primary_key"`
	Course                Course
	SegmentNumber         uint
	SchoolSegmentsSession SchoolSegmentsSession
	SegmentCategories     SegmentCategory
	Archived              bool
}

// Table for course
type Course struct {
	ID              uint `gorm:"primary_key"`
	Schools         School
	CourseCode      string
	CourseName      string
	CourseStartDate string
	CourseEndDate   string
	Archived        bool
	Segment         []Segment
}

// Course can have on or more Segments
type Segment struct {
	ID                    uint `gorm:"primary_key"`
	CourseID              uint
	SegmentName           string
	TeacherID             string
	Scope                 uint
	SegmentCategories     SegmentCategory
	ExpectedAttendance    uint
	SchoolSegmentsSession SchoolSegmentsSession
}

// Schools Segment table has data fo students and where to find their Session for the Segement.
type SchoolSegmentsSession struct {
	ID                      uint `gorm:"primary_key"`
	AnonID                  string
	StudentSegmentsSessions string
	Privacy                 string
}

// Segment has different Categories for tracking and settings for them.
// All SubCategories belong in to one of the three MainCategoy
// Maybe should be belongs to o one-to-one with two structs?
type SegmentCategory struct {
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
/*
type StudentSegment struct {
	ID                     uint `gorm:"primary_key"`
	Course                 Course
	SegmentNumber          uint
	StudentSegmentSessions StudentSegmentSession
	SegmentCategory        SegmentCategory
	Archived               bool
}
*/
// Revamped
type StudentSegment struct {
	ID                     uint `gorm:"primary_key"`
	Course                 Course
	SegmentNumber          uint
	StudentSegmentSessions string
	SegmentCategory        string
	Archived               bool
}

// Students Sessions for Segment
/*
type StudentSegmentSession struct {
	// Maybe use gorm.Model that automatically give ID, CreatedAt,UpdatedAt and DeletedAt fields. ?
	//gorm.Model
	ID              uint `gorm:"primary_key"`
	StartTime       string
	EndTime         string
	CreatedAt       string
	UpdateAt        string
	DeletedAt       string
	SegmentCategory SegmentCategory
	Comment         string
	Version         uint
	Locations       string
}
*/
// Restruct
type StudentSegmentSession struct {
	// Maybe use gorm.Model that automatically give ID, CreatedAt,UpdatedAt and DeletedAt fields. ?
	//gorm.Model
	ID              uint `gorm:"primary_key"`
	StartTime       string
	EndTime         string
	CreatedAt       string
	UpdateAt        string
	DeletedAt       string
	SegmentCategory string
	Comment         string
	Version         uint
	Locations       string
}

// Table for School. School can have multiple campuses
type School struct {
	ID        uint `gorm:"primary_key"`
	Shorthand string
	Finnish   string
	English   string
	Campuses  []Campus `gorm:"association_foreignkey:ID;AssociationForeignKey:ID"`
}

// Campus of the School, Campus can have multiple Apartments
type Campus struct {
	ID         uint `gorm:"primary_key"`
	SchoolID   uint
	Shorthand  string
	Finnish    string
	English    string
	Apartments []Apartment `gorm:"association_foreignkey:ID;AssociationForeignKey:ID"`
}

/*
func (Campus) TableName() string {
	return "OAMK_Campuses"
}
*/
// Table for different Apartment in Campus, Apartment can have multiple Degrees
type Apartment struct {
	ID        uint `gorm:"primary_key"`
	CampusID  uint
	Shorthand string
	Finnish   string
	English   string
	Degrees   []Degree `gorm:"association_foreignkey:ID;AssociationForeignKey:ID"`
}

// Degree in the Apartment.
type Degree struct {
	ID          uint `gorm:"primary_key"`
	ApartmentID uint
	Shorthand   string
	Finnish     string
	English     string
}

/*
func (Degree) TableName() string {
	return "OAMK_Degrees"
}

func (Apartment) TableName() string {
	return "OAMK_Apartments"
}
*/
