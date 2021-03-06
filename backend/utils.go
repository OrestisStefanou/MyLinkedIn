package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func validImgExtension(extension string) bool {
	for _, ext := range validImageExtensions {
		if extension == ext {
			return true
		}
	}
	return false
}

func validFileExtension(extension string) bool {
	for _, ext := range validAttachedFileExtensions {
		if extension == ext {
			return true
		}
	}
	return false
}

func getProfessionalFromSession(c *gin.Context) (Professional, error) {
	professional := Professional{}
	//Create professional object from session
	session := sessions.Default(c)
	professionalID := session.Get("userID")
	if professionalID == nil {
		return professional, errors.New("Not authenticated")
	}
	firstName := session.Get("firstName")
	lastName := session.Get("lastName")
	email := session.Get("userEmail")
	phoneNumber := session.Get("userPhone")
	photo := session.Get("userPhoto")
	professional.ID = professionalID.(int)
	professional.FirstName = firstName.(string)
	professional.LastName = lastName.(string)
	professional.Email = email.(string)
	professional.PhoneNumber = phoneNumber.(string)
	professional.Photo = photo.(string)
	return professional, nil

}

//Create or update a professional's session
func setProfessionalSession(c *gin.Context, prof Professional) {
	session := sessions.Default(c)
	session.Set("userID", prof.ID)
	session.Set("firstName", prof.FirstName)
	session.Set("lastName", prof.LastName)
	session.Set("userEmail", prof.Email)
	session.Set("userPhone", prof.PhoneNumber)
	session.Set("userPhoto", prof.Photo)
	session.Save()
}

//ArticleCommentResponse json struct
type ArticleCommentResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Comment   string `json:"comment"`
}

//ChatMessage json struct
type ChatMessage struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Photo     string `json:"photo"`
	Msg       string `json:"msg"`
}

//Create profile photo url
func (c *ChatMessage) setPhotoURL() {
	photoURL := mediaURL + c.Photo
	c.Photo = photoURL
}

//ChatDialog json struct
type ChatDialog struct {
	ProfessionalID    int    `json:"professionalID"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	ProfessionalPhoto string `json:"professionalPhoto"`
	UnreadMessages    int    `json:"unreadMessages"`
}

//Create profile photo url
func (c *ChatDialog) setPhotoURL() {
	photoURL := mediaURL + c.ProfessionalPhoto
	c.ProfessionalPhoto = photoURL
}

//JobAdCommentResponse json struct
type JobAdCommentResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Comment   string `json:"comment"`
}

//UserResponse json struct
type UserResponse struct {
	User    Professional `json:"userInfo"`
	Checked bool         `json:"checked"`
}

//UserDetailInfo json struct
type UserDetailInfo struct {
	XMLName                xml.Name         `xml:"userDetailInfo"`
	UserInfo               Professional     `json:"userInfo"`
	EducationInfo          []Education      `json:"education"`
	ExperienceInfo         []Experience     `json:"experience" `
	SkillsInfo             []Skill          `json:"skills" `
	ArticlesInfo           []Article        `json:"articles" `
	ArticleLikes           []Article        `json:"articleLikes" `
	ArticleComments        []ArticleComment `json:"articleComments" `
	ConnectedProfessionals []Professional   `json:"connectedProfessionals" `
	JobAds                 []JobAd          `json:"jobAds" `
	JobInterests           []JobAd          `json:"jobInterests"`
	JobComments            []JobComment     `json:"jobComments"`
}
