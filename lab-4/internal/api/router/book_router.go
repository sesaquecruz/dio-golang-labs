package router

import (
	"lab-4/internal/entity"

	"github.com/gin-gonic/gin"
)

func StartBookRouter(addr string, c entity.BookControllerInterface) {
	router := gin.Default()
	router.GET("books/", c.FindAll)
	router.GET("books/:id", c.FindById)
	router.POST("books/", c.Save)
	router.PUT("books/:id", c.Update)
	router.DELETE("books/:id", c.Delete)
	router.Run(addr)
}
