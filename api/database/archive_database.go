package database

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

// Toggle Archive status of course, it's segments and categories, true to archive, false to un-archive
// W0rks
func ArchiveCourse(courseID uint) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var errorFlag bool = false
	// Archiving the Course

	// archive as input removed as implementing un-Archiving all the stuff sounds like hell.
	var archive bool = true
	//courseToArchive.Archived = archive
	//Tiukudb.Save(&courseToArchive)
	courseToArchive := GetCourseTableById(courseID)
	// Set Courses Segment to Archived

	var tempSegment []Segment
	result := Tiukudb.Table(SegmentTableToEdit).Where("course_id = ?", courseToArchive.ID).Find(&tempSegment)
	if result == nil {
		log.Println(result)
		errorFlag = true
	} else {
		result2, _ := result.Rows()
		var tempSegment2 Segment
		// Going through Segments
		for result2.Next() {
			if err3 := result.ScanRows(result2, &tempSegment2); err3 != nil {
				log.Printf("Error: Problem with Segment data <database/archive_database.go->ArchiveCourse> %v \n", err3)
				errorFlag = true
			} else {
				// Get shared data for Archiving.
				var tempArchive ArchivedSessionsTable
				tempArchive, _ = ArchivedSessionsTemplate(tempSegment2.ID)
				// Get Users to that Segment
				//var tempStudents []StudentUser
				tempStudents := GetStudentsJoinedOnSegment(tempSegment2.ID)
				// Loop to Archive users data
				for i := range tempStudents {
					errArch := ArchiveToSchoolTable(tempStudents[i], tempSegment2.ID, tempArchive)
					if errArch {
						log.Printf("Error. Could not Archive Student users data. <database/archive_database.go->ArchiveCourse")
						errorFlag = true
					} else {
						//log.Printf("Starting ArchiveSegmentOnPersonalTable...")
						errArch2 := ArchiveSegmentOnPersonalTable(tempStudents[i], tempSegment2.ID)
						if errArch2 {
							log.Printf("Error. Could not Archive Student users data. <database/archive_database.go->ArchiveCourse")
							errorFlag = true
						}
					}
				}

				// Archiving Segment
				tempSegment2.Archived = archive
				Tiukudb.Save(&tempSegment2)
				// Change Categories for Segment to Archived
				var tempCat []SegmentCategory
				resultSeg := Tiukudb.Table(CategoriesTableToEdit).Where("segment_id = ?", tempSegment2.ID).Find(&tempCat)
				if resultSeg == nil {
					log.Printf("Error: Problem with retrieving Categories for segments. <database/archive_database.go->ArchiveCourse> %v \n", resultSeg)
					errorFlag = true
				}

				resultSeg2, _ := resultSeg.Rows()
				var tempCat2 SegmentCategory
				for resultSeg2.Next() {
					if err4 := result.ScanRows(resultSeg2, &tempCat2); err4 != nil {
						log.Printf("Error: Problem with Categories transfering. <database/archive_databse.go->ArchiveCourse> %v \n", err4)
						errorFlag = true
					}
					// Archiving Segments Categories
					tempCat2.Archived = archive
					Tiukudb.Save(&tempCat2)
				}

			}
		}
		//log.Printf("What is this shit.")
	}
	return errorFlag
}

// Get ArchivedSessions shared data. Returns ArchivedSessionTable and bool true if everything ok
// T35T errorFlag fixed
func ArchivedSessionsTemplate(segmentId uint) (ArchivedSessionsTable, bool) {
	if Tiukudb == nil {
		ConnectToDB()
	}
	// Set variables needed
	var errorFlag bool = false
	var returnArchive ArchivedSessionsTable
	var tempSegment Segment
	var tempCourse Course
	var tempDegree Degree
	var tempApartment Apartment
	var tempCampus Campus

	// Get Segments data
	if segErr := Tiukudb.Table(SegmentTableToEdit).Where("id = ?", segmentId).Find(&tempSegment).Error; segErr != nil {
		log.Printf("Error: Can not retrieve segment data at <database/archive_database.go->ArchivedSessionsTemplate> %v \n", segErr)
		errorFlag = true
	} else {
		// Get Course data
		if courseErr := Tiukudb.Table(CourseTableToEdit).Where("id = ? ", tempSegment.CourseID).Find(&tempCourse).Error; courseErr != nil {
			log.Printf("Error: Can not retrieve course data at <database/archive_database.go->ArchivedSessionsTemplate> %v \n", segErr)
			errorFlag = true
		} else {
			// Get Degree data
			if degErr := Tiukudb.Table(DegreeTableToEdit).Where("id = ?", tempCourse.Degree).Find(&tempDegree).Error; degErr != nil {
				log.Printf("Error: Can not retrieve Degree data at <database/archive_database.go->ArchivedSessionsTemplate> %v \n", segErr)
				errorFlag = true
			} else {
				// Get and set Apartment data
				if apartErr := Tiukudb.Table(ApartmentTableToEdit).Where("id = ?", tempDegree.ApartmentID).Find(&tempApartment).Error; apartErr != nil {
					log.Printf("Error: Can not retrieve Apartment data at <database/archive_database.go->ArchivedSessionsTemplate> %v \n", segErr)
					errorFlag = true
				} else {
					// Get and set Campus data
					if campusErr := Tiukudb.Table(SchoolsTableToEdit).Where("id = ?", tempApartment.CampusID).Find(&tempCampus).Error; campusErr != nil {
						log.Printf("Error: Can not retrieve campus data at <database/archive_database.go->ArchivedSessionsTemplate> %v \n", segErr)
						errorFlag = true
					} else {
						// Set Data
						// Set Segment Data
						returnArchive.SegmentID = tempSegment.ID
						returnArchive.SegmentName = tempSegment.SegmentName
						returnArchive.TeacherID = tempSegment.TeacherID
						returnArchive.Scope = tempSegment.Scope
						returnArchive.ExpectedAttendance = tempSegment.ExpectedAttendance
						// Set Course Data
						returnArchive.CourseID = tempCourse.ID
						returnArchive.CourseCode = tempCourse.CourseCode
						returnArchive.CourseName = tempCourse.CourseName
						returnArchive.CourseStartDate = tempCourse.CourseStartDate
						returnArchive.CourseEndDate = tempCourse.CourseEndDate
						// Set Degree
						returnArchive.DegreeID = tempCourse.Degree
						// Set Apartment
						returnArchive.ApartmentID = tempDegree.ApartmentID
						// Set Campus
						returnArchive.CampusID = tempApartment.CampusID
						// Set School
						returnArchive.SchoolID = tempCampus.SchoolID
					}
				}
			}
		}
	}
	return returnArchive, errorFlag
}

// Archive Sessions to Archive_Tables and input user, segmentsId, populated ArchiveSessions
// T35T errorFlag fixed
func ArchiveToSchoolTable(tempStudent StudentUser, segmentId uint, tempArchive ArchivedSessionsTable) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var errorFlag bool = false
	//var tempStudent StudentUser
	// Move tempArchive creation to other level?
	//tempArchive := ArchivedSessionsTemplate(segmentId)
	var tempStudentSessions []StudentSegmentSession
	var resFrom *gorm.DB
	//tempStudent = GetStudentUserWithStudentID(user)
	tempArchive.AnonID = tempStudent.AnonID
	tableToCopyFrom := tempStudent.AnonID + "_sessions"
	tableToCopyTo := tempStudent.AnonID + "_sessions_archived"
	resFrom = Tiukudb.Table(tableToCopyFrom).Where("segment_id = ?", segmentId).Find(&tempStudentSessions)
	log.Printf("Transfering data from student %v", tempStudent.StudentName)
	var tempCategories []SegmentCategory
	Tiukudb.Table(CategoriesTableToEdit).Where("segment_id = ?", segmentId).Find(&tempCategories)
	resTo, _ := resFrom.Rows()
	var tempRes StudentSegmentSession
	for resTo.Next() {
		if err2 := resFrom.ScanRows(resTo, &tempRes); err2 != nil {
			log.Printf("Error in <database/maintenance_database.go->ArchiveToSchoolTable> %v", err2)
			errorFlag = true
		} else {

			// T0D0 Check if already in table
			tempRes.ID = 0
			Tiukudb.Table(tableToCopyTo).Create(&tempRes)
			// Copy to Schools ArchiveTable
			tempArchive.ID = 0
			tempArchive.StartTime = tempRes.StartTime
			tempArchive.EndTime = tempRes.EndTime
			tempArchive.Created = tempRes.Created
			tempArchive.Updated = tempRes.Updated
			tempArchive.Deleted = tempRes.Deleted
			// None of this category stuff works.
			for i := range tempCategories {
				if tempCategories[i].ID == tempRes.Category {
					log.Printf("Found Category %v", tempCategories[i].SubCategory)
					tempArchive.MainCategory = tempCategories[i].MainCategory
					tempArchive.SubCategory = tempCategories[i].SubCategory
					tempArchive.MandatoryToComment = tempCategories[i].MandatoryToComment
					tempArchive.MandatoryToTrack = tempCategories[i].MandatoryToTrack
					tempArchive.Tickable = tempCategories[i].Tickable
					break
				}
			}

			// Check if requirement for category
			if tempArchive.MandatoryToComment {
				tempArchive.Comment = tempRes.Comment
			} else {
				if tempRes.Comment == "" {
					tempArchive.Comment = StringForEmpy
				} else {
					tempArchive.Comment = "Commented."
				}
			}
			// T0D0 add check if already in table.
			if err := Tiukudb.Table(ArchiveTableToEdit).Create(&tempArchive).Error; err != nil {
				log.Printf("Error in ArchiveToSchoolTable line 259: %v", err)
			} else {
				log.Printf("Added Session to School Archive")
			}
		}
	}
	if !errorFlag {
		// Remove Archived Sessions from Student Uses Sessions table
		if err := Tiukudb.Table(tableToCopyFrom).Where("segment_id = ?", segmentId).Delete(&StudentSegmentSession{}).Error; err != nil {
			log.Printf("Error: Problem deleting Session from student users sessions table. <database/archive_database.go->ArchiveToSchoolTable> %v", err)
			errorFlag = true
		}
		// Remove Student User from ParticipationsList
		if err := Tiukudb.Table(SchoolParticipationList).Where("segment_id = ? AND anon_id = ?", segmentId, tempStudent.AnonID).Delete(&SchoolSegmentsSession{}).Error; err != nil {
			log.Printf("Error: Problem removing student users from Participation list. <database/archive_database.go->ArchiveToSchoolTable> %v", err)
			errorFlag = true
		}

	}
	//slog.Printf("Got to the end of ArchiveToSchoolTable")
	return errorFlag
}

// Archive Segment on Student User Segments table
// T35T errorFlag fixed
func ArchiveSegmentOnPersonalTable(tempStudent StudentUser, segId uint) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	//log.Printf("In ArchiveSegmentOnPersonalTable...")
	var errorFlag bool = false
	var tempSegmentTable StudentSegment
	//tempStudent := GetStudentUserWithStudentID(user)
	tableToEdit := tempStudent.AnonID + "_segments"
	if err := Tiukudb.Table(tableToEdit).Where("segment_id = ? ", segId).Find(&tempSegmentTable).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error: Could not retrieve segments table for student user in <database/archive_database.go->ArchiveSegmentOnPersonalTable %v \n", err)
			errorFlag = true
		}
	} else {
		//log.Printf("Archiving Personal Segment Table...")
		tempSegmentTable.Archived = true
		if err2 := Tiukudb.Table(tableToEdit).Save(&tempSegmentTable).Error; err2 != nil {
			log.Printf("Error: Could not save status to Archived")
			errorFlag = true
		}
	}
	//log.Printf("Returning from ArchiveSegmentOnPersonalTable...")
	return errorFlag
}
