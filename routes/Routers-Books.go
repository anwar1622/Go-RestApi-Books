package routes

import (
	"Go-RestApi-Books/controllers"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()

	bookRouter := router.Group("/api/v1")
	{
		bookRouter.POST("/books", controllers.CreateBooks)
		bookRouter.GET("/books", controllers.GetBook)
		bookRouter.GET("/books/:id", controllers.GetBookById)
		bookRouter.PUT("/books/:id", controllers.UpdateBook)
		bookRouter.DELETE("/books/:id", controllers.DeleteBook)
	}

	return router
}
