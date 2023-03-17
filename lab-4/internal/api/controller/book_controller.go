package controller

import (
	"lab-4/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	R entity.BookRepositoryInterface
}

func NewBookController(r entity.BookRepositoryInterface) *BookController {
	return &BookController{R: r}
}

func (c *BookController) FindAll(g *gin.Context) {
	books, err := c.R.FindAll()
	if err != nil {
		g.Status(http.StatusInternalServerError)
		return
	}
	g.JSON(http.StatusOK, books)
}

func (c *BookController) FindById(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.Status(http.StatusBadRequest)
		return
	}
	book, err := c.R.FindById(int64(id))
	if err != nil {
		g.Status(http.StatusNotFound)
		return
	}
	g.JSON(http.StatusOK, book)
}

func (c *BookController) Save(g *gin.Context) {
	var book entity.Book
	err := g.BindJSON(&book)
	if err != nil {
		g.Status(http.StatusBadRequest)
		return
	}
	id, err := c.R.Save(&book)
	if err != nil {
		g.Status(http.StatusBadRequest)
		return
	}
	book.ID = id
	g.JSON(http.StatusCreated, book)
}

func (c *BookController) Update(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.Status(http.StatusBadRequest)
		return
	}
	var book entity.Book
	err = g.BindJSON(&book)
	if err != nil {
		g.Status(http.StatusBadRequest)
		return
	}
	book.ID = int64(id)
	row, err := c.R.Update(book.ID, &book)
	if err != nil || row == 0 {
		g.Status(http.StatusBadRequest)
		return
	}
	g.JSON(http.StatusOK, book)
}

func (c *BookController) Delete(g *gin.Context) {
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		g.Status(http.StatusBadRequest)
		return
	}
	row, err := c.R.Delete(int64(id))
	if err != nil || row == 0 {
		g.Status(http.StatusBadRequest)
		return
	}
	g.Status(http.StatusNoContent)
}
