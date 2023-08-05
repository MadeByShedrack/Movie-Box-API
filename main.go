package main

import (
	"github.com/MadeByShedrack/controllers"
	"github.com/MadeByShedrack/database"
	"github.com/gin-gonic/gin"
)

func main() {
	movieRouter := gin.Default()

	database.ConnectDatabase()

	movieRouter.GET("/movies", controllers.FindBooks)

	movieRouter.Run()
}
