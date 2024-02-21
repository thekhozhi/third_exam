package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	_ "develop/api/docs"
	"develop/api/handler"

	"develop/service"
	"time"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func New(services service.IServiceManager) *gin.Engine {
	h := handler.New(services)

	r := gin.New()
	
	r.Use(traceRequest)

	{
		 
		r.POST("/book", h.CreateBook)
		r.GET("/book/:id", h.GetBook)
		r.GET("/books", h.GetBookList)
		r.PUT("/book/:id", h.UpdateBook)
		r.DELETE("/book/:id", h.DeleteBook)
		r.PATCH("/book/:id", h.UpdateBookPageNumber)


		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return r
}


func traceRequest(c *gin.Context) {
	beforeRequest(c)

	c.Next()

	afterRequest(c)
}

func beforeRequest(c *gin.Context) {
	startTime := time.Now()

	c.Set("start_time", startTime)

	log.Println("start time:", startTime.Format("2006-01-02 15:04:05.000 MST"), "path:", c.Request.URL.Path)
}

func afterRequest(c *gin.Context) {
	startTime, exists := c.Get("start_time")
	if !exists {
		startTime = time.Now()
	}

	duration := time.Since(startTime.(time.Time))

	log.Printf("end time: %s, duration: %d milliseconds, method: %s, status: %d\n", time.Now().Format("2006-01-02 15:04:05.000 MST"), duration.Milliseconds(), c.Request.Method, c.Writer.Status())
	fmt.Println()
}
