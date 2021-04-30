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
	feed, err := dbclient.getArticles()
	return feed, err
}

//Check if a professional liked an article
func (prof *Professional) likedArticle(article Article) (bool, error) {
	liked, err := dbclient.professionalLikedArticle(prof.ID, article)
	return liked, err
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
