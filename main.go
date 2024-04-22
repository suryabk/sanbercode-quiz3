package main

import (
	"net/http"
	"quiz3/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/bangun-datar/:jenis", func(c *gin.Context) {
		jenis := c.Param("jenis")
		switch jenis {
		case "segitiga-sama-sisi":
			controllers.HitungSegitigaSamaSisi(c)
		case "persegi":
			controllers.HitungPersegi(c)
		case "persegi-panjang":
			controllers.HitungPersegiPanjang(c)
		case "lingkaran":
			controllers.HitungLingkaran(c)
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Jenis bangun datar tidak valid"})
		}
	})

	router.Run(":8080")
}
