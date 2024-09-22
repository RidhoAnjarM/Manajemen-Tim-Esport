package main

import (
	"os"
	"tim-esport/controllers"
	"tim-esport/middleware"
	"tim-esport/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	models.ConnectDatabase()

	r := gin.Default()

	r.Static("/uploads", "./uploads")

	//Registrasi dan login
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// Upload file
		auth.POST("/upload", controllers.UploadFile)

		// CRUD Team
		auth.POST("/team/create", controllers.CreateTeam)
		auth.GET("/team/get", controllers.GetTeams)
		auth.GET("/team/get/:id", controllers.GetTeamByID)
		auth.PUT("/team/update/:id", controllers.UpdateTeam)
		auth.DELETE("/team/delete/:id", controllers.DeleteTeam)

		// CRUD Player
		auth.POST("/player/create", controllers.CreatePlayer)
		auth.GET("/player/get", controllers.GetPlayers)
		auth.GET("/player/get/:id", controllers.GetPlayer)
		auth.PUT("/player/update/:id", controllers.UpdatePlayer)
		auth.DELETE("/player/delete/:id", controllers.DeletePlayer)
	}

	host := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	err = r.Run(host)
	if err != nil {
		panic(err)
	}
}
