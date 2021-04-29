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

//Professional model related functions
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
	if len(prof.Password) != 0 {
		stmt, err := driver.db.Prepare(`UPDATE Professionals SET
		First_Name=?,
		Last_Name=?,
		Email=?,
		Password=?,
		Phone_Number=?,
		Photo=? WHERE ProfessionalID=?`,
		)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(prof.FirstName, prof.LastName, prof.Email, prof.Password, prof.PhoneNumber, prof.Photo, prof.ID)
		if err != nil {
			return err
		}
		return nil
	}
	stmt, err := driver.db.Prepare(`UPDATE Professionals SET
		First_Name=?,
		Last_Name=?,
		Email=?,
		Phone_Number=?,
		Photo=? WHERE ProfessionalID=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(prof.FirstName, prof.LastName, prof.Email, prof.PhoneNumber, prof.Photo, prof.ID)
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

//Education model related functions
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

//Experience model related functions
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

//Skill model related functions
func (driver *DBClient) createSkill(skill *Skill) error {
	stmt, err := driver.db.Prepare(`INSERT INTO Skills SET 
		ProfessionalID=?,
		Name=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(skill.ProfessionalID, skill.Name)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) getProfessionalSkills(professionalID int) ([]Skill, error) {
	skillInfo := Skill{}
	rows, err := driver.db.Query("SELECT * FROM Skills WHERE ProfessionalID=?", professionalID)
	if err != nil {
		return nil, err
	}
	var skillsArray []Skill
	for rows.Next() {
		err = rows.Scan(&skillInfo.ID, &skillInfo.ProfessionalID, &skillInfo.Name)
		if err != nil {
			return skillsArray, err
		}
		skillsArray = append(skillsArray, skillInfo)
	}
	return skillsArray, nil
}

func (driver *DBClient) deleteProfessionalSkill(skill Skill) error {
	stmt, err := driver.db.Prepare("DELETE FROM Skills WHERE ProfessionalID=? AND Name=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(skill.ProfessionalID, skill.Name)
	if err != nil {
		return err
	}
	return nil
}

//Article model related functions
func (driver *DBClient) createArticle(article *Article) error {
	stmt, err := driver.db.Prepare(`INSERT INTO Articles SET 
		UploaderID=?,
		Title=?,
		Content=?,
		Attached_File=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(article.UploaderID, article.Title, article.Content, article.AttachedFile)
	if err != nil {
		return err
	}
	return nil
}

//I SINARTISI DAME THA ALLAKSI JE THA PIANI ORISMA TO ID TOU
//PROFESSIONAL POU THELOUME NA PIASOUME TA ARTHRA POU THA EMFANISTOUN
//STO XRONOLOGIO TOU
func (driver *DBClient) getArticles() ([]Article, error) {
	articleInfo := Article{}
	rows, err := driver.db.Query("SELECT * FROM Articles ORDER BY Created")
	if err != nil {
		return nil, err
	}
	var articlesArray []Article
	for rows.Next() {
		err = rows.Scan(&articleInfo.ID, &articleInfo.UploaderID, &articleInfo.Title, &articleInfo.Content, &articleInfo.AttachedFile, &articleInfo.Created)
		if err != nil {
			return articlesArray, err
		}
		articleInfo.setFileURL() //Change the file directory to a url
		articlesArray = append(articlesArray, articleInfo)
	}
	return articlesArray, nil
}

func (driver *DBClient) getArticleUploader(professionalID int) (Professional, error) {
	professional := Professional{}
	rows, err := driver.db.Query("SELECT * FROM Professionals WHERE ProfessionalID=?", professionalID)
	if err != nil {
		return professional, err
	}
	for rows.Next() {
		err = rows.Scan(&professional.ID, &professional.FirstName, &professional.LastName, &professional.Email, &professional.Password, &professional.PhoneNumber, &professional.Photo)
		if err != nil {
			return professional, err
		}
	}
	return professional, nil
}

func (driver *DBClient) getArticleFilePath(articleID int) (string, error) {
	var path string
	rows, err := driver.db.Query("SELECT Attached_File FROM Articles WHERE id=?", articleID)
	if err != nil {
		return path, err
	}
	for rows.Next() {
		err = rows.Scan(&path)
		if err != nil {
			return path, err
		}
	}
	return path, nil
}

//ArticleLike related functions
func (driver *DBClient) createArticleLike(like *ArticleLike) error {
	stmt, err := driver.db.Prepare(`INSERT INTO Article_Likes SET 
		ProfessionalID=?,
		ArticleID=?,`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(like.ProfessionalID, like.ArticleID)
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
