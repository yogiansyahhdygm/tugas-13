package main

import (
	"tugas13/config"
	"tugas13/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()

	// ROUTES CRUD
	r.POST("/bioskop", controllers.TambahBioskop)
	r.GET("/bioskop", controllers.GetSemuaBioskop)
	r.GET("/bioskop/:id", controllers.GetBioskopByID)
	r.PUT("/bioskop/:id", controllers.UpdateBioskop)
	r.DELETE("/bioskop/:id", controllers.HapusBioskop)
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Aplikasi Bioskop API Aktif"})
	})

	r.Run(":8080")
}
