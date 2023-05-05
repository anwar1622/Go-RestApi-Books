package controllers

import (
	"Go-RestApi-Books/database"
	"Go-RestApi-Books/models"
	"Go-RestApi-Books/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var appJSON = "application/json"

func CreateBooks(c *gin.Context) {
	db := database.GetDB()
	books := models.Book{}

	if appJSON == utils.GetContentType(c) {
		c.ShouldBindJSON(&books)
	} else {
		c.ShouldBind(&books)
	}

	err := db.Debug().Create(&books).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":     books.ID,
		"title":  books.Title,
		"author": books.Author,
	})
}

func GetBook(c *gin.Context) {
	db := database.GetDB()
	contentType := utils.GetContentType(c)
	_, _ = db, contentType
	books := []models.Book{}
	if contentType == appJSON {
		c.ShouldBindJSON(&books)
	} else {
		c.ShouldBind(&books)
	}
	err := db.Find(&books).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    books,
	})
}

func GetBookById(c *gin.Context) {
	db := database.GetDB()
	contentType := utils.GetContentType(c)
	_, _ = db, contentType
	books := models.Book{}
	if contentType == appJSON {
		c.ShouldBindJSON(&books)
	} else {
		c.ShouldBind(&books)
	}
	id := c.Param("id")
	err := db.Where("id = ?", id).First(&books).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    books,
	})
}

func UpdateBook(c *gin.Context) {
	db := database.GetDB()
	contentType := utils.GetContentType(c)
	_, _ = db, contentType
	books := models.Book{}
	if err := db.Where("id = ?", c.Param("id")).First(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&books); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Save(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Success",
		"data":    books,
	})
}

func DeleteBook(c *gin.Context) {
	db := database.GetDB()
	contentType := utils.GetContentType(c)
	_, _ = db, contentType
	books := models.Book{}
	if err := db.Where("id = ?", c.Param("id")).First(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Delete(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
		"data":    books,
	})
}
