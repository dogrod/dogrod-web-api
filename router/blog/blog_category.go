package blog

import (
	"dogrod-web-service/service"

	"github.com/gin-gonic/gin"
)

type BlogCategoryRouter struct{}

func (s *BlogCategoryRouter) InitBlogCategoryRouter(Router *gin.RouterGroup) {
	blogCategoryRouter := Router.Group("blog/categories")
	blogCategoryService := service.ServiceGroupApp.BlogServiceGroup.BlogCategoryService

	blogCategoryRouter.GET("", blogCategoryService.GetBlogCategories)
	blogCategoryRouter.POST("", blogCategoryService.AddBlogCategory)
}
