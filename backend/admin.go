package main

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//Admin json struct
type Admin struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	PhoneNumber string `json:"phoneNumber"`
}

//Create or update an admin's session
func setAdminSession(c *gin.Context, admin Admin) {
	session := sessions.Default(c)
	session.Set("adminID", admin.ID)
	session.Set("adminFirstName", admin.FirstName)
	session.Set("adminLastName", admin.LastName)
	session.Set("adminEmail", admin.Email)
	session.Set("adminPhone", admin.PhoneNumber)
	session.Save()
}

func getAdminFromSession(c *gin.Context) (Admin, error) {
	admin := Admin{}
	//Create admin object from session
	session := sessions.Default(c)
	adminID := session.Get("adminID")
	if adminID == nil {
		return admin, errors.New("Not authenticated")
	}
	firstName := session.Get("adminFirstName")
	lastName := session.Get("adminLastName")
	email := session.Get("adminEmail")
	phoneNumber := session.Get("adminPhone")
	admin.ID = adminID.(int)
	admin.FirstName = firstName.(string)
	admin.LastName = lastName.(string)
	admin.Email = email.(string)
	admin.PhoneNumber = phoneNumber.(string)
	return admin, nil

}

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

func (driver *DBClient) getAdmin(email string) (Admin, error) {
	adminInfo := Admin{}
	rows, err := driver.db.Query("SELECT * FROM Admins WHERE Email=?", email)
	if err != nil {
		return adminInfo, err
	}
	for rows.Next() {
		err = rows.Scan(&adminInfo.ID, &adminInfo.FirstName, &adminInfo.LastName, &adminInfo.Email, &adminInfo.Password, &adminInfo.PhoneNumber)
		if err != nil {
			return adminInfo, err
		}
	}
	return adminInfo, nil
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
