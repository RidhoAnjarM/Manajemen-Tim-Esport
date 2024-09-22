package controllers

import (
	"net/http"
	"tim-esport/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreatePlayer(c *gin.Context) {
	var player models.Player

	if err := c.ShouldBindJSON(&player); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string

		// Custom error messages untuk field yang tidak valid
		for _, fieldError := range validationErrors {
			switch fieldError.Field() {
			case "Name":
				errorMessages = append(errorMessages, "Nama pemain harus diisi.")
			case "Position":
				errorMessages = append(errorMessages, "Posisi harus diisi.")
			case "Game":
				errorMessages = append(errorMessages, "Game harus diisi dengan pilihan yang benar (Dota2, CS:GO, Valorant, PUBGM, MLBB).")
			case "Profil":
				errorMessages = append(errorMessages, "profil harus diisi juga yaa.")
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	if err := models.DB.Create(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pemain berhasil dibuat", "player": player})
}

func UpdatePlayer(c *gin.Context) {
	var player models.Player
	id := c.Param("id")

	if err := models.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player tidak ditemukan."})
		return
	}

	if err := c.ShouldBindJSON(&player); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var errorMessages []string

		// Custom error messages untuk field yang tidak valid
		for _, fieldError := range validationErrors {
			switch fieldError.Field() {
			case "Name":
				errorMessages = append(errorMessages, "Nama pemain harus diisi.")
			case "Position":
				errorMessages = append(errorMessages, "Posisi harus diisi.")
			case "Game":
				errorMessages = append(errorMessages, "Game harus diisi dengan pilihan yang benar (Dota2, CS:GO, Valorant, PUBGM, MLBB).")
			case "Profil":
				errorMessages = append(errorMessages, "profil harus diisi.")
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	// Simpan perubahan
	if err := models.DB.Save(&player).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Player berhasil diperbarui", "player": player})
}

func GetPlayers(c *gin.Context) {
	var players []models.Player
	if err := models.DB.Find(&players).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, players)
}

func GetPlayer(c *gin.Context) {
	var player models.Player
	id := c.Param("id")
	if err := models.DB.First(&player, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, player)
}

func DeletePlayer(c *gin.Context) {
	id := c.Param("id")
	if err := models.DB.Delete(&models.Player{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Player berhasil dihapus"})
}
