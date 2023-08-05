package controllers

import (
	"net/http"

	"github.com/MadeByShedrack/database"
	"github.com/MadeByShedrack/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	var movies []models.Movies
	database.DB.Find(&movies)

	c.JSON(http.StatusOK, gin.H{"message": movies})
}
