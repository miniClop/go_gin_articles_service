package handlers

import (
	"github.com/gin-gonic/gin"
	"go_gin_articles_service/models"
	"net/http"
	"strconv"
)

// ShowIndexPage Вид главной страницы
func ShowIndexPage(c *gin.Context) {
	articles := models.AllArticles()

	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

// Article Получение одного экземпляра артикула, по переданному параметру `article_id'
func Article(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("article_id"))
	if err != nil {
		panic("article not passed")
	}

	article, err := models.ArticleById(articleId)
	if err != nil {
		panic(err)
	}

	render(c, gin.H{
		"title":   article.Title,
		"payload": article,
	}, "article.html")
}

// render Вспомогательный хелпер для возвращения данных в необходимом формате переданного заголовка `Accept`
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}
