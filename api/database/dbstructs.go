package database

/*
// dbstructs.go
// Description: Structs used in tiuku
*/

// Table for StudentUsers

type StudentUser struct {
	ID              uint   `gorm:"primary_key"`
	StudentID       string `gorm:"not null"`
	AnonID          string `json:"-"` // `gorm:"not null"`
	StudentName     string
	StudentSegments string `json:"-"`
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
	Degree          uint
	CourseCode      string `gorm:"not null"`
	CourseName      string
	CourseStartDate string
	CourseEndDate   string
	Archived        bool
	Segment         []Segment
	//Degree          Degree
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

// Schools Segment table has data for students and where to find their Session for the Segement.
type SchoolSegmentsSession struct {
	ID                      uint   `gorm:"primary_key"`
	SegmentID               uint   `gorm:"not null"`
	AnonID                  string `gorm:"not null"`
	StudentSegmentsSessions string
	Privacy                 string //Allowed to see name of the student?
}

// Session Report Struct for Segment
type SegmentSessionReport struct {
	StudentID  string
	ResourceID uint
	SegmentID  uint //`gorm:"not null"`
	Category   uint
	StartTime  string //`gorm:"type:datetime" json:"start_time,omitempty"`
	EndTime    string //`gorm:"default=FUGorm"` //`gorm:"type:datetime" json:"end_time,omitempty"`
	Created    string //time.Time
	Updated    string //time.Time
	//Deleted    string //`gorm:"default=FUGorm"` //`default:"NotSet"` //*time.Time
	Comment string
	//Version    uint
	//Locations  string
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
	MandatoryToTrack   bool // Have to see student name, Warning if not
	MandatoryToComment bool // Comment field should not be empty, Warning if not
	Tickable           bool // Faculty only wants to see if tracked or not, time spent not needed
	LocationNeeded     bool // Not used
	Active             bool // Visible to Student Users
	Archived           bool // Archived, Not visible to Student Users and maybe helps with sorting if many categories.
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
	ResourceID uint
	SegmentID  uint //`gorm:"not null"`
	//Segment    Segment
	Category  uint
	StartTime string //`gorm:"type:datetime" json:"start_time,omitempty"`
	EndTime   string //`gorm:"default=FUGorm"` //`gorm:"type:datetime" json:"end_time,omitempty"`
	Created   string //time.Time
	Updated   string //time.Time
	Deleted   string //`gorm:"default=FUGorm"` //`default:"NotSet"` //*time.Time
	Comment   string
	Version   uint
	Locations string
	Privacy   bool
	//SegmentCategory string `gorm:"not null"`
}

//mysql.NullTime
// Table for School. School can have multiple campuses
type School struct {
	ID        uint `gorm:"primary_key"`
	Shorthand string
	Finnish   string
	English   string
	Campuses  []Campus //`gorm:"association_foreignkey:ID;AssociationForeignKey:ID"`
}

// Campus of the School, Campus can have multiple Apartments
type Campus struct {
	ID         uint `gorm:"primary_key"`
	SchoolID   uint
	Shorthand  string
	Finnish    string
	English    string
	Apartments []Apartment //`gorm:"association_foreignkey:ID;AssociationForeignKey:ID"`
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
