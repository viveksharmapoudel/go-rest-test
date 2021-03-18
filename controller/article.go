package controller

import (
	"rest-api-test/models"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

//GetArticles ... Get all Articles
func GetArticles(c *gin.Context) {
	var articles []models.Article
	err := models.GetAllArticles(&articles)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, articles)
	}
}

//CreateArticle ... Create Article
func CreateArticle(c *gin.Context) {
	var article models.Article
	c.BindJSON(&article)
	err := models.CreateArticle(&article)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, article)
	}
}

//GetArticleByID ... Get the article by id
func GetArticleByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var article models.Article
	err := models.GetArticleByID(&article, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, article)
	}
}

//UpdateArticle ... Update the article information
func UpdateArticle(c *gin.Context) {
	var article models.Article
	id := c.Params.ByName("id")
	err := models.GetArticleByID(&article, id)
	if err != nil {
		c.JSON(http.StatusNotFound, article)
	}
	c.BindJSON(&article)
	err = models.UpdateArticle(&article, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, article)
	}
}

//DeleteArticle ... Delete the Article
func DeleteArticle(c *gin.Context) {
	var article models.Article
	id := c.Params.ByName("id")
	err := models.DeleteArticle(&article, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

