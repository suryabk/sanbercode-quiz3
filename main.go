package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"quiz3/controllers"
	"quiz3/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "books"
// )

var (
	DB  *sql.DB
	err error
)

func Auth(c *gin.Context) {
	uname, pwd, ok := c.Request.BasicAuth()
	if !ok {
		c.String(http.StatusUnauthorized, "Username atau Password tidak boleh kosong")
		c.Abort()
		return
	}

	if (uname == "admin" && pwd == "password") || (uname == "editor" && pwd == "secret") {
		c.Next()
		return
	}
	c.String(http.StatusUnauthorized, "Username atau Password tidak sesuai")
	c.Abort()
}

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success read file environment")
	}

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed: ", err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

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

	router.GET("/categories", controllers.GetAllCategories)
	router.POST("/categories", Auth, controllers.InsertCategories)
	router.PUT("/categories/:id", Auth, controllers.UpdateCategories)
	router.DELETE("/categories/:id", Auth, controllers.DeleteCategories)
	router.GET("/categories/:id/books", controllers.ShowCategoriesBook)

	router.GET("/books", controllers.GetAllBook)
	router.POST("/books", Auth, controllers.InsertBook)
	router.PUT("/books/:id", Auth, controllers.UpdateBook)
	router.DELETE("/books/:id", Auth, controllers.DeleteBook)

	// router.Run("localhost:8080")
	router.Run(":" + os.Getenv("PORT"))
}
