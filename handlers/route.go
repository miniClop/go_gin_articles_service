package handlers

import (
	"github.com/gin-gonic/gin"
	"go_gin_articles_service/models"
	"net/http"
	"strconv"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.AllArticles()

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

func Article(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		panic("article not passed")
	}

	article, err := models.ArticleById(articleId)
	if err != nil {
		panic(err)
	}

	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"article.html",
		// Pass the data that the page uses
		gin.H{
			"title":   article.Title,
			"payload": article,
		},
	)
}
