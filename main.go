package main

import (
	"dogrod-web-service/router"

	"github.com/gin-gonic/gin"
)

func main() {
	Router := gin.Default()
	RouterGroup := Router.Group("")

	router.RouterGroupApp.Blog.InitBlogPostRouter(RouterGroup)

	Router.Run("0.0.0.0:9009")
}
