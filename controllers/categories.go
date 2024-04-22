package controllers

import (
	"net/http"
	"quiz3/database"
	"quiz3/repository"
	"quiz3/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ShowCategoriesBook(c *gin.Context) {
	var (
		result   gin.H
		category structs.Category
	)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category.ID = id
	books, err := repository.FilterCategoriesBook(database.DbConnection, category)
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

func GetAllCategories(c *gin.Context) {
	var (
		result gin.H
	)

	category, err := repository.GetAllCategories(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": category,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertCategories(c *gin.Context) {
	var category structs.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Category",
	})
}

func UpdateCategories(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = int(id)
	err = repository.UpdateCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Category",
	})
}

func DeleteCategories(c *gin.Context) {
	var category structs.Category
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category.ID = id
	err = repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Category",
	})
}
