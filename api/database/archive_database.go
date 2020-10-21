package database

import "log"

// Toggle Archive status of course, it's segments and categories, true to archive, false to un-archive
// W1P , Archiving Course and it's Segments + Categories works already.
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
	log.Printf("Course to Archive %v", courseToArchive.ID)
	var tempSegment []Segment
	result := Tiukudb.Table(SegmentTableToEdit).Where("course_id = ?", courseToArchive.ID).Find(&tempSegment)
	if result == nil {
		log.Println(result)
		errorFlag = true
	} else {
		log.Printf("ArchiveCourse phase #2")
		result2, _ := result.Rows()
		log.Printf("Result2 is %v", result2)
		var tempSegment2 Segment
		// Going through Segments
		for result2.Next() {
			if err3 := result.ScanRows(result2, &tempSegment2); err3 != nil {
				log.Printf("Error: Problem with Segment data <database/archive_database.go->ArchiveCourse> %v \n", err3)
				errorFlag = true
			} else {
				log.Println(tempSegment2)
				log.Printf("ArchiveCourse phase #3")
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
		log.Printf("What is this shit.")
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
						returnArchive.segmentID = tempSegment.ID
						returnArchive.SegmentName = tempSegment.SegmentName
						returnArchive.TeacherID = tempSegment.TeacherID
						returnArchive.Scope = tempSegment.Scope
						returnArchive.ExpectedAttendance = tempSegment.ExpectedAttendance
						// Set Course Data
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

	//tempStudent = GetStudentUserWithStudentID(user)
	tempArchive.AnonID = tempStudent.AnonID
	tableToCopyFrom := tempStudent.AnonID + "_sessions"
	tableToCopyTo := tempStudent.AnonID + "_sessions_archived"
	resFrom := Tiukudb.Table(tableToCopyFrom).Where("segment_id = ?", segmentId).Find(&tempStudentSessions)
	if resFrom.Error != nil {
		errorFlag = true
		log.Printf("Error reading user table in <database/maintenance_database.go->ArchiveToSchoolTable> %v", resFrom.Error)
	} else {
		var tempCategories []SegmentCategory
		Tiukudb.Table(CategoriesTableToEdit).Where("segment_id = ?", segmentId).Find(&tempCategories)

		resTo, _ := resFrom.Rows()
		var tempRes StudentSegmentSession
		for resTo.Next() {
			if err2 := resFrom.ScanRows(resTo, &tempRes); err2 != nil {
				log.Printf("Error in <database/maintenance_database.go->ArchiveToSchoolTable> %v", err2)
				errorFlag = true
			} else {

				// Copy to Student Users Archive Table if has one.
				// T35T ? T0D0 needs check if the Archive already has it.
				if DebugMode {
					log.Printf("Starting to Copy Session into User Archive... ")
				}
				if Tiukudb.HasTable(tableToCopyTo) {
					if DebugMode {
						log.Printf("User has Sessions_Archive table..")
					}
					/*
						var tempToTable StudentSegmentSession
						result := Tiukudb.Table(tableToCopyTo).Where(&StudentSegmentSession{ResourceID: tempRes.ResourceID,
							SegmentID: tempRes.SegmentID, Category: tempRes.Category, StartTime: tempRes.StartTime,
							EndTime: tempRes.EndTime, Created: tempRes.Created, Updated: tempRes.Updated,
							Deleted: tempRes.Deleted, Comment: tempRes.Comment, Version: tempRes.Version,
							Locations: tempRes.Locations, Privacy: tempRes.Privacy}).Find(&tempToTable).RowsAffected
						if result == 0 {
							if DebugMode {
								log.Printf("Copying Session to User Sessions_Archive table...")
							}
					*/
					tempRes.ID = 0
					Tiukudb.Table(tableToCopyTo).Create(&tempRes)
					/*
						} else {
							log.Printf("Error: Session already in Users Sessions_Archive table...")
						}
					*/
				}
			}

			// Copy to Schools ArchiveTable
			if DebugMode {
				log.Printf("Starting transfer to Schools Archive table...")
			}
			tempArchive.ID = 0
			tempArchive.StartTime = tempRes.StartTime
			tempArchive.EndTime = tempRes.EndTime
			tempArchive.Created = tempRes.Created
			tempArchive.Updated = tempRes.Updated
			for i := range tempCategories {
				if tempCategories[i].ID == tempRes.Category {
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
				if tempRes.Comment == "No Comment." {
					tempArchive.Comment = StringForEmpy
				} else {
					tempArchive.Comment = "Commented."
				}
			}
			if DebugMode {
				log.Printf("Currently tempArchive is... %v", tempArchive)
			}
			// Save the data
			// T35T if the Archive already has it
			//var tempArchiveCheck ArchivedSessionsTable
			//resArch := Tiukudb.Table(ArchiveTableToEdit).Where(&tempArchive).Find(&tempArchiveCheck).RowsAffected
			//if resArch == 0 {
			if err0r := Tiukudb.Table(ArchiveTableToEdit).Create(&tempArchive).Error; err0r != nil {
				log.Printf("Error in ArchiveToSchoolTable line 258: %v", err0r)
			}
			//} else {
			//	log.Printf("Error: Table %s already has this Session <database/archive_database.go->ArchiveToSchoolTable>. ", ArchiveTableToEdit)
			//}
		}
	}
	return errorFlag
}

// Archive Segment on Student User Segments table
// T35T errorFlag fixed
func ArchiveSegmentOnPersonalTable(tempStudent StudentUser, segId uint) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var errorFlag bool = false
	var tempSegmentTable StudentSegment
	//tempStudent := GetStudentUserWithStudentID(user)
	tableToEdit := tempStudent.AnonID + "_segments"
	if err := Tiukudb.Table(tableToEdit).Where("segment_id = ? ", segId).Find(&tempSegmentTable).Error; err != nil {
		log.Printf("Error: Could not retrieve segments table for student user in <database/archive_database.go->ArchiveSegmentOnPersonalTable %v \n", err)
		errorFlag = true
	} else {
		tempSegmentTable.Archived = true
		if err2 := Tiukudb.Table(tableToEdit).Save(tempSegmentTable).Error; err2 != nil {
			log.Printf("Error: Could not save status to Archived")
			errorFlag = true
		}
	}
	return errorFlag
}

// Sequence to when Course is Archived
// W1P
func ArchiveCourseSequence(courseToArchive Course) bool {
	if Tiukudb == nil {
		ConnectToDB()
	}
	var errorFlag bool = false
	//

	return errorFlag
}
