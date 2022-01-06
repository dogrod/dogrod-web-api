package blog

import (
	"dogrod-web-service/service"

	"github.com/gin-gonic/gin"
)

type BlogPostRouter struct{}

func (s *BlogPostRouter) InitBlogPostRouter(Router *gin.RouterGroup) {
	blogRouter := Router.Group("blog")
	blogPostService := service.ServiceGroupApp.BlogServiceGroup.BlogPostService

	blogRouter.GET("posts", blogPostService.GetBlogPosts)
	blogRouter.GET("posts/:slug", blogPostService.GetBlogPostBySlug)
	blogRouter.POST("posts", blogPostService.AddBlogPost)
}
