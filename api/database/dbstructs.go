package database

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

/*
// dbstructs.go
// Description: Structs used in tiuku
*/

// Table for StudentUsers

type StudentUser struct {
	ID              uint   `gorm:"primary_key"`
	StudentID       string `gorm:"not null"`
	AnonID          string `gorm:"not null"`
	StudentName     string
	StudentSegments string
	StudentEmail    string
	StudentClass    string
}

// Table for FacultyUsers of the API
type FacultyUser struct {
	ID           uint   `gorm:"primary_key"`
	FacultyID    string `gorm:"not null"`
	FacultyName  string
	FacultyEmail string
	Apartment    Apartment
	//School       School
	//FacultySegment FacultySegment
	//FacultySegment string
}

type StudentClass struct {
	School    School
	Campus    Campus
	Apartment Apartment
	Degree    Degree
}

// Table for course
type Course struct {
	ID              uint `gorm:"primary_key"`
	Degree          Degree
	CourseCode      string `gorm:"not null"`
	CourseName      string
	CourseStartDate string
	CourseEndDate   string
	Archived        bool
	Segment         []Segment
}

// Course can have on or more Segments
type Segment struct {
	ID          uint `gorm:"primary_key"`
	CourseID    uint `gorm:"not null"`
	Course      Course
	SegmentName string
	TeacherID   uint
	Scope       uint
	//SegmentCategories     SegmentCategory
	//SegmentCategories     string
	ExpectedAttendance    uint
	SchoolSegmentsSession SchoolSegmentsSession
	Archived              bool
}

// Schools Segment table has data fo students and where to find their Session for the Segement.
type SchoolSegmentsSession struct {
	ID                      uint   `gorm:"primary_key"`
	SegmentID               uint   `gorm:"not null"`
	AnonID                  string `gorm:"not null"`
	StudentSegmentsSessions string
	Privacy                 string //Allowed to see name of the student?
}

// Segment has different Categories for tracking and settings for them.
// All SubCategories belong in to one of the three MainCategoy
// Maybe should be belongs to o one-to-one with two structs?
type SegmentCategory struct {
	ID                 uint `gorm:"primary_key"`
	SegmentID          uint
	Segment            Segment
	MainCategory       uint `gorm:"not null"`
	SubCategory        string
	MandatoryToTrack   bool
	MandatoryToComment bool
	Tickable           bool
	LocationNeeded     bool
	Active             bool
	Archived           bool
}

type MainCategory struct {
	ID        uint `gorm:"primary_key"`
	Shorthand string
	Finnish   string
	English   string
}

// Students Segments list
type StudentSegment struct {
	ID                     uint `gorm:"primary_key"`
	SegmentID              uint
	StudentSegmentSessions string
	Archived               bool
	// Not needed as categories are searched with segment id
	//SegmentCategory        string
	// Can get course from segmentID when needed
	//Course                 Course

}

// Students Sessions for Segment
type StudentSegmentSession struct {
	// Maybe use gorm.Model that automatically give ID, CreatedAt,UpdatedAt and DeletedAt fields. ?
	//gorm.Model
	ID         uint `gorm:"primary_key"`
	ResourceID uint `gorm:"not null"`
	SegmentID  uint `gorm:"not null"`
	Segment    Segment
	StartTime  time.Time
	EndTime    mysql.NullTime
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
	Comment    string
	Version    uint
	Locations  string
	//SegmentCategory string `gorm:"not null"`
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
	Shorthand   string
	Finnish     string
	English     string
	ApartmentID uint
}

// Old and obselete Structs here for storage until sure they're not needed anymore

// Old StudentUser struct
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

// Table of Faculty(in this case Teachers), to save their Segments
// Not used as decided to use single table for faculty and their segments
/*
type FacultySegment struct {
	ID                    uint `gorm:"primary_key"`
	Course                Course
	SegmentNumber         uint
	SchoolSegmentsSession SchoolSegmentsSession
	SegmentCategories     SegmentCategory
	Archived              bool
}
*/

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

/*
func (Degree) TableName() string {
	return "OAMK_Degrees"
}

func (Apartment) TableName() string {
	return "OAMK_Apartments"
}
*/
