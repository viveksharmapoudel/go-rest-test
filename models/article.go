package models

import (
	"rest-api-test/config"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//GetAllarticles Fetch all article data
func GetAllArticles(article *[]Article) (err error) {
	if err = config.DB.Find(article).Error; err != nil {
		return err
	}
	return nil
}

//Createarticle ... Insert New data
func CreateArticle(article *Article) (err error) {
	fmt.Println("check:", article)

	if err = config.DB.Create(article).Error; err != nil {
		return err
	}
	return nil
}

//GetarticleByID ... Fetch only one article by Id
func GetArticleByID(article *Article, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(article).Error; err != nil {
		return err
	}
	return nil
}

//Updatearticle ... Update article
func UpdateArticle(article *Article, id string) (err error) {
	fmt.Println(article)
	config.DB.Save(article)
	return nil
}

//Deletearticle ... Delete article
func DeleteArticle(article *Article, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(article)
	return nil
}
