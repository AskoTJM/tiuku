package database

import "log"

func PopulateSchool() {
	if db == nil {
		ConnectToDB()
	}

	if err := db.Create(&School{
		ID:       0,
		Finnish:  "Oulun Ammattikorkeakoulu",
		English:  "Oulu University of Applied Sciences",
		Campuses: []Campus{},
	}).Error; err != nil {
		log.Panic("Problems populating table of Schools. <go/populate.go->populateSchool>")
	}

	if err := db.Table(schoolShortName + "_Campuses").Create(&Campus{
		ID:         0,
		Finnish:    "Linnanmaan Kampus",
		English:    "Campus Linnanmaa",
		Apartments: []Apartment{},
	}).Error; err != nil {
		log.Panic("Problems populating table of Campuses. <go/populate.go->populateSchool>")
	}

	if err := db.Table(schoolShortName + "_Apartments").Create(&Apartment{
		ID:      0,
		Finnish: "Informaatioteknologia",
		English: "Information Technology",
		Degrees: []Degree{},
	}).Error; err != nil {
		log.Panic("Problems populating table of Apartments. <go/populate.go->populateSchool>")
	}

	if err := db.Table(schoolShortName + "_Degrees").Create(&Degree{
		ID:      0,
		Finnish: "Insinööri (AMK), tieto- ja viestintätekniikka",
		English: "Bachelor of Engineering, Information Technology",
	}).Error; err != nil {
		log.Panic("Problems populating table of Degrees. <go/populate.go->populateSchool>")
	}
	/*
		db.Model(&Schools{
			ID:       0,
			Finnish:  "",
			English:  "",
			Campuses: []Campuses{},
		}).AddForeignKey()
	*/
}

func PopulateStudents() {
	if db == nil {
		ConnectToDB()
	}

	if err := db.Table(schoolShortName + "_StudentUsers").Create(&StudentUser{
		ID:              0,
		StudentID:       "oppi1",
		AnonID:          "Anon1",
		StudentName:     "Oppilas 1",
		StudentSegments: StudentSegment{},
		StudentEmail:    "oppilas1@oppilaitos.fi",
		StudentClass:    "tit1",
	}).Error; err != nil {
		log.Panic("Problems populating table of StudentUsers. <go/populate.go->populateStudents>")
	}

	if err := db.Table(schoolShortName + "_StudentUsers").Create(&StudentUser{
		ID:              0,
		StudentID:       "oppi2",
		AnonID:          "Anon2",
		StudentName:     "Oppilas 2",
		StudentSegments: StudentSegment{},
		StudentEmail:    "oppilas2@oppilaitos.fi",
		StudentClass:    "tit1",
	}).Error; err != nil {
		log.Panic("Problems populating table of StudentUsers. <go/populate.go->populateStudents>")
	}

	if err := db.Table(schoolShortName + "_StudentUsers").Create(&StudentUser{
		ID:              0,
		StudentID:       "oppi3",
		AnonID:          "Anon3",
		StudentName:     "Oppilas 3",
		StudentSegments: StudentSegment{},
		StudentEmail:    "oppilas3@oppilaitos.fi",
		StudentClass:    "tit2",
	}).Error; err != nil {
		log.Panic("Problems populating table of StudentUsers. <go/populate.go->populateStudents>")
	}

	// Creating StudentSegments table for oppi1
	/*
		if err := db.Table("oppi1_StudentSegments").CreateTable(&StudentSegment{
			ID:                     0,
			ResourceID:             "",
			StudentID:              "oppi1",
			Course:                 Course{},
			SegmentNumber:          0,
			StudentSegmentSessions: StudentSegmentSession{},
			SegmentCategory:        SegmentCategory{},
			Archived:               false,
		}).Error; err != nil {
			log.Panic("Problems populating table of table StudentUsers. <go/populate.go->populateStudents>")
		}
	*/
	// Creating new row in StudentSegments
	//newSegment := StudentSegment
	//if err := db.Create(&oppi1_StudentSegments)

}
