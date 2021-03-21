package routes

import (
	"rest-api-test/controller"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	//grouped the links 
	articleGrp := r.Group("/article-api")
	{
		articleGrp.GET("article", controller.GetArticles)
		articleGrp.POST("article", controller.CreateArticle)
		articleGrp.GET("article/:id", controller.GetArticleByID)
		articleGrp.PUT("article/:id", controller.UpdateArticle)
		articleGrp.DELETE("article/:id", controller.DeleteArticle)
	}
	return r
}
