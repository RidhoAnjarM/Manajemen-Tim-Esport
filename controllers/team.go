package controllers

import (
	"net/http"
	"tim-esport/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CustomErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

// Fungsi untuk mengeluarkan pesan kesalahan kustom
func GetCustomErrorMessages(err error) []CustomErrorResponse {
	var errors []CustomErrorResponse
	for _, err := range err.(validator.ValidationErrors) {
		var element CustomErrorResponse
		element.FailedField = err.Field()
		element.Tag = err.Tag()
		element.Value = err.Param()
		errors = append(errors, element)
	}
	return errors
}

// Create Teams
func CreateTeam(c *gin.Context) {
	var team models.Team

	// Bind JSON dan validasi otomatis
	if err := c.ShouldBindJSON(&team); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string

		for _, fieldError := range validationErrors {
			switch fieldError.Field() {
			case "TeamName":
				errorMessages = append(errorMessages, "Nama tim harus diisi.")
			case "Game":
				errorMessages = append(errorMessages, "Game harus diisi dengan pilihan yang benar (Dota2, CS:GO, Valorant, PUBGM, MLBB).")
			case "Logo":
				errorMessages = append(errorMessages, "Logo harus diisi.")
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	// Jika validasi berhasil, simpan data tim
	if err := models.DB.Create(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tim berhasil dibuat", "team": team})
}

// Update Teams
func UpdateTeam(c *gin.Context) {
	var team models.Team
	id := c.Param("id")

	if err := models.DB.First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tim tidak ditemukan."})
		return
	}

	if err := c.ShouldBindJSON(&team); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string

		for _, fieldError := range validationErrors {
			switch fieldError.Field() {
			case "TeamName":
				errorMessages = append(errorMessages, "Nama tim harus diisi.")
			case "Game":
				errorMessages = append(errorMessages, "Game harus diisi dengan pilihan yang benar (Dota2, CS:GO, Valorant, PUBGM, MLBB).")
			case "Logo":
				errorMessages = append(errorMessages, "Logo harus diisi.")
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	// Simpan perubahan
	if err := models.DB.Save(&team).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tim berhasil diperbarui", "team": team})
}

// Get All Teams
func GetTeams(c *gin.Context) {
	var teams []models.Team
	if err := models.DB.Preload("Players").Find(&teams).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teams)
}

// Get Team by ID
func GetTeamByID(c *gin.Context) {
	var team models.Team
	id := c.Param("id")
	if err := models.DB.Preload("Players").First(&team, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tim tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, team)
}

// Delete Team
func DeleteTeam(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.Team{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tim tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tim berhasil dihapus"})
}
