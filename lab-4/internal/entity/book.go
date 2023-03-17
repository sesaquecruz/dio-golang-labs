package entity

import "github.com/gin-gonic/gin"

type Book struct {
	ID       int64  `json:"id,omitempty"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Year     string `json:"year"`
}

type BookRepositoryInterface interface {
	FindAll() ([]Book, error)
	FindById(id int64) (*Book, error)
	Save(book *Book) (int64, error)
	Update(id int64, book *Book) (int64, error)
	Delete(id int64) (int64, error)
}

type BookControllerInterface interface {
	FindAll(g *gin.Context)
	FindById(g *gin.Context)
	Save(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
}
