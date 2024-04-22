package controllers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HitungSegitigaSamaSisi(c *gin.Context) {
	alasStr := c.Query("alas")
	tinggiStr := c.Query("tinggi")
	hitung := c.Query("hitung")

	alas, err1 := strconv.ParseFloat(alasStr, 64)
	tinggi, err2 := strconv.ParseFloat(tinggiStr, 64)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter alas dan tinggi harus angka"})
		return
	}

	var result float64

	if hitung == "luas" {
		result = 0.5 * alas * tinggi
	} else if hitung == "keliling" {
		result = alas * 3
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter hitung tidak valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{hitung: result})
}

func HitungPersegi(c *gin.Context) {
	sisiStr := c.Query("sisi")
	hitung := c.Query("hitung")

	sisi, err1 := strconv.ParseFloat(sisiStr, 64)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter alas dan tinggi harus angka"})
		return
	}

	var result float64

	if hitung == "luas" {
		result = sisi * sisi
	} else if hitung == "keliling" {
		result = 4 * sisi
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter hitung tidak valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{hitung: result})
}

func HitungPersegiPanjang(c *gin.Context) {
	panjangStr := c.Query("panjang")
	lebarStr := c.Query("lebar")
	hitung := c.Query("hitung")

	panjang, err1 := strconv.ParseFloat(panjangStr, 64)
	lebar, err2 := strconv.ParseFloat(lebarStr, 64)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter alas dan tinggi harus angka"})
		return
	}

	var result float64

	if hitung == "luas" {
		result = panjang * lebar
	} else if hitung == "keliling" {
		result = 2 * (panjang + lebar)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter hitung tidak valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{hitung: result})

}

func HitungLingkaran(c *gin.Context) {
	jariJari := c.Query("jariJari")
	hitung := c.Query("hitung")

	r, err1 := strconv.ParseFloat(jariJari, 64)

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter alas dan tinggi harus angka"})
		return
	}

	var result float64

	if hitung == "luas" {
		result = math.Pi * r * r
	} else if hitung == "keliling" {
		result = 2 * math.Pi * r
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameter hitung tidak valid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{hitung: result})

}
