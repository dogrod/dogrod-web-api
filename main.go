package main

import (
	"dogrod-web-service/router"
	"dogrod-web-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	blogPostService := service.ServiceGroupApp.BlogServiceGroup.BlogPostService
	blogPostService.InitDatabase()

	Router := gin.Default()
	RouterGroup := Router.Group("")

	router.RouterGroupApp.Blog.InitBlogPostRouter(RouterGroup)

	Router.Run("0.0.0.0:9009")
}
