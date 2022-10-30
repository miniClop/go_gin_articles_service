package models

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article Content 1"},
	{ID: 2, Title: "Article 2", Content: "Article Content 2"},
}

// AllArticles Получение списка артикулов
func AllArticles() []article {
	return articleList
}

// ArticleById Получение артикула по ID
func ArticleById(id int) (*article, error) {
	for _, a := range AllArticles() {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("article not found")
}
