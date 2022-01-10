package main

import (
	"dogrod-web-service/router"
	"dogrod-web-service/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Init all databases
	blogPostService := service.ServiceGroupApp.BlogServiceGroup.BlogPostService
	blogPostService.InitDatabase()

	blogCategoryService := service.ServiceGroupApp.BlogServiceGroup.BlogCategoryService
	blogCategoryService.InitDatabase()

	// Init routers
	Router := gin.Default()
	RouterGroup := Router.Group("")

	router.RouterGroupApp.Blog.InitBlogPostRouter(RouterGroup)
	router.RouterGroupApp.Blog.InitBlogCategoryRouter(RouterGroup)

	Router.Run("0.0.0.0:9009")
}
