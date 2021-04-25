package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	dbclient.initialize()
	router := gin.Default()

	//Session middleware
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	//Cors middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
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
		//Logout and session endpoints
		v1.GET("/logout", logout)
		v1.GET("/authenticated", authenticated)
		//Media endpoint
		v1.StaticFS("/media", http.Dir("./media"))
	}
	router.Run()
}
