package controllers

import (
	"net/http"
	"quiz3/database"
	"quiz3/repository"
	"quiz3/structs"
	"quiz3/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context) {
	var (
		result gin.H
	)

	books, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertBook(c *gin.Context) {
	var input structs.InputBook

	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(err)
	}

	// Validate release_year
	if input.ReleaseYear < 1980 || input.ReleaseYear > 2021 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year harus berada di antara 1980 dan 2021"})
		return
	}

	book := structs.Book{
		Title:       input.Title,
		Description: input.Description,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		Thickness:   utils.ThicknessStatus(input.TotalPage),
		CategoryID:  input.CategoryID,
	}

	err = repository.InsertBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Book",
	})
}

func UpdateBook(c *gin.Context) {
	var input structs.InputBook
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&input)
	if err != nil {
		panic(err)
	}

	if input.ReleaseYear < 1980 || input.ReleaseYear > 2021 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "release_year harus berada di antara 1980 dan 2021"})
		return
	}

	book := structs.Book{
		Title:       input.Title,
		Description: input.Description,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		Thickness:   utils.ThicknessStatus(input.TotalPage),
		CategoryID:  input.CategoryID,
	}

	book.ID = int(id)
	err = repository.UpdateBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Book",
	})
}

func DeleteBook(c *gin.Context) {
	var book structs.Book
	id, err := strconv.Atoi(c.Param("id"))

	book.ID = int(id)

	err = repository.DeleteBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Book",
	})
}
