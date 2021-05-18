package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

//GET /admin/LinkedIn/authenticated
func adminAuthenticated(c *gin.Context) {
	admin, err := getAdminFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		//fmt.Println(professional)
		c.JSON(http.StatusAccepted, gin.H{"status": "Authenticated", "admin": admin})
	}
}

//POST /admin/LinkedIn/signin
func adminSignin(c *gin.Context) {
	type adminLoginInfo struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	var adminInfo adminLoginInfo
	if err := c.ShouldBindJSON(&adminInfo); err == nil {
		admin, err := dbclient.getAdmin(adminInfo.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Interval server error"})
		} else {
			if admin.ID == 0 { //An admin with this email does not exist
				c.JSON(http.StatusNotFound, gin.H{"error": "Wrong email or password"})
			} else {
				//Get md5 hash of password
				md5pass := getMD5Hash(adminInfo.Password)
				if md5pass == admin.Password {
					setAdminSession(c, admin)
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

//GET /admin/LinkedIn/allUsers
func getAllUsers(c *gin.Context) {
	_, err := getAdminFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		users, err := dbclient.getAllUsers()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		} else {
			c.JSON(http.StatusOK, gin.H{"users": users})
		}
	}
}

//POST /admin/LinkedIn/jsonUsers
func jsonUsers(c *gin.Context) {
	type UsersIDArray struct {
		UsersIDs []int `json:"ids"`
	}
	//First check if admin is authencticated
	var idArray UsersIDArray
	if err := c.ShouldBindJSON(&idArray); err == nil {
		fmt.Println(idArray.UsersIDs)
		var data []UserDetailInfo
		for i := 0; i < len(idArray.UsersIDs); i++ {
			detailInfo, err := dbclient.getUserDetails(idArray.UsersIDs[i])
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
				return
			}
			data = append(data, detailInfo)
		}
		//Create the json file
		path := filepath.Join(adminDir, "users.json")
		jsonFile, _ := json.MarshalIndent(data, "", " ")
		_ = ioutil.WriteFile(path, jsonFile, 0644)
		//Return the file url
		fileURL := mediaURL + "admin/" + "users.json"
		c.JSON(http.StatusOK, gin.H{"users": fileURL})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are necessary"})
	}
}
