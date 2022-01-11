package blog

import (
	"dogrod-web-service/controllers"
	"dogrod-web-service/service"

	"github.com/gin-gonic/gin"
)

type BlogPostRouter struct{}

func (s *BlogPostRouter) InitBlogPostRouter(Router *gin.RouterGroup) {
	blogPostRouter := Router.Group("blog/posts")
	blogPostService := service.ServiceGroupApp.BlogServiceGroup.BlogPostService
	blogPostController := controllers.ControllerGroupApp.BlogControllerGroup.BlogPostController

	blogPostRouter.GET("", blogPostService.GetBlogPosts)
	blogPostRouter.GET(":slug", blogPostService.GetBlogPostBySlug)
	blogPostRouter.POST("", blogPostController.CreatePost)
}
