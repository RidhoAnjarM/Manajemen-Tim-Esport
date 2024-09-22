package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ada file yang diterima"})
		return
	}

	if file.Size > 1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ukuran file harus dibawah 1MB"})
		return
	}

	if !strings.HasSuffix(file.Filename, ".jpg") && !strings.HasSuffix(file.Filename, ".png") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak valid. Hanya .jpg dan .png yang diperbolehkan"})
		return
	}

	filePath := fmt.Sprintf("uploads/%s", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File berhasil diunggah", "file_url": filePath})
}