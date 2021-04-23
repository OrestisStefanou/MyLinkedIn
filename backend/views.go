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
		//Upload the photo
		photoPath = filepath.Join(mediaDir, user.Email, "profilePhoto", file.Filename)
		c.SaveUploadedFile(file, photoPath)
		//Path to save in the database
		photoPath = filepath.Join(user.Email, "profilePhoto", file.Filename)
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
					session.Set("userID", professional.ID)
					session.Set("firstName", professional.FirstName)
					session.Set("lastName", professional.LastName)
					session.Set("userEmail", professional.Email)
					session.Set("userPhone", professional.PhoneNumber)
					session.Set("userPhoto", professional.Photo)
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

//POST /v1/LinkedIn/updateProfessional
func updateProfessional(c *gin.Context) {
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		file, filerError := c.FormFile("photo")
		newEmail := c.PostForm("email")
		newFirstName := c.PostForm("firstName")
		newLastName := c.PostForm("lastName")
		newPassword := c.PostForm("password")
		newPhoneNumber := c.PostForm("phoneNumber")
		fmt.Println("NEW EMAIL", newEmail)
		fmt.Println("NEW FIRST NAME", newFirstName)
		fmt.Println("NEW LAST NAME", newLastName)
		fmt.Println("NEW PHOTO", file)
		fmt.Println("NEW PASSWORD", newPassword)
		fmt.Println("NEW PHONE NUMBER", newPhoneNumber)
		if newEmail != professional.Email {
			fmt.Println("TRYING TO CHANGE EMAIL")
			//Check if a user with this email already exists
			checkUser, err := dbclient.getProfessional(newEmail)
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
			professional.Email = newEmail
		}
		professional.FirstName = newFirstName
		professional.LastName = newLastName
		if newPassword != "" {
			fmt.Println("TRYING TO CHANGE PASSWORD")
			professional.Password = getMD5Hash(newPassword)
		}
		professional.PhoneNumber = newPhoneNumber
		var photoPath string
		if filerError == nil { //If image given
			fmt.Println("TRYING TO UPDATE PROFILE PICTURE")
			//Check if file is an image
			extension := filepath.Ext(file.Filename)
			if !validImgExtension(extension) {
				fmt.Println("Not a valid file")
				c.JSON(http.StatusBadRequest, gin.H{"error": "File is not an image"})
				return
			}
			//Upload the photo
			photoPath = filepath.Join(mediaDir, professional.Email, "profilePhoto", file.Filename)
			c.SaveUploadedFile(file, photoPath)
			//Path to save in the database
			photoPath = filepath.Join(professional.Email, "profilePhoto", file.Filename)
			professional.Photo = photoPath
		}
		professional.update() //Update user in the database
		fmt.Println("GOT HERE")
		//UPDATE THE SESSION
		c.JSON(http.StatusOK, gin.H{"message": "Profile updated"})
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
			//Create professional object
			professional := Professional{}
			professional.ID = professionalID.(int)
			educationInfo.setID()
			err = professional.addEducation(educationInfo)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			} else {
				c.JSON(http.StatusCreated, gin.H{"educationInfo": educationInfo})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Necessary fields not given"})
		}
	}
}

//POST /v1/LinkedIn/removeEducation
func removeEducation(c *gin.Context) {
	session := sessions.Default(c)
	professionalID := session.Get("userID")
	if professionalID == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var educationInfo Education
		if err := c.ShouldBindJSON(&educationInfo); err == nil {
			//Create professional object
			professional := Professional{}
			professional.ID = professionalID.(int)
			err = professional.removeEducation(educationInfo)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Necessary fields not given"})
		}
	}
}

//GET /v1/LinkedIn/getEducation
func getEducation(c *gin.Context) {
	session := sessions.Default(c)
	professionalID := session.Get("userID")
	if professionalID == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		//Create professional object
		professional := Professional{}
		professional.ID = professionalID.(int)
		education, err := professional.getEducation()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		} else {
			c.JSON(http.StatusAccepted, gin.H{"education": education})
		}
	}
}

//POST /v1/LinkedIn/addExperience
func addExperience(c *gin.Context) {
	session := sessions.Default(c)
	professionalID := session.Get("userID")
	if professionalID == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var experienceInfo Experience
		if err := c.ShouldBindJSON(&experienceInfo); err == nil {
			//Create professional object
			professional := Professional{}
			professional.ID = professionalID.(int)
			experienceInfo.setID()
			err = professional.addExperience(experienceInfo)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			} else {
				c.JSON(http.StatusCreated, gin.H{"experienceInfo": experienceInfo})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Necessary fields not given"})
		}
	}
}

//POST /v1/LinkedIn/removeExperience
func removeExperience(c *gin.Context) {
	session := sessions.Default(c)
	professionalID := session.Get("userID")
	if professionalID == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var experienceInfo Experience
		if err := c.ShouldBindJSON(&experienceInfo); err == nil {
			//Create professional object
			professional := Professional{}
			professional.ID = professionalID.(int)
			err = professional.removeExperience(experienceInfo)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Necessary fields not given"})
		}
	}
}

//GET /v1/LinkedIn/getExperience
func getExperience(c *gin.Context) {
	session := sessions.Default(c)
	professionalID := session.Get("userID")
	if professionalID == nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		//Create professional object
		professional := Professional{}
		professional.ID = professionalID.(int)
		experience, err := professional.getExperience()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		} else {
			c.JSON(http.StatusAccepted, gin.H{"experience": experience})
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
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		fmt.Println(professional)
		professional.setPhotoURL() //Create photo url
		c.JSON(http.StatusAccepted, gin.H{"status": "Authenticated", "professional": professional})
	}
}
