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
	user.Password = c.PostForm("password")
	user.PhoneNumber = c.PostForm("phoneNumber")
	log.Println(file.Filename, file.Size)
	fmt.Println("User info is ", user)
	//Create a directory for the user's media
	os.MkdirAll(filepath.Join(mediaDir, user.Email, "profilePhoto"), 0755)
	c.SaveUploadedFile(file, filepath.Join(mediaDir, user.Email, "profilePhoto"))
	//Hash the password before saving to the database
	c.JSON(http.StatusOK, gin.H{"Professional": user})
}
