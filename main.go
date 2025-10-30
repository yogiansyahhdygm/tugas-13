package main

import (
	"tugas13/config"
	"tugas13/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Koneksi ke database
	config.ConnectDatabase()

	// Buat router
	r := gin.Default()

	// Route POST /bioskop
	r.POST("/bioskop", controllers.TambahBioskop)

	// Jalankan server
	r.Run(":8080")
}
