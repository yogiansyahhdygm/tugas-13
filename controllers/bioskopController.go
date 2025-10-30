package controllers

import (
	"net/http"
	"tugas13/config"
	"tugas13/models"

	"github.com/gin-gonic/gin"
)

func TambahBioskop(c *gin.Context) {
	var input models.Bioskop

	// Ambil data dari JSON body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	// Validasi
	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	sqlStatement := `
		INSERT INTO bioskop (nama, lokasi, rating)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := config.DB.QueryRow(sqlStatement, input.Nama, input.Lokasi, input.Rating).Scan(&input.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil ditambahkan", "data": input})
}
