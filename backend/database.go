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

func (driver *DBClient) getAllProfessionalIDs() ([]int, error) {
	var id int
	rows, err := driver.db.Query("SELECT ProfessionalID FROM Professionals")
	if err != nil {
		return nil, err
	}
	var ids []int
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (driver *DBClient) getProfessionalByID(id int) (Professional, error) {
	prof := Professional{}
	rows, err := driver.db.Query("SELECT * FROM Professionals WHERE ProfessionalID=?", id)
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

//Get the articles that a professional liked
func (driver *DBClient) getProfessionalLikedArticles(id int) ([]Article, error) {
	articleInfo := Article{}
	rows, err := driver.db.Query("SELECT a.* FROM Articles a,Article_Likes al WHERE al.ProfessionalID = ? AND a.id = al.ArticleID", id)
	if err != nil {
		return nil, err
	}
	var articles []Article
	for rows.Next() {
		err = rows.Scan(&articleInfo.ID, &articleInfo.UploaderID, &articleInfo.Title, &articleInfo.Content, &articleInfo.AttachedFile, &articleInfo.Created)
		if err != nil {
			return nil, err
		}
		articleInfo.setFileURL()
		articles = append(articles, articleInfo)
	}
	return articles, nil
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

func (driver *DBClient) getProfessionalArticles(professionalID int) ([]Article, error) {
	articleInfo := Article{}
	rows, err := driver.db.Query(`SELECT * FROM Articles WHERE UploaderID=?`, professionalID)
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

func (driver *DBClient) deleteArticle(articleID int) error {
	//First delete the likes
	stmt, err := driver.db.Prepare("DELETE FROM Article_Likes WHERE ArticleID=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(articleID)
	if err != nil {
		return err
	}
	//Then delete the comments
	stmt, err = driver.db.Prepare("DELETE FROM Article_Comments WHERE ArticleID=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(articleID)
	if err != nil {
		return err
	}
	//Finally delete the article
	stmt, err = driver.db.Prepare("DELETE FROM Articles WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(articleID)
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

func (driver *DBClient) getArticles(professionalID int) ([]Article, error) {
	articleInfo := Article{}
	rows, err := driver.db.Query(`SELECT * FROM Articles WHERE 
	(UploaderID IN (SELECT ProfessionalID2 FROM Friendships WHERE ProfessionalID1 = ? AND Status="friends") 
	OR  UploaderID IN (SELECT ProfessionalID1 FROM Friendships WHERE ProfessionalID2=?  AND Status="friends")) 
	OR (id IN(SELECT al.ArticleID FROM Article_Likes al,Friendships f WHERE al.ProfessionalID = f.ProfessionalID2 AND f.ProfessionalID1=? AND f.Status="friends") 
	OR id IN(SELECT al.ArticleID FROM Article_Likes al,Friendships f WHERE al.ProfessionalID = f.ProfessionalID1 AND f.ProfessionalID2=? AND f.Status="friends")) 
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

//Get all the article commments that a Professional made
func (driver *DBClient) getProfessionalArticleComments(professionalID int) ([]ArticleComment, error) {
	comment := ArticleComment{}
	rows, err := driver.db.Query("SELECT * FROM Article_Comments WHERE ProfessionalID=?", professionalID)
	if err != nil {
		return nil, err
	}
	var commentsArray []ArticleComment
	for rows.Next() {
		err = rows.Scan(&comment.ID, &comment.ProfessionalID, &comment.ArticleID, &comment.Comment, &comment.Created)
		if err != nil {
			return nil, err
		}
		commentsArray = append(commentsArray, comment)
	}
	return commentsArray, nil
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

//Function to use for getting the users that a professioanl already has a chat and how many
//unread messages he has from each one
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

//JobAd related functions
func (driver *DBClient) createJobAd(ad *JobAd) error {
	stmt, err := driver.db.Prepare(`INSERT INTO JobAds SET 
		UploaderID=?,
		Title=?,
		Job_Description=?,
		Attached_File=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(ad.UploaderID, ad.Title, ad.JobDescription, ad.AttachedFile)
	if err != nil {
		return err
	}
	return nil
}

//Get the jobads that a professional is interested to
func (driver *DBClient) getProfessionalJobInterest(professionalID int) ([]JobAd, error) {
	ad := JobAd{}
	rows, err := driver.db.Query("SELECT j.* FROM JobAds j,Job_Interest ji WHERE ji.ProfessionalID = ? AND j.id = ji.JobID", professionalID)
	if err != nil {
		return nil, err
	}
	var jobAdsArray []JobAd
	for rows.Next() {
		err = rows.Scan(&ad.ID, &ad.UploaderID, &ad.Title, &ad.JobDescription, &ad.AttachedFile, &ad.Created)
		if err != nil {
			return nil, err
		}
		jobAdsArray = append(jobAdsArray, ad)
	}
	return jobAdsArray, nil
}

//Get the jobAd comments that a professional made
func (driver *DBClient) getProfessionalJobAdComments(professionalID int) ([]JobComment, error) {
	comment := JobComment{}
	rows, err := driver.db.Query("SELECT * FROM Job_Comments WHERE ProfessionalID=?", professionalID)
	if err != nil {
		return nil, err
	}
	var commentsArray []JobComment
	for rows.Next() {
		err = rows.Scan(&comment.ID, &comment.ProfessionalID, &comment.JobID, &comment.Comment, &comment.Created)
		if err != nil {
			return nil, err
		}
		commentsArray = append(commentsArray, comment)
	}
	return commentsArray, nil
}

func (driver *DBClient) getJobAds(professionalID int) ([]JobAd, error) {
	ad := JobAd{}
	rows, err := driver.db.Query(`SELECT * FROM JobAds WHERE 
	(UploaderID IN (SELECT ProfessionalID2 FROM Friendships WHERE ProfessionalID1 = ?) 
	OR  UploaderID IN (SELECT ProfessionalID1 FROM Friendships WHERE ProfessionalID2=?)) ORDER BY Created DESC`, professionalID, professionalID)

	if err != nil {
		return nil, err
	}
	var jobAds []JobAd
	for rows.Next() {
		err = rows.Scan(&ad.ID, &ad.UploaderID, &ad.Title, &ad.JobDescription, &ad.AttachedFile, &ad.Created)
		if err != nil {
			return nil, err
		}
		if len(ad.AttachedFile) > 0 {
			ad.setFileURL() //Change the path of the file to a url
		}
		jobAds = append(jobAds, ad)
	}
	return jobAds, nil
}

func (driver *DBClient) createJobInterest(interest *JobInterest) error {
	stmt, err := driver.db.Prepare(`INSERT INTO Job_Interest SET 
		ProfessionalID=?,
		JobID=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(interest.ProfessionalID, interest.JobID)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) deleteJobInterest(interest JobInterest) error {
	stmt, err := driver.db.Prepare("DELETE FROM Job_Interest WHERE ProfessionalID=? AND JobID=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(interest.ProfessionalID, interest.JobID)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) createJobComment(comment *JobComment) error {
	stmt, err := driver.db.Prepare(`INSERT INTO Job_Comments SET 
		ProfessionalID=?,
		JobID=?,
		Comment=?`,
	)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(comment.ProfessionalID, comment.JobID, comment.Comment)
	if err != nil {
		return err
	}
	return nil
}

func (driver *DBClient) getJobAdUploader(professionalID int) (Professional, error) {
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

//Get the comments of a job ad
func (driver *DBClient) getJobAdComments(ad *JobAd) ([]JobAdCommentResponse, error) {
	comment := JobAdCommentResponse{}
	rows, err := driver.db.Query("SELECT c.id,p.First_Name,p.Last_Name,c.Comment FROM Professionals p,Job_Comments c WHERE c.JobID = ? AND c.ProfessionalID = p.ProfessionalID ORDER BY c.Created", ad.ID)
	if err != nil {
		return nil, err
	}
	var commentsArray []JobAdCommentResponse
	for rows.Next() {
		err = rows.Scan(&comment.ID, &comment.FirstName, &comment.LastName, &comment.Comment)
		if err != nil {
			return commentsArray, err
		}
		commentsArray = append(commentsArray, comment)
	}
	return commentsArray, nil
}

//Get the interest of a jobAd
func (driver *DBClient) getJobAdInterest(ad *JobAd) (int, error) {
	var interestCount int
	rows, err := driver.db.Query("SELECT COUNT(*) FROM Job_Interest WHERE JobID = ?", ad.ID)
	if err != nil {
		return 0, err
	}
	for rows.Next() {
		err = rows.Scan(&interestCount)
		if err != nil {
			return 0, err
		}
	}
	return interestCount, nil
}

func (driver *DBClient) getJobAdFilePath(adID int) (string, error) {
	var path string
	rows, err := driver.db.Query("SELECT Attached_File FROM JobAds WHERE id=?", adID)
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

//Check if a professional is interested in a  particular job
func (driver *DBClient) professionalInterestedForJob(professionalID int, ad JobAd) (bool, error) {
	var interest JobInterest
	rows, err := driver.db.Query("SELECT * FROM Job_Interest WHERE ProfessionalID=? AND JobID=?", professionalID, ad.ID)
	if err != nil {
		return false, err
	}
	for rows.Next() {
		err = rows.Scan(&interest.ID, &interest.ProfessionalID, &interest.JobID)
		if err != nil {
			return false, err
		}
	}
	if interest.ID > 0 {
		return true, nil
	}
	return false, nil
}

func (driver *DBClient) getJobAd(adID int) (JobAd, error) {
	ad := JobAd{}
	rows, err := driver.db.Query("SELECT * FROM JobAds WHERE id=?", adID)
	if err != nil {
		return ad, err
	}
	for rows.Next() {
		err = rows.Scan(&ad.ID, &ad.UploaderID, &ad.Title, &ad.JobDescription, &ad.AttachedFile, &ad.Created)
		if err != nil {
			return ad, err
		}
	}
	return ad, nil
}

func (driver *DBClient) getUserUploadedJobAds(profeesionalID int) ([]JobAd, error) {
	ad := JobAd{}
	rows, err := driver.db.Query("SELECT * FROM JobAds WHERE UploaderID=?", profeesionalID)
	if err != nil {
		return nil, err
	}
	var jobsArray []JobAd
	for rows.Next() {
		err = rows.Scan(&ad.ID, &ad.UploaderID, &ad.Title, &ad.JobDescription, &ad.AttachedFile, &ad.Created)
		if err != nil {
			return nil, err
		}
		if len(ad.AttachedFile) > 0 {
			ad.setFileURL()
		}
		jobsArray = append(jobsArray, ad)
	}
	return jobsArray, nil
}

func (driver *DBClient) getInterestedProfessionals(jobID int) ([]Professional, error) {
	prof := Professional{}
	rows, err := driver.db.Query("SELECT * FROM Professionals WHERE ProfessionalID IN (SELECT ProfessionalID FROM Job_Interest WHERE JobID=?)", jobID)
	if err != nil {
		return nil, err
	}
	var usersArray []Professional
	for rows.Next() {
		err = rows.Scan(&prof.ID, &prof.FirstName, &prof.LastName, &prof.Email, &prof.Password, &prof.PhoneNumber, &prof.Photo)
		if err != nil {
			return nil, err
		}
		if len(prof.Photo) > 0 {
			prof.setPhotoURL()
		}
		usersArray = append(usersArray, prof)
	}
	return usersArray, nil
}

func (driver *DBClient) deleteJobAd(jobAdID int) error {
	//First delete all the jobInterest for this job
	stmt, err := driver.db.Prepare("DELETE FROM Job_Interest WHERE JobID=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(jobAdID)
	if err != nil {
		return err
	}
	//Delete all the comments on this job ad
	stmt, err = driver.db.Prepare("DELETE FROM Job_Comments WHERE JobID=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(jobAdID)
	if err != nil {
		return err
	}
	//Finally delete the job ad
	stmt, err = driver.db.Prepare("DELETE FROM JobAds WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(jobAdID)
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
