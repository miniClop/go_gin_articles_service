package main

import (
	"github.com/gin-gonic/gin"
	"go_gin_articles_service/handlers"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", handlers.ShowIndexPage)

	router.GET("/article/view/:article_id", handlers.Article)

	_ = router.Run()
}
