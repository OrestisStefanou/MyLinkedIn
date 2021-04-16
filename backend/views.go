package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user Professional
	file, _ := c.FormFile("photo")
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
	//Check if file is an image
	extension := filepath.Ext(file.Filename)
	if !validImgExtension(extension) {
		fmt.Println("Not a valid file")
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is not an image"})
		return
	}
	//Create a directory for the user's media
	os.MkdirAll(filepath.Join(mediaDir, user.Email, "profilePhoto"), 0755)
	photoPath := filepath.Join(mediaDir, user.Email, "profilePhoto", file.Filename)
	c.SaveUploadedFile(file, photoPath)
	user.Photo = photoPath
	user.save() //Save the new user in the database
	c.JSON(http.StatusCreated, gin.H{"message": "Signed up successfully"})
}
