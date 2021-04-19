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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
