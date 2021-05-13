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

func (driver *DBClient) searchProfessional(query string) ([]Professional, error) {
	prof := Professional{}
	sql := "SELECT * FROM Professionals WHERE First_Name LIKE '%" + query + "?%' OR Last_Name LIKE '%" + query + "%' OR Email LIKE '%" + query + "%'"
	rows, err := driver.db.Query(sql)
	if err != nil {
		return nil, err
	}
	var searchResults []Professional
	for rows.Next() {
		err = rows.Scan(&prof.ID, &prof.FirstName, &prof.LastName, &prof.Email, &prof.Password, &prof.PhoneNumber, &prof.Photo)
		if err != nil {
			return nil, err
		}
		prof.setPhotoURL() //Change the path of a photo to a url
		searchResults = append(searchResults, prof)
	}
	return searchResults, nil
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

func (driver *DBClient) getArticle(articleID int) (Article, error) {
	article := Article{}
	rows, err := driver.db.Query("SELECT * FROM Articles WHERE id=?", articleID)
	if err != nil {
		return article, err
	}
	for rows.Next() {
		err = rows.Scan(&article.ID, &article.UploaderID, &article.Title, &article.Content, &article.AttachedFile, &article.Created)
		if err != nil {
			return article, err
		}
	}
	return article, nil
}

//I SINARTISI DAME THA ALLAKSI JE THA PIANI ORISMA TO ID TOU
//PROFESSIONAL POU THELOUME NA PIASOUME TA ARTHRA POU THA EMFANISTOUN
//STO XRONOLOGIO TOU
func (driver *DBClient) getArticles(professionalID int) ([]Article, error) {
	articleInfo := Article{}
	rows, err := driver.db.Query(`SELECT * FROM Articles WHERE 
	(UploaderID IN (SELECT ProfessionalID2 FROM Friendships WHERE ProfessionalID1 = ?) 
	OR  UploaderID IN (SELECT ProfessionalID1 FROM Friendships WHERE ProfessionalID2=?)) 
	OR (id IN(SELECT al.ArticleID FROM Article_Likes al,Friendships f WHERE al.ProfessionalID = f.ProfessionalID2 AND f.ProfessionalID1=?) 
	OR id IN(SELECT al.ArticleID FROM Article_Likes al,Friendships f WHERE al.ProfessionalID = f.ProfessionalID1 AND f.ProfessionalID2=?)) 
	ORDER BY Created DESC;`, professionalID, professionalID, professionalID, professionalID)
	if err != nil {
		return nil, err
	}
	var articlesArray []Article
	for rows.Next() {
		err = rows.Scan(&articleInfo.ID, &articleInfo.UploaderID, &articleInfo.Title, &articleInfo.Content, &articleInfo.AttachedFile, &articleInfo.Created)
		if err != nil {
			return articlesArray, err
		}
		if len(articleInfo.AttachedFile) > 0 { //If there is an attached file
			articleInfo.setFileURL() //Change the file directory to a url
		}
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

//Get the comments of an article
func (driver *DBClient) getArticleComments(article *Article) ([]ArticleCommentResponse, error) {
	comment := ArticleCommentResponse{}
	rows, err := driver.db.Query("SELECT c.id,p.First_Name,p.Last_Name,c.Comment FROM Professionals p,Article_Comments c WHERE c.ArticleID = ? AND c.ProfessionalID = p.ProfessionalID ORDER BY c.Created", article.ID)
	if err != nil {
		return nil, err
	}
	var commentsArray []ArticleCommentResponse
	for rows.Next() {
		err = rows.Scan(&comment.ID, &comment.FirstName, &comment.LastName, &comment.Comment)
		if err != nil {
			return commentsArray, err
		}
		commentsArray = append(commentsArray, comment)
	}
	return commentsArray, nil
}

//Get the likes of an article
func (driver *DBClient) getArticleLikes(article *Article) (int, error) {
	var likes int
	rows, err := driver.db.Query("SELECT COUNT(*) FROM Article_Likes WHERE ArticleID = ?", article.ID)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		err = rows.Scan(&likes)
		if err != nil {
			return 0, err
		}
	}
	return likes, nil
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
		ArticleID=?`,
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

func (driver *DBClient) deleteArticleLike(like ArticleLike) error {
	stmt, err := driver.db.Prepare("DELETE FROM Article_Likes WHERE ProfessionalID=? AND ArticleID=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(like.ProfessionalID, like.ArticleID)
	if err != nil {
		return err
	}
	return nil
}

//Check if a professional liked a particular article
func (driver *DBClient) professionalLikedArticle(professionalID int, article Article) (bool, error) {
	var like ArticleLike
	rows, err := driver.db.Query("SELECT * FROM Article_Likes WHERE ProfessionalID=? AND ArticleID=?", professionalID, article.ID)
	if err != nil {
		return false, err
	}
	for rows.Next() {
		err = rows.Scan(&like.ID, &like.ProfessionalID, &like.ArticleID)
		if err != nil {
			return false, err
		}
	}
	if like.ID > 0 {
		return true, nil
	}
	return false, nil
}

//ArticleComment related functins
func (driver *DBClient) createArticleComment(comment *ArticleComment) error {
	stmt, err := driver.db.Prepare(`INSERT INTO Article_Comments SET 
		ProfessionalID=?,
		ArticleID=?,
		Comment=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(comment.ProfessionalID, comment.ArticleID, comment.Comment)
	if err != nil {
		return err
	}
	return nil
}

//Notifications related functions
func (driver *DBClient) createNotification(n *Notification) error {
	stmt, err := driver.db.Prepare(`INSERT INTO Notifications SET 
		ProfessionalID=?,
		Msg=?,
		Seen=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(n.ProfessionalID, n.Msg, n.Seen)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) getProfessionalNotifications(professionalID int) ([]Notification, error) {
	n := Notification{}
	rows, err := driver.db.Query("SELECT * FROM Notifications WHERE ProfessionalID=? ORDER BY Created DESC", professionalID)
	if err != nil {
		return nil, err
	}
	var notifications []Notification
	for rows.Next() {
		err = rows.Scan(&n.ID, &n.ProfessionalID, &n.Msg, &n.Seen, &n.Created)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, n)
	}
	return notifications, nil
}

func (driver *DBClient) getProfessionalNewNotifications(professionalID int) ([]Notification, error) {
	n := Notification{}
	rows, err := driver.db.Query("SELECT * FROM Notifications WHERE ProfessionalID=? AND Seen=False", professionalID)
	if err != nil {
		return nil, err
	}
	var notifications []Notification
	for rows.Next() {
		err = rows.Scan(&n.ID, &n.ProfessionalID, &n.Msg, &n.Seen, &n.Created)
		if err != nil {
			return nil, err
		}
		notifications = append(notifications, n)
	}
	return notifications, nil
}

func (driver *DBClient) clearProfessionalNotifications(professionalID int) error {
	stmt, err := driver.db.Prepare(`UPDATE Notifications SET
		Seen=True WHERE ProfessionalID=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(professionalID)
	if err != nil {
		return err
	}
	return nil
}

//Friendship related functions
func (driver *DBClient) createFriendRequest(professionalID1, professionalID2 int) error {
	stmt, err := driver.db.Prepare(`INSERT INTO Friendships SET 
		ProfessionalID1=?,
		ProfessionalID2=?,
		Status=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(professionalID1, professionalID2, "pending")
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) deleteFriendRequest(professionalID1, professionalID2 int) error {
	stmt, err := driver.db.Prepare("DELETE FROM Friendships WHERE ProfessionalID1=? AND ProfessionalID2=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(professionalID1, professionalID2)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) createFriendship(professionalID1, professionalID2 int) error {
	stmt, err := driver.db.Prepare(`UPDATE Friendships SET
		Status=? WHERE ProfessionalID1=? AND ProfessionalID2=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec("friends", professionalID1, professionalID2)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) getFriendshipStatus(professionalID1, professionalID2 int) (string, error) {
	var status string
	rows, err := driver.db.Query("SELECT Status FROM Friendships WHERE ProfessionalID1=? AND ProfessionalID2=?", professionalID1, professionalID2)
	if err != nil {
		return status, err
	}
	for rows.Next() {
		err = rows.Scan(&status)
		if err != nil {
			return status, err
		}
	}
	return status, nil
}

//Get the professionals that sent a friend request to professional with id:professionalID
func (driver *DBClient) getProfessionalFriendRequests(professionalID int) ([]Professional, error) {
	prof := Professional{}
	rows, err := driver.db.Query(`SELECT p.ProfessionalID ,p.First_Name,p.Last_Name,p.Email FROM Professionals p,Friendships f 
	WHERE p.ProfessionalID = f.ProfessionalID1 AND f.ProfessionalID2=? AND f.Status="pending"`, professionalID)
	if err != nil {
		return nil, err
	}
	var professionals []Professional
	for rows.Next() {
		err = rows.Scan(&prof.ID, &prof.FirstName, &prof.LastName, &prof.Email)
		if err != nil {
			return nil, err
		}
		professionals = append(professionals, prof)
	}
	return professionals, nil
}

func (driver *DBClient) getProfessionalFriends(professionalID int) ([]Professional, error) {
	prof := Professional{}
	rows, err := driver.db.Query(`SELECT p.ProfessionalID ,p.First_Name,p.Last_Name,p.Email,p.Photo
	FROM Professionals p,Friendships f 
	WHERE p.ProfessionalID = f.ProfessionalID1 AND f.ProfessionalID2=? AND f.Status="friends"`, professionalID)
	if err != nil {
		return nil, err
	}
	var professionals []Professional
	for rows.Next() {
		err = rows.Scan(&prof.ID, &prof.FirstName, &prof.LastName, &prof.Email, &prof.Photo)
		if err != nil {
			return nil, err
		}
		prof.setPhotoURL() //Change directory of photo to a url
		professionals = append(professionals, prof)
	}
	rows, err = driver.db.Query(`SELECT p.ProfessionalID ,p.First_Name,p.Last_Name,p.Email,p.Photo
	FROM Professionals p,Friendships f 
	WHERE p.ProfessionalID = f.ProfessionalID2 AND f.ProfessionalID1=? AND f.Status="friends"`, professionalID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&prof.ID, &prof.FirstName, &prof.LastName, &prof.Email, &prof.Photo)
		if err != nil {
			return nil, err
		}
		prof.setPhotoURL() //Change directory of photo to a url
		professionals = append(professionals, prof)
	}
	return professionals, nil
}

//Message related functions
func (driver *DBClient) createMessage(message *Message) error {
	stmt, err := driver.db.Prepare(`INSERT INTO Messages SET 
		Sender=?,
		Receiver=?,
		Msg=?,
		Seen=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(message.Sender, message.Receiver, message.Msg, false)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) getChat(professionalID1, professionalID2 int) ([]ChatMessage, error) {
	msg := ChatMessage{}
	rows, err := driver.db.Query(`SELECT p.First_Name,p.Last_Name,p.Photo,m.Msg FROM Professionals p,Messages m 
	WHERE p.ProfessionalID=m.Sender 
	AND ((m.Sender=? AND m.Receiver=?) OR (m.Sender=? AND m.Receiver=?)) ORDER BY m.Created`, professionalID1, professionalID2, professionalID2, professionalID1)
	if err != nil {
		return nil, err
	}
	var chatMessages []ChatMessage
	for rows.Next() {
		err = rows.Scan(&msg.FirstName, &msg.LastName, &msg.Photo, &msg.Msg)
		if err != nil {
			return nil, err
		}
		//Change photo path to a url
		msg.setPhotoURL()
		chatMessages = append(chatMessages, msg)
	}
	return chatMessages, nil
}

//Set unseen messages of a professional to seen
func (driver *DBClient) updateMessagesStatus(professionalID, sender int) error {
	stmt, err := driver.db.Prepare(`UPDATE Messages SET
		Seen=? WHERE Receiver=? AND Sender=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(true, professionalID, sender)
	if err != nil {
		return err
	}
	return nil
}

//Get the number of unread dialogs
func (driver *DBClient) getUnreadDialogs(professionalID int) (int, error) {
	count := 0
	rows, err := driver.db.Query(`SELECT COUNT(*) FROM Messages WHERE Receiver=? AND Seen=0 GROUP BY Sender`, professionalID)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		count = count + 1
	}
	return count, nil
}

//Functions to use for getting the user that a professioanl already has a chat and how many
//unread messages he has from each one
//SELECT DISTINCT p.ProfessionalID, p.First_Name,p.Last_Name,p.Photo FROM Professionals p,Messages m WHERE m.Sender=p.ProfessionalID AND m.receiver=?;
//Get the dialogs(users that he has messaged) of the user
func (driver *DBClient) getProfessionalDialogs(professionalID int) ([]ChatDialog, error) {
	dialog := ChatDialog{}
	//Get the users that the professional has messaged
	rows, err := driver.db.Query(`SELECT DISTINCT p.ProfessionalID, p.First_Name,p.Last_Name,p.Photo FROM Professionals p,Messages m WHERE m.Sender=p.ProfessionalID AND m.receiver=?`, professionalID)
	if err != nil {
		return nil, err
	}
	var chatDialogs []ChatDialog
	for rows.Next() {
		err = rows.Scan(&dialog.ProfessionalID, &dialog.FirstName, &dialog.LastName, &dialog.ProfessionalPhoto)
		if err != nil {
			return nil, err
		}
		//Change photo path to a url
		dialog.setPhotoURL()
		chatDialogs = append(chatDialogs, dialog)
	}
	//Get number of unread messages from each user
	for i := 0; i < len(chatDialogs); i++ {
		rows, err := driver.db.Query(`SELECT COUNT(*) FROM Messages m WHERE m.Sender = ? AND m.Receiver=? AND m.Seen = 0`, chatDialogs[i].ProfessionalID, professionalID)
		if err != nil {
			return nil, err
		}
		var count int
		for rows.Next() {
			err = rows.Scan(&count)
			if err != nil {
				return nil, err
			}
			//Change photo path to a url
			chatDialogs[i].UnreadMessages = count
		}
	}
	return chatDialogs, nil
}

//SELECT COUNT(*) FROM Messages m WHERE m.Sender = ? AND m.Receiver=? AND m.Seen = 0;
//Order them by count?

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
