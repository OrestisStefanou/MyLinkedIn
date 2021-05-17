package main

import (
	"errors"

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

func (driver *DBClient) getAllUsers() ([]UserResponse, error) {
	prof := Professional{}
	user := UserResponse{}
	sql := "SELECT * FROM Professionals"
	rows, err := driver.db.Query(sql)
	if err != nil {
		return nil, err
	}
	var response []UserResponse
	for rows.Next() {
		err = rows.Scan(&prof.ID, &prof.FirstName, &prof.LastName, &prof.Email, &prof.Password, &prof.PhoneNumber, &prof.Photo)
		if err != nil {
			return nil, err
		}
		prof.setPhotoURL() //Change the path of a photo to a url
		user.User = prof
		user.Checked = false
		response = append(response, user)
	}
	return response, nil
}
