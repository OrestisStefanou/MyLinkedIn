package main

import (
	"crypto/md5"
	"encoding/hex"
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
