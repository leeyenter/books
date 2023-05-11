package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leeyenter/books/backend/auth"
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
	config.AllowOrigins = []string{"http://localhost:5173", "https://ytbooks.pages.dev"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"auth", "Content-Type"}
	r.Use(cors.New(config))

	authApi := r.Group("/auth")
	authApi.GET("/", auth.Middleware, checkLogin)
	authRegis := authApi.Group("/registration")
	authRegis.GET("/", getNumCredentials)
	authRegis.POST("/begin", beginRegistration)
	authRegis.POST("/finish", finishRegistration)
	authLogin := authApi.Group("/login")
	authLogin.POST("/begin", beginLogin)
	authLogin.POST("/finish", finishLogin)

	book := r.Group("/book")
	book.GET("/", getBooks)
	book.POST("/", auth.Middleware, addBook)

	singleBookApi := book.Group("/:id")
	singleBookApi.Use(auth.Middleware)
	singleBookApi.POST("", editBook)
	singleBookApi.POST("/bought", markBookAsBought)
	singleBookApi.POST("/price", setBookPrice)
	singleBookApi.DELETE("/price", removeBookPrice)
	singleBookApi.PUT("/status", setBookStatus)

	return r
}
