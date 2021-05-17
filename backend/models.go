package main

import (
	"path/filepath"
)

//Professional json struct
type Professional struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber"`
	Photo       string `json:"photo"`
}

//Save method for Professional
func (prof *Professional) save() error {
	err := dbclient.createProfessional(prof)
	return err
}

//Update method for Professional
func (prof *Professional) update() error {
	err := dbclient.updateProfessional(prof)
	return err
}

//Create profile photo url
func (prof *Professional) setPhotoURL() {
	photoURL := mediaURL + prof.Photo
	prof.Photo = photoURL
}

//Add education info to a professional
func (prof *Professional) addEducation(educationInfo Education) error {
	educationInfo.ProfessionalID = prof.ID
	err := dbclient.createEducation(&educationInfo)
	return err
}

//Get the education of a professional
func (prof *Professional) getEducation() ([]Education, error) {
	education, err := dbclient.getProfessionalEducation(prof.ID)
	return education, err
}

//Remove education of a professional
func (prof *Professional) removeEducation(educationInfo Education) error {
	educationInfo.ProfessionalID = prof.ID
	err := dbclient.deleteProfessionalEducation(educationInfo)
	return err
}

//Add experience info to a professional
func (prof *Professional) addExperience(experienceInfo Experience) error {
	experienceInfo.ProfessionalID = prof.ID
	err := dbclient.createExperience(&experienceInfo)
	return err
}

//Get the experience of a Professional
func (prof *Professional) getExperience() ([]Experience, error) {
	experience, err := dbclient.getProfessionalExperience(prof.ID)
	return experience, err
}

//Remove experience of a Professional
func (prof *Professional) removeExperience(experienceInfo Experience) error {
	experienceInfo.ProfessionalID = prof.ID
	err := dbclient.deleteProfessionalExperience(experienceInfo)
	return err
}

//Add skill info to a professional
func (prof *Professional) addSkill(skillInfo Skill) error {
	skillInfo.ProfessionalID = prof.ID
	err := dbclient.createSkill(&skillInfo)
	return err
}

//Get the skills of a Professional
func (prof *Professional) getSkills() ([]Skill, error) {
	skills, err := dbclient.getProfessionalSkills(prof.ID)
	return skills, err
}

//Remove skill of a Professional
func (prof *Professional) removeSkill(skillInfo Skill) error {
	skillInfo.ProfessionalID = prof.ID
	err := dbclient.deleteProfessionalSkill(skillInfo)
	return err
}

//Get the feed of a professional
func (prof *Professional) getFeed() ([]Article, error) {
	feed, err := dbclient.getArticles(prof.ID)
	return feed, err
}

//Get the articles that the professional uploaded
func (prof *Professional) getMyArticles() ([]Article, error) {
	articles, err := dbclient.getProfessionalArticles(prof.ID)
	return articles, err
}

//Get the articles that the professional liked
func (prof *Professional) getLikedArticles() ([]Article, error) {
	likedArticles, err := dbclient.getProfessionalLikedArticles(prof.ID)
	return likedArticles, err
}

//Get the jobs that the professional is interested
func (prof *Professional) getInterestedJobs() ([]JobAd, error) {
	jobs, err := dbclient.getProfessionalJobInterest(prof.ID)
	return jobs, err
}

//Get the JobAd comments that the professional made
func (prof *Professional) getMyJobAdComments() ([]JobComment, error) {
	comments, err := dbclient.getProfessionalJobAdComments(prof.ID)
	return comments, err
}

//Check if a professional liked an article
func (prof *Professional) likedArticle(article Article) (bool, error) {
	liked, err := dbclient.professionalLikedArticle(prof.ID, article)
	return liked, err
}

//Get the notifications of a Professional
func (prof *Professional) getNotifications() ([]Notification, error) {
	notifications, err := dbclient.getProfessionalNotifications(prof.ID)
	return notifications, err
}

//Get the new notifications of a Professional
func (prof *Professional) getNewNotifications() ([]Notification, error) {
	notifications, err := dbclient.getProfessionalNewNotifications(prof.ID)
	return notifications, err
}

//Clear the notifications of a Professional
func (prof *Professional) clearNotifications() error {
	err := dbclient.clearProfessionalNotifications(prof.ID)
	return err
}

//Get the friends of a Professional
func (prof *Professional) getFriends() ([]Professional, error) {
	friends, err := dbclient.getProfessionalFriends(prof.ID)
	return friends, err
}

//Get unread dialogs of a Professional
func (prof *Professional) getUnreadDialogs() (int, error) {
	count, err := dbclient.getUnreadDialogs(prof.ID)
	return count, err
}

//Get the chat dialogs of a Professional
func (prof *Professional) getChatDialogs() ([]ChatDialog, error) {
	dialogs, err := dbclient.getProfessionalDialogs(prof.ID)
	return dialogs, err
}

//Get the job ads related to the professional
func (prof *Professional) getAds() ([]JobAd, error) {
	ads, err := dbclient.getJobAds(prof.ID)
	return ads, err
}

//Check if a professional is interested in a job
func (prof *Professional) isInterestedAtJob(ad JobAd) (bool, error) {
	interested, err := dbclient.professionalInterestedForJob(prof.ID, ad)
	return interested, err
}

//Get the job ads that the user uploaded
func (prof *Professional) getMyJobAds() ([]JobAd, error) {
	jobAds, err := dbclient.getUserUploadedJobAds(prof.ID)
	return jobAds, err
}

//Education json struct
type Education struct {
	ID             int    `json:"id"`
	ProfessionalID int    `json:"professionalId"`
	DegreeName     string `json:"degreeName" binding:"required"`
	SchoolName     string `json:"schoolName" binding:"required"`
	StartDate      string `json:"startDate" binding:"required"`
	FinishDate     string `json:"finishDate"`
}

//Save method for Education
func (education *Education) save() error {
	err := dbclient.createEducation(education)
	return err
}

//Method to set id of Education object
func (education *Education) setID() error {
	id, err := dbclient.getEducationID()
	education.ID = id + 1
	return err
}

//Experience json struct
type Experience struct {
	ID             int    `json:"id"`
	ProfessionalID int    `json:"professionalId"`
	EmployerName   string `json:"employerName" binding:"required"`
	JobTitle       string `json:"jobTitle" binding:"required"`
	StartDate      string `json:"startDate" binding:"required"`
	FinishDate     string `json:"finishDate"`
}

//Save method for experience
func (experience *Experience) save() error {
	err := dbclient.createExperience(experience)
	return err
}

//Method to set id of Experience object
func (experience *Experience) setID() error {
	id, err := dbclient.getExperienceID()
	experience.ID = id + 1
	return err
}

//Skill json struct
type Skill struct {
	ID             int    `json:"id"`
	ProfessionalID int    `json:"professionalId"`
	Name           string `json:"name" binding:"required"`
}

//Save method for skill
func (skill *Skill) save() error {
	err := dbclient.createSkill(skill)
	return err
}

//Article json struct
type Article struct {
	ID           int     `json:"id"`
	UploaderID   int     `json:"uploaderId"`
	Title        string  `json:"title" binding:"required"`
	Content      string  `json:"content" binding:"required"`
	AttachedFile string  `json:"file"`
	Created      []uint8 `json:"created"`
}

//Save method for article
func (article *Article) save() error {
	err := dbclient.createArticle(article)
	return err
}

//Create a url for the attached file of the article
func (article *Article) setFileURL() {
	fileURL := mediaURL + article.AttachedFile
	article.AttachedFile = fileURL
}

//Get the comments of an article
func (article *Article) getComments() ([]ArticleCommentResponse, error) {
	comments, err := dbclient.getArticleComments(article)
	return comments, err
}

//Get the likes of an article
func (article *Article) getLikes() (int, error) {
	likes, err := dbclient.getArticleLikes(article)
	return likes, err
}

//Get the info of the professional who uploaded the video
func (article *Article) getUploader() (Professional, error) {
	professional, err := dbclient.getArticleUploader(article.UploaderID)
	professional.Password = ""
	return professional, err
}

//Check if the attached file of the article is an image
func (article *Article) fileIsImage() (bool, error) {
	fileName, err := dbclient.getArticleFilePath(article.ID)
	if err != nil {
		return false, err
	}
	extension := filepath.Ext(fileName)
	if validImgExtension(extension) {
		return true, nil
	}
	return false, nil
}

//Add a like to the article
func (article *Article) addLike(like ArticleLike) error {
	like.ArticleID = article.ID
	err := dbclient.createArticleLike(&like)
	return err
}

//Remove a like from the article
func (article *Article) removeLike(like ArticleLike) error {
	like.ArticleID = article.ID
	err := dbclient.deleteArticleLike(like)
	return err
}

//Add a comment to the article
func (article *Article) addComment(comment ArticleComment) error {
	comment.ArticleID = article.ID
	err := dbclient.createArticleComment(&comment)
	return err
}

//ArticleLike json struct
type ArticleLike struct {
	ID             int `json:"id"`
	ProfessionalID int `json:"professionalId"`
	ArticleID      int `json:"articleId" binding:"required"`
}

//ArticleLike save method
func (like *ArticleLike) save() error {
	err := dbclient.createArticleLike(like)
	return err
}

//ArticleComment json struct
type ArticleComment struct {
	ID             int     `json:"id"`
	ProfessionalID int     `json:"professionalId"`
	ArticleID      int     `json:"articleId" binding:"required"`
	Comment        string  `json:"comment" binding:"required"`
	Created        []uint8 `json:"created"`
}

//ArticleComment save method
func (comment *ArticleComment) save() error {
	err := dbclient.createArticleComment(comment)
	return err
}

//Notification json struct
type Notification struct {
	ID             int     `json:"id"`
	ProfessionalID int     `json:"professionalId"`
	Msg            string  `json:"msg" binding:"required"`
	Seen           bool    `json:"seen"`
	Created        []uint8 `json:"created"`
}

//Notification save method
func (n *Notification) save() error {
	err := dbclient.createNotification(n)
	return err
}

//Friendship json struct
type Friendship struct {
	ID              int    `json:"id"`
	ProfessionalID1 int    `json:"professionalId1"`
	ProfessionalID2 int    `json:"professionalId2"`
	Status          string `json:"status" `
}

//Message json struct
type Message struct {
	ID       int     `json:"id"`
	Sender   int     `json:"sender"`
	Receiver int     `json:"receiver"`
	Msg      string  `json:"msg" binding:"required" `
	Seen     bool    `json:"seen"`
	Created  []uint8 `json:"created"`
}

//Message save method
func (m *Message) save() error {
	err := dbclient.createMessage(m)
	return err
}

//JobAd json struct
type JobAd struct {
	ID             int     `json:"id"`
	UploaderID     int     `json:"uploaderId"`
	Title          string  `json:"title" binding:"required"`
	JobDescription string  `json:"jobDescription" binding:"required"`
	AttachedFile   string  `json:"file"`
	Created        []uint8 `json:"created"`
}

//JobAd save method
func (ad *JobAd) save() error {
	err := dbclient.createJobAd(ad)
	return err
}

//Create a url for the attached file of the job ad
func (ad *JobAd) setFileURL() {
	fileURL := mediaURL + ad.AttachedFile
	ad.AttachedFile = fileURL
}

//Get the comments of a jobAd
func (ad *JobAd) getComments() ([]JobAdCommentResponse, error) {
	comments, err := dbclient.getJobAdComments(ad)
	return comments, err
}

//Get the interest(how many are interested) of a jobAd
func (ad *JobAd) getInterest() (int, error) {
	likes, err := dbclient.getJobAdInterest(ad)
	return likes, err
}

//Get the info of the professional who uploaded the ad
func (ad *JobAd) getUploader() (Professional, error) {
	professional, err := dbclient.getJobAdUploader(ad.UploaderID)
	professional.Password = ""
	return professional, err
}

//Check if the attached file of the article is an image
func (ad *JobAd) fileIsImage() (bool, error) {
	fileName, err := dbclient.getJobAdFilePath(ad.ID)
	if err != nil {
		return false, err
	}
	extension := filepath.Ext(fileName)
	if validImgExtension(extension) {
		return true, nil
	}
	return false, nil
}

func (ad *JobAd) addInterest(interest JobInterest) error {
	interest.JobID = ad.ID
	err := dbclient.createJobInterest(&interest)
	return err
}

func (ad *JobAd) removeInterest(interest JobInterest) error {
	interest.JobID = ad.ID
	err := dbclient.deleteJobInterest(interest)
	return err
}

func (ad *JobAd) addComment(comment JobComment) error {
	comment.JobID = ad.ID
	err := dbclient.createJobComment(&comment)
	return err
}

func (ad *JobAd) getInterestedProfessionals() ([]Professional, error) {
	interested, err := dbclient.getInterestedProfessionals(ad.ID)
	return interested, err
}

//JobInterest json struct
type JobInterest struct {
	ID             int `json:"id"`
	ProfessionalID int `json:"professionalId"`
	JobID          int `json:"jobId" binding:"required"`
}

//ArticleLike save method
func (interest *JobInterest) save() error {
	err := dbclient.createJobInterest(interest)
	return err
}

//JobComment json struct
type JobComment struct {
	ID             int     `json:"id"`
	ProfessionalID int     `json:"professionalId"`
	JobID          int     `json:"jobId" binding:"required"`
	Comment        string  `json:"comment" binding:"required"`
	Created        []uint8 `json:"created"`
}

//ArticleComment save method
func (comment *JobComment) save() error {
	err := dbclient.createJobComment(comment)
	return err
}
