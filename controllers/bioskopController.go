package controllers

import (
	"net/http"
	"strconv"
	"tugas13/config"
	"tugas13/models"

	"github.com/gin-gonic/gin"
)

// CREATE - POST
func TambahBioskop(c *gin.Context) {
	var input models.Bioskop
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	//validasi
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

// GET /bioskop
func GetSemuaBioskop(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, nama, lokasi, rating FROM bioskop")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	defer rows.Close()

	var hasil []models.Bioskop
	for rows.Next() {
		var b models.Bioskop
		if err := rows.Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca data"})
			return
		}
		hasil = append(hasil, b)
	}

	c.JSON(http.StatusOK, gin.H{"data": hasil})
}

// GET /bioskop/:id
func GetBioskopByID(c *gin.Context) {
	id := c.Param("id")

	var b models.Bioskop
	err := config.DB.QueryRow("SELECT id, nama, lokasi, rating FROM bioskop WHERE id=$1", id).
		Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": b})
}

// PUT /bioskop/:id
func UpdateBioskop(c *gin.Context) {
	id := c.Param("id")
	var input models.Bioskop

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	if input.Nama == "" || input.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi tidak boleh kosong"})
		return
	}

	sqlStatement := `
		UPDATE bioskop
		SET nama=$1, lokasi=$2, rating=$3
		WHERE id=$4
	`
	res, err := config.DB.Exec(sqlStatement, input.Nama, input.Lokasi, input.Rating, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui data"})
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	idInt, _ := strconv.Atoi(id)
	input.ID = idInt
	c.JSON(http.StatusOK, gin.H{"message": "Data bioskop berhasil diperbarui", "data": input})
}

// DELETE /bioskop/:id
func HapusBioskop(c *gin.Context) {
	id := c.Param("id")

	sqlStatement := `DELETE FROM bioskop WHERE id=$1`
	res, err := config.DB.Exec(sqlStatement, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus data"})
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil dihapus"})
}
