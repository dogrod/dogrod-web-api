package blog

import (
	"dogrod-web-service/controllers"

	"github.com/gin-gonic/gin"
)

type BlogPostRouter struct{}

func (s *BlogPostRouter) InitBlogPostRouter(Router *gin.RouterGroup) {
	blogPostRouter := Router.Group("blog/posts")
	blogPostController := controllers.ControllerGroupApp.BlogControllerGroup.BlogPostController

	blogPostRouter.GET("", blogPostController.GetBlogPosts)
	blogPostRouter.GET(":slug", blogPostController.GetBlogPostBySlug)
	blogPostRouter.POST("", blogPostController.CreatePost)
}
