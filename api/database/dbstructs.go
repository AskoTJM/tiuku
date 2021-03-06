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
	Apartment    uint
	Active       bool
	Teacher      bool
	Admin        bool
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
	//Segment         []Segment
	//Degree          Degree
}

// Course can have on or more Segments
type Segment struct {
	ID       uint `gorm:"primary_key"`
	CourseID uint `gorm:"not null"`
	//Course      Course
	SegmentName string
	TeacherID   uint
	Scope       uint
	//SegmentCategories     SegmentCategory
	//SegmentCategories     string
	ExpectedAttendance uint
	//SchoolSegmentsSession SchoolSegmentsSession
	Archived bool
}

// Schools Segment table has data for students and where to find their Session for the Segement.
type SchoolSegmentsSession struct {
	ID                      uint   `gorm:"primary_key"`
	SegmentID               uint   `gorm:"not null"`
	AnonID                  string `json:"-"`
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
	ID        uint `gorm:"primary_key"`
	SegmentID uint
	//Segment            Segment
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
	CourseID               uint
	SegmentID              uint
	StudentSegmentSessions string `json:"-"`
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
	ID         uint `gorm:"primary_key" json:"-"`
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
	Active    bool     // Show as possible choice
}

// Campus of the School, Campus can have multiple Apartments
type Campus struct {
	ID         uint `gorm:"primary_key"`
	SchoolID   uint
	Shorthand  string
	Finnish    string
	English    string
	Apartments []Apartment //`gorm:"association_foreignkey:ID;AssociationForeignKey:ID"`
	Active     bool        // Show as possible choice
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
	Active    bool     // Show as possible choice
}

// Degree in the Apartment.
type Degree struct {
	ID          uint `gorm:"primary_key"`
	ApartmentID uint
	Shorthand   string
	Finnish     string
	English     string
	Active      bool // Show as possible choice
}

// Struct to Save Archived Sessions
// T35T
type ArchivedSessionsTable struct {
	ID                 uint `gorm:"primary_key"`
	SchoolID           uint
	CampusID           uint
	ApartmentID        uint
	DegreeID           uint
	CourseID           uint
	CourseCode         string `gorm:"not null"`
	CourseName         string
	CourseStartDate    string
	CourseEndDate      string
	SegmentID          uint
	SegmentName        string
	TeacherID          uint
	Scope              uint
	ExpectedAttendance uint
	MainCategory       uint `gorm:"not null"`
	SubCategoryID      uint // Added later
	SubCategory        string
	MandatoryToTrack   bool
	MandatoryToComment bool
	Tickable           bool
	AnonID             string
	StartTime          string
	EndTime            string
	Created            string
	Updated            string
	Deleted            string // Necessary ? Should these even get here?
	Comment            string // Set only if MandatoryToComment = true, otherwise only "Commented"/"NotCommented"
	Version            uint
	Locations          string // Not in use.
	Privacy            bool   // Necessary?
}
