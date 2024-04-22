package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"quiz3/controllers"
	"quiz3/database"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB               *sql.DB
	err              error
	usernamePassword = map[string]string{
		"admin":  "password",
		"editor": "secret",
	}
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if (uname == "admin" && pwd == "password") || (uname == "editor" && pwd == "secret") {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success read file environment")
	}

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

	router.Use(basicAuthMiddleware())
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

	// router.GET("/books", getBooks)
	// router.POST("/books", createBook)
	// router.PUT("/books/:id", updateBook)
	// router.DELETE("/books/:id", deleteBook)

	router.Run(":8080")
}

func basicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == http.MethodPost || method == http.MethodPut || method == http.MethodDelete {
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				unauthorized(c)
				return
			}

			auth := strings.SplitN(authHeader, " ", 2)
			if len(auth) != 2 || auth[0] != "Basic" {
				unauthorized(c)
				return
			}

			payload, err := base64.StdEncoding.DecodeString(auth[1])
			if err != nil {
				unauthorized(c)
				return
			}

			pair := strings.SplitN(string(payload), ":", 2)
			if len(pair) != 2 {
				unauthorized(c)
				return
			}

			username := pair[0]
			password := pair[1]

			if storedPassword, ok := usernamePassword[username]; !ok || storedPassword != password {
				unauthorized(c)
				return
			}
		}

		c.Next()
	}
}
func unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
	c.Abort()
}
