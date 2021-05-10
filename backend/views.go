package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
					setProfessionalSession(c, professional)
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

//GET /v1/LinkedIn/searchProfessional/:query
func searchProfessional(c *gin.Context) {
	_, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		queryString := c.Query("query")
		fmt.Println("QUERY STRING IS ", queryString)
		searchResults, err := dbclient.searchProfessional(queryString)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Interval server error"})
		} else {
			c.JSON(http.StatusOK, gin.H{"results": searchResults})
		}
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
		if newEmail != professional.Email {
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
		if len(newPassword) != 0 {
			professional.Password = getMD5Hash(newPassword)
		}
		professional.PhoneNumber = newPhoneNumber
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
			photoPath = filepath.Join(mediaDir, professional.Email, "profilePhoto", file.Filename)
			c.SaveUploadedFile(file, photoPath)
			//Path to save in the database
			photoPath = filepath.Join(professional.Email, "profilePhoto", file.Filename)
			professional.Photo = photoPath
		}
		professional.update() //Update user in the database
		//UPDATE THE SESSION
		setProfessionalSession(c, professional)
		professional.setPhotoURL() //Change the path to a url
		c.JSON(http.StatusOK, gin.H{"profile": professional})
	}
}

//GET /v1/LinkedIn/professional?:id
func getProfessionalProfile(c *gin.Context) {
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		email := c.Query("email")
		fmt.Println("QUERY STRING IS ", email)
		if email == "self" {
			email = professional.Email
		}
		prof, err := dbclient.getProfessional(email)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Interval server error"})
			return
		}
		prof.setPhotoURL() //Change the path of the photo to a url
		education, err := prof.getEducation()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Interval server error"})
			return
		}
		experience, err := prof.getExperience()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Interval server error"})
			return
		}
		skills, err := prof.getSkills()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Interval server error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"professional": prof, "education": education, "experience": experience, "skills": skills})
	}
}

//POST /v1/LinkedIn/addEducation
func addEducation(c *gin.Context) {
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var educationInfo Education
		if err := c.ShouldBindJSON(&educationInfo); err == nil {
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
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var educationInfo Education
		if err := c.ShouldBindJSON(&educationInfo); err == nil {
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
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
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
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var experienceInfo Experience
		if err := c.ShouldBindJSON(&experienceInfo); err == nil {
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
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var experienceInfo Experience
		if err := c.ShouldBindJSON(&experienceInfo); err == nil {
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
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		experience, err := professional.getExperience()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		} else {
			c.JSON(http.StatusAccepted, gin.H{"experience": experience})
		}
	}
}

//POST /v1/LinkedIn/addSkill
func addSkill(c *gin.Context) {
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var skillInfo Skill
		if err := c.ShouldBindJSON(&skillInfo); err == nil {
			err = professional.addSkill(skillInfo)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			} else {
				c.JSON(http.StatusCreated, gin.H{"skillInfo": skillInfo})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Necessary fields not given"})
		}
	}
}

//POST /v1/LinkedIn/removeSkill
func removeSkill(c *gin.Context) {
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var skillInfo Skill
		if err := c.ShouldBindJSON(&skillInfo); err == nil {
			err = professional.removeSkill(skillInfo)
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

//GET /v1/LinkedIn/getSkills
func getSkills(c *gin.Context) {
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		skills, err := professional.getSkills()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		} else {
			c.JSON(http.StatusAccepted, gin.H{"skills": skills})
		}
	}
}

//POST /v1/LinkedIn/addArticle
func addArticle(c *gin.Context) {
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var newArticle Article
		file, filerError := c.FormFile("file")
		newArticle.Title = c.PostForm("title")
		newArticle.Content = c.PostForm("content")
		newArticle.UploaderID = professional.ID
		var filePath string
		if filerError == nil { //In case of attached file
			//Check if is a valid file
			extension := filepath.Ext(file.Filename)
			if !validFileExtension(extension) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type"})
				return
			}
			//Create a directory for the artcile
			now := time.Now()
			os.MkdirAll(filepath.Join(mediaDir, professional.Email, newArticle.Title, now.Format("2006-02-01 15:04:05.000 MST")), 0755)
			//Upload the file
			filePath = filepath.Join(mediaDir, professional.Email, newArticle.Title, now.Format("2006-02-01 15:04:05.000 MST"), file.Filename)
			c.SaveUploadedFile(file, filePath)
			//Path to save in the database
			filePath = filepath.Join(professional.Email, newArticle.Title, now.Format("2006-02-01 15:04:05.000 MST"), file.Filename)
		} else {
			filePath = ""
		}
		newArticle.AttachedFile = filePath
		newArticle.save()
		c.JSON(http.StatusCreated, gin.H{"articleInfo": newArticle})
	}
}

//GET /v1/LinkedIn/getArticles
func getArticles(c *gin.Context) {
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		feed, err := professional.getFeed()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		} else {
			c.JSON(http.StatusAccepted, gin.H{"articles": feed})
		}
	}
}

//POST /v1/LinkedIn/getArticleDetails
func getArticleDetails(c *gin.Context) {
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var article Article
		if err := c.ShouldBindJSON(&article); err == nil {
			articleUploader, err := article.getUploader()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
				return
			}
			hasImage, err := article.fileIsImage()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
				return
			}
			//Check if user liked this article
			liked, err := professional.likedArticle(article)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
				return
			}
			//Get the comments of the article
			comments, err := article.getComments()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
				return
			}
			//Get the likes of the article
			likes, err := article.getLikes()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"uploader": articleUploader, "hasImage": hasImage, "liked": liked, "comments": comments, "likes": likes})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		}
	}
}

//POST /v1/LinkedIn/article/addLike
func addArticleLike(c *gin.Context) {
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var article Article
		if err := c.ShouldBindJSON(&article); err == nil {
			//Create like object
			like := ArticleLike{}
			like.ProfessionalID = professional.ID //ID of the user who liked the article
			err = article.addLike(like)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
				return
			}
			if professional.ID != article.UploaderID {
				//Create a notification
				notification := Notification{}
				notification.ProfessionalID = article.UploaderID
				notification.Seen = false
				notification.Msg = fmt.Sprintf("%s %s liked your article with title %s", professional.FirstName, professional.LastName, article.Title)
				err = notification.save()
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
					return
				}
			}
			c.JSON(http.StatusOK, gin.H{"message": "Like added"})

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		}
	}
}

//POST /v1/LinkedIn/article/removeLike
func removeArticleLike(c *gin.Context) {
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var article Article
		if err := c.ShouldBindJSON(&article); err == nil {
			//Create like object
			like := ArticleLike{}
			like.ProfessionalID = professional.ID //ID of the user who unliked the article
			err = article.removeLike(like)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Like removed"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		}
	}
}

//POST /v1/LinkedIn/article/addComment
func addArticleComment(c *gin.Context) {
	professional, err := getProfessionalFromSession(c) //Get professional object from session
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var comment ArticleComment
		if err := c.ShouldBindJSON(&comment); err == nil {
			//Get article Object
			article, err := dbclient.getArticle(comment.ArticleID)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
				return
			}
			comment.ProfessionalID = professional.ID //ID of the user who liked the article
			err = article.addComment(comment)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			} else {
				//Create a notification
				uploader, err := article.getUploader()
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
					return
				}
				if professional.ID != uploader.ID {
					notification := Notification{}
					notification.ProfessionalID = uploader.ID
					notification.Seen = false
					notification.Msg = fmt.Sprintf("%s %s commented %s on your article with title %s", professional.FirstName, professional.LastName, comment.Comment, article.Title)
					err = notification.save()
					if err != nil {
						fmt.Println(err)
						c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
						return
					}
				}
				response := ArticleCommentResponse{}
				response.ID = -1 //temporary ID until the user refreshes the page
				response.FirstName = professional.FirstName
				response.LastName = professional.LastName
				response.Comment = comment.Comment
				c.JSON(http.StatusOK, gin.H{"comment": response})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		}
	}
}

//GET v1/LinkedIn/homepage
//Get the number of notifications and unread messages of a professional
func homepage(c *gin.Context) {
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		//Get professional's new notifications
		notifications, err := professional.getNewNotifications()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}
		//Get professional's unread message dialogs
		unreadCount, err := professional.getUnreadDialogs()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"notifications": len(notifications), "unreadMessages": unreadCount})
	}
}

//GET v1/LinkedIn/notifications
//Get the notifications and messages of a professional
func getNotifications(c *gin.Context) {
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		//Get professional notifications
		notifications, err := professional.getNotifications()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}
		//Set the new notifications to seen
		err = professional.clearNotifications()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"notifications": notifications})
	}
}

//POST /v1/LinkedIn/addFriendRequest
func addFriendRequest(c *gin.Context) {
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var f Friendship
		if err := c.ShouldBindJSON(&f); err == nil {
			//Create a friend request
			err = dbclient.createFriendRequest(professional.ID, f.ProfessionalID2)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			} else {
				//Create a notification
				notification := Notification{}
				notification.ProfessionalID = f.ProfessionalID2
				notification.Seen = false
				notification.Msg = fmt.Sprintf("%s %s sent you a friend request", professional.FirstName, professional.LastName)
				err = notification.save()
				if err != nil {
					fmt.Println(err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
					return
				}
				c.JSON(http.StatusOK, gin.H{"message": "Sent friend request"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		}
	}
}

//POST /v1/LinkedIn/addFriend
func addFriend(c *gin.Context) {
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var f Friendship
		if err := c.ShouldBindJSON(&f); err == nil {
			//Create a friend request
			err = dbclient.createFriendship(f.ProfessionalID2, professional.ID)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Frienship created"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		}
	}
}

//POST /v1/LinkedIn/removeFriendRequest
func removeFriendRequest(c *gin.Context) {
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var f Friendship
		if err := c.ShouldBindJSON(&f); err == nil {
			//Delete friend request
			err = dbclient.deleteFriendRequest(f.ProfessionalID2, professional.ID)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Request deleted"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		}
	}
}

//GET /v1/LinkedIn/friendRequests
func getFriendRequests(c *gin.Context) {
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		requests, err := dbclient.getProfessionalFriendRequests(professional.ID)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		} else {
			c.JSON(http.StatusOK, gin.H{"requests": requests})
		}
	}
}

//GET /v1/LinkedIn/friendshipStatus?id
func friendshipStatus(c *gin.Context) {
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		id := c.Query("id")
		professionalID2, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}
		status, err := dbclient.getFriendshipStatus(professional.ID, professionalID2)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}
		if len(status) == 0 {
			//Check if there is a pending request from professional2 to professional
			//Or if there is a connection from professional2 to professional
			status, err = dbclient.getFriendshipStatus(professionalID2, professional.ID)
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
				return
			}
			if status == "pending" { //Professional can accept the friend request
				status = "accept"
			}
			c.JSON(http.StatusOK, gin.H{"status": status})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": status})
		}
	}
}

//GET /v1/LinkedIn/connectedProfessionals
func getConnectedProfessionals(c *gin.Context) {
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		connectedProfessionals, err := professional.getFriends()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		} else {
			c.JSON(http.StatusOK, gin.H{"connectedProfessionals": connectedProfessionals})
		}
	}
}

//POST /v1/LinkedIn/sendMessage
func sendMessage(c *gin.Context) {
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		var m Message
		if err := c.ShouldBindJSON(&m); err == nil {
			m.Sender = professional.ID
			err = m.save() //Save message in the database
			if err != nil {
				fmt.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			} else {
				//Create ChatMessage response
				response := ChatMessage{}
				response.FirstName = professional.FirstName
				response.LastName = professional.LastName
				response.Photo = professional.Photo
				response.Msg = m.Msg
				response.setPhotoURL()
				c.JSON(http.StatusOK, gin.H{"message": response})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
		}
	}
}

//GET /v1/LinkedIn/chat?id
func getChatMessages(c *gin.Context) {
	professional, err := getProfessionalFromSession(c)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not authenticated"})
	} else {
		id := c.Query("id")
		professionalID2, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}
		chatMessages, err := dbclient.getChat(professional.ID, professionalID2)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		} else {
			//Update unseen messages to seen
			dbclient.updateMessagesStatus(professional.ID, professionalID2)
			c.JSON(http.StatusOK, gin.H{"chat": chatMessages})
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
		//fmt.Println(professional)
		professional.setPhotoURL() //Create photo url
		c.JSON(http.StatusAccepted, gin.H{"status": "Authenticated", "professional": professional})
	}
}
