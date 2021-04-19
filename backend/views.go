package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//POST /v1/LinkedIn/signup
func signup(c *gin.Context) {
	var user Professional
	file, filerError := c.FormFile("photo")
	user.Email = c.PostForm("email")
	user.FirstName = c.PostForm("firstName")
	user.LastName = c.PostForm("lastName")
	user.Password = getMD5Hash(c.PostForm("password"))
	user.PhoneNumber = c.PostForm("phoneNumber")
	//Check if a user with this email already exists
	checkUser, err := dbclient.getProfessional(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal server error"})
		return
	}
	if checkUser.ID > 0 { //A user with this email already exists
		fmt.Println("User already exists")
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "A user with this email already exists"})
		return
	}
	//Create a directory for the user's media
	os.MkdirAll(filepath.Join(mediaDir, user.Email, "profilePhoto"), 0755)
	var photoPath string
	if filerError == nil { //If image given
		//Check if file is an image
		extension := filepath.Ext(file.Filename)
		if !validImgExtension(extension) {
			fmt.Println("Not a valid file")
			c.JSON(http.StatusBadRequest, gin.H{"error": "File is not an image"})
			return
		}
		photoPath = filepath.Join(mediaDir, user.Email, "profilePhoto", file.Filename)
		c.SaveUploadedFile(file, photoPath)
	} else {
		photoPath = ""
	}
	user.Photo = photoPath
	user.save() //Save the new user in the database
	c.JSON(http.StatusCreated, gin.H{"message": "Signed up successfully"})
}

//POST /v1/LinkedIn/signin
func signin(c *gin.Context) {
	type loginInfo struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var userLoginInfo loginInfo
	if err := c.ShouldBindJSON(&userLoginInfo); err == nil {
		professional, err := dbclient.getProfessional(userLoginInfo.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Interval server error"})
		} else {
			if professional.ID == 0 { //A user with this email does not exist
				c.JSON(http.StatusNotFound, gin.H{"error": "Wrong email or password"})
			} else {
				//Get md5 hash of password
				md5pass := getMD5Hash(userLoginInfo.Password)
				if md5pass == professional.Password {
					session := sessions.Default(c)
					session.Set("userEmail", professional.Email)
					session.Set("userID", professional.ID)
					session.Save()
					c.JSON(http.StatusAccepted, gin.H{"message": "Login successfull"})
				} else {
					c.JSON(http.StatusNotFound, gin.H{"error": "Wrong email or password"})
				}
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are necessary"})
	}
}

//POST /v1/LinkedIn/addEducation
func addEducation(c *gin.Context) {
	session := sessions.Default(c)
	professionalID := session.Get("userID")
	if professionalID == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var educationInfo Education
		if err := c.ShouldBindJSON(&educationInfo); err == nil {
			educationInfo.ProfessionalID = professionalID.(int)
			educationInfo.save()
			c.JSON(http.StatusCreated, gin.H{"educationInfo": educationInfo})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Necessary fields not given"})
		}
	}
}

//GET /v1/LinkedIn/logout
func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{"status": "Log out successfull"})
}

// GET /v1/LinkedIn/authenticated
func authenticated(c *gin.Context) {
	session := sessions.Default(c)
	email := session.Get("userEmail")
	professionalID := session.Get("userID")
	if email != nil {
		c.JSON(http.StatusAccepted, gin.H{"status": "Authenticated", "email": email.(string), "id": professionalID.(int)})
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	}
}
