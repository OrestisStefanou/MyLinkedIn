package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	dbclient.initialize()
	defer dbclient.db.Close()
	//Log file
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()

	//Session middleware
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	//Cors middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	v1 := router.Group("/v1/LinkedIn")
	{
		//Registration endpoints
		v1.POST("/signup", signup)
		v1.POST("/signin", signin)
		//Education endpoints
		v1.POST("/addEducation", addEducation)
		v1.GET("/getEducation", getEducation)
		v1.POST("/removeEducation", removeEducation)
		//Experience endpoints
		v1.POST("/addExperience", addExperience)
		v1.POST("/removeExperience", removeExperience)
		v1.GET("/getExperience", getExperience)
		//Skills endpoints
		v1.POST("/addSkill", addSkill)
		v1.POST("/removeSkill", removeSkill)
		v1.GET("/getSkills", getSkills)
		//Professional endpoints
		v1.POST("/updateProfessional", updateProfessional)
		v1.GET("/homepage", homepage)
		v1.GET("/searchProfessional", searchProfessional)
		v1.GET("/professional", getProfessionalProfile)
		//Article endpoints
		v1.POST("/addArticle", addArticle)
		v1.GET("/getArticles", getArticles)
		v1.POST("/getArticleDetails", getArticleDetails)
		v1.POST("/article/addLike", addArticleLike)
		v1.POST("/article/removeLike", removeArticleLike)
		v1.POST("/article/addComment", addArticleComment)
		v1.GET("/getProfessionalArticles", getProfessionalArticles)
		//Notifications endpoints
		v1.GET("/notifications", getNotifications)
		//Friendship endpoints
		v1.GET("/friendshipStatus", friendshipStatus)
		v1.POST("/addFriendRequest", addFriendRequest)
		v1.POST("/removeFriendRequest", removeFriendRequest)
		v1.POST("/addFriend", addFriend)
		v1.GET("/friendRequests", getFriendRequests)
		v1.GET("/connectedProfessionals", getConnectedProfessionals)
		//Message endpoints
		v1.POST("/sendMessage", sendMessage)
		v1.GET("/chat", getChatMessages)
		v1.GET("/chatDialogs", chatDialogs)
		//JobAd endpoints
		v1.POST("/addJobAd", addJobAd)
		v1.GET("/jobAds", getJobAds)
		v1.POST("/jobAd/addInterest", addJobInterest)
		v1.POST("/jobAd/removeInterest", removeJobInterest)
		v1.POST("/jobAd/addComment", addJobAdComment)
		v1.POST("/getJobAdDetails", getJobAdDetails)
		v1.GET("/professionalJobAds", getProfeesionalJobAds)
		v1.POST("/removeJobAd", removeJobAd)
		//Logout and session endpoints
		v1.GET("/logout", logout)
		v1.GET("/authenticated", authenticated)
		//Media endpoint
		v1.StaticFS("/media", http.Dir("./media"))
	}
	adminEndpoints := router.Group("/admin/LinkedIn")
	{
		adminEndpoints.POST("/signin", adminSignin)
		adminEndpoints.GET("/authenticated", adminAuthenticated)
		adminEndpoints.GET("/allUsers", getAllUsers)
		adminEndpoints.POST("/jsonUsers", jsonUsers)
		adminEndpoints.POST("/xmlUsers", xmlUsers)
		adminEndpoints.GET("/professional", getUserProfile)
	}
	router.Run()
}
