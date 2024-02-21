package handler

import (
	"context"
	"database/sql"
	"develop/api/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
// @Router       /book [POST]
// @Summary      Creates a new book
// @Description  create a new book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        book body models.CreateBook false "book"
// @Success      201  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateBook(c *gin.Context)  {
	createBook := models.CreateBook{}
	
	err := c.ShouldBindJSON(&createBook)
	if err != nil{
		handleResponse(c,"Error in handlers, while binding json!",http.StatusBadRequest, err)
		return
	}

	book, err := h.services.Book().Create(context.Background(),createBook)
	if err != nil{
		handleResponse(c, "Error in handlers, while creating book!", http.StatusInternalServerError,err)
		return
	}

	handleResponse(c,"Success", http.StatusOK,book)
}

// GetBook godoc
// @Router       /book/{id} [GET]
// @Summary      Get book by id
// @Description  get book by id
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id path string true "book_id"
// @Success      201  {object}  models.Book
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBook(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handleResponse(c, "uuid is not valid!", http.StatusInternalServerError,errors.New("invaliid type uuid"))
		return
	}

	book, err := h.services.Book().GetByID(context.Background(),id)
	if err != nil{
		if errors.Is(err, sql.ErrNoRows){
		handleResponse(c,"Error in handlers, while getting book by id!",http.StatusInternalServerError,err)
		return
		}else{
			handleResponse(c,"We don't have a book you want!",http.StatusNotFound, errors.New("we don't have a book you want"))
		}
	}
	handleResponse(c, "",http.StatusOK, book)
}

// GetBookList godoc
// @Router       /books [GET]
// @Summary      Get book list
// @Description  get book list
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        page query string false "page"
// @Param        limit query string false "limit"
// @Param        search query string false "search"
// @Success      201  {object}  models.BooksResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetBookList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "Error in handlers, while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "Error in handlers, while converting limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	books, err := h.services.Book().GetList(context.Background(), models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	if err != nil {
		handleResponse(c, "Error in handlers, while getting list of books", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c,  "", http.StatusOK, books)
}

// UpdateBook godoc
// @Router       /book/{id} [PUT]
// @Summary      Update book
// @Description  update book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id path string true "book_id"
// @Param        book body models.UpdateBook false "book"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBook(c *gin.Context) {
	book := models.UpdateBook{}
	uid := c.Param("id")

	err := c.ShouldBindJSON(&book)
	if err != nil {
		handleResponse(c, "Error in hanlders, while reading body", http.StatusBadRequest, err.Error())
		return
	}

	book.ID = uid

	updBook, err := h.services.Book().Update(context.Background(), book)
	if err != nil {
		handleResponse(c, "Error in handlers, while updating book!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, updBook)
}

// DeleteBook godoc
// @Router       /book/{id} [DELETE]
// @Summary      Delete book
// @Description  delete book
// @Tags         book
// @Accept       json
// @Produce      json
// @Param        id path string true "book_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteBook(c *gin.Context) {
	uid := c.Param("id")

	err := h.services.Book().Delete(context.Background(), uid)
	if err != nil {
		handleResponse(c,  "Error in handlers, while deleting book!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Deleted!", http.StatusOK, "successfully deleted!")
}

// UpdateBookPageNumber godoc
// @Router       /book/{id} [PATCH]
// @Summary      Update book page number
// @Description  update book page number
// @Tags         book
// @Accept       json
// @Produce      json
// @Param 		 id path string true "book_id"
// @Param        book body models.UpdatePageNumber true "book"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateBookPageNumber(c * gin.Context) {
	updPageNum := models.UpdatePageNumber{}
	uid := c.Param("id")

	err := c.ShouldBindJSON(&updPageNum)
	if err != nil{
		handleResponse(c, "Error in handlers, while reading body!", http.StatusBadRequest,err)
		return
	}

	updPageNum.ID = uid

	book, err := h.services.Book().UpdatePageNumber(context.Background(),updPageNum)
	if err != nil{
		handleResponse(c,"Error in handlers, while updating page num!",http.StatusInternalServerError,err)
		return
	}

	handleResponse(c, "updated!", http.StatusOK, book)
}