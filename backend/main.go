package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := setupRouter()
	err := r.Run(":5000")
	if err != nil {
		log.Fatalln(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"auth", "Content-Type"}
	r.Use(cors.New(config))

	book := r.Group("/book")
	book.GET("/", getBooks)
	book.POST("/", addBook)
	book.POST("/:id", editBook)
	book.POST("/:id/bought", markBookAsBought)
	book.POST("/:id/price", setBookPrice)
	book.DELETE("/:id/price", removeBookPrice)
	book.PUT("/:id/status", setBookStatus)

	return r
}
