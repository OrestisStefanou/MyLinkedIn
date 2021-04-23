package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DBClient stores the database session imformation. Needs to be initialized once
type DBClient struct {
	db *sql.DB
}

func (driver *DBClient) initialize() {
	var err error
	driver.db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", databaseUser, userPassword, databaseName))
	checkErr(err)

	driver.db.SetConnMaxLifetime(time.Minute * 3)
	driver.db.SetMaxOpenConns(10)
	driver.db.SetMaxIdleConns(10)
}

func (driver *DBClient) createProfessional(prof *Professional) error {
	stmt, err := driver.db.Prepare(`INSERT INTO Professionals SET 
		First_Name=?,
		Last_Name=?,
		Email=?,
		Password=?,
		Phone_Number=?,
		Photo=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(prof.FirstName, prof.LastName, prof.Email, prof.Password, prof.PhoneNumber, prof.Photo)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) updateProfessional(prof *Professional) error {
	stmt, err := driver.db.Prepare(`UPDATE Professionals SET
		First_Name=?,
		Last_Name=?,
		Email=?,
		Password=?,
		Phone_Number=?,
		Photo=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(prof.FirstName, prof.LastName, prof.Email, prof.Password, prof.PhoneNumber, prof.Photo)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) getProfessional(email string) (Professional, error) {
	prof := Professional{}
	rows, err := driver.db.Query("SELECT * FROM Professionals WHERE Email=?", email)
	if err != nil {
		return prof, err
	}
	for rows.Next() {
		err = rows.Scan(&prof.ID, &prof.FirstName, &prof.LastName, &prof.Email, &prof.Password, &prof.PhoneNumber, &prof.Photo)
		if err != nil {
			return prof, err
		}
	}
	return prof, nil
}

func (driver *DBClient) createEducation(education *Education) error {
	stmt, err := driver.db.Prepare(`INSERT INTO Education SET 
		ProfessionalID=?,
		Degree_Name=?,
		School_Name =?,
		Start_Date=?,
		Finish_Date =?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(education.ProfessionalID, education.DegreeName, education.SchoolName, education.StartDate, education.FinishDate)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) getEducationID() (int, error) {
	var id int
	rows, err := driver.db.Query("SELECT MAX(id) FROM Education")
	if err != nil {
		return -1, err
	}
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return -1, err
		}
	}
	return id, nil
}

func (driver *DBClient) deleteProfessionalEducation(education Education) error {
	stmt, err := driver.db.Prepare("DELETE FROM Education WHERE ProfessionalID=? AND Degree_Name=? AND School_Name=? AND Start_Date=? AND Finish_Date=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(education.ProfessionalID, education.DegreeName, education.SchoolName, education.StartDate, education.FinishDate)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) getProfessionalEducation(professionalID int) ([]Education, error) {
	educationInfo := Education{}
	rows, err := driver.db.Query("SELECT * FROM Education WHERE ProfessionalID=?", professionalID)
	if err != nil {
		return nil, err
	}
	var educationArray []Education
	for rows.Next() {
		err = rows.Scan(&educationInfo.ID, &educationInfo.ProfessionalID, &educationInfo.DegreeName, &educationInfo.SchoolName, &educationInfo.StartDate, &educationInfo.FinishDate)
		if err != nil {
			return educationArray, err
		}
		educationArray = append(educationArray, educationInfo)
	}
	return educationArray, nil
}

func (driver *DBClient) createExperience(experience *Experience) error {
	stmt, err := driver.db.Prepare(`INSERT INTO Experience SET 
		ProfessionalID=?,
		Employer_Name=?,
		Job_Title =?,
		Start_Date=?,
		Finish_Date =?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(experience.ProfessionalID, experience.EmployerName, experience.JobTitle, experience.StartDate, experience.FinishDate)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) getExperienceID() (int, error) {
	var id int
	rows, err := driver.db.Query("SELECT MAX(id) FROM Experience")
	if err != nil {
		return -1, err
	}
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return -1, err
		}
	}
	return id, nil
}

func (driver *DBClient) getProfessionalExperience(professionalID int) ([]Experience, error) {
	experienceInfo := Experience{}
	rows, err := driver.db.Query("SELECT * FROM Experience WHERE ProfessionalID=?", professionalID)
	if err != nil {
		return nil, err
	}
	var experienceArray []Experience
	for rows.Next() {
		err = rows.Scan(&experienceInfo.ID, &experienceInfo.ProfessionalID, &experienceInfo.EmployerName, &experienceInfo.JobTitle, &experienceInfo.StartDate, &experienceInfo.FinishDate)
		if err != nil {
			return experienceArray, err
		}
		experienceArray = append(experienceArray, experienceInfo)
	}
	return experienceArray, nil
}

func (driver *DBClient) deleteProfessionalExperience(experience Experience) error {
	stmt, err := driver.db.Prepare("DELETE FROM Experience WHERE ProfessionalID=? AND Employer_Name=? AND Job_Title=? AND Start_Date=? AND Finish_Date=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(experience.ProfessionalID, experience.EmployerName, experience.JobTitle, experience.StartDate, experience.FinishDate)
	if err != nil {
		return err
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
