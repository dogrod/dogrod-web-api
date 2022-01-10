package blog

import (
	"dogrod-web-service/global"
	"dogrod-web-service/model/blog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogCategoryService struct{}

func (blogCategoryService *BlogCategoryService) InitDatabase() {
	global.ConnectDatabase().AutoMigrate(&blog.BlogCategory{})
}

// create blog category item
func (blogCategoryService *BlogCategoryService) AddBlogCategory(c *gin.Context) {
	var newCategory blog.BlogCategory

	if err := c.BindJSON(&newCategory); err != nil {
		panic("An error occurred when parse request data")
	}

	// TODO generate slug if slug not provide
	// TODO check if slug duplicated

	global.ConnectDatabase().Create(&newCategory)

	c.IndentedJSON(http.StatusOK, newCategory)
}

// retrieve all blog categories
func (blogCategoryService *BlogCategoryService) GetBlogCategories(c *gin.Context) {
	var categories []blog.BlogCategory

	err := global.ConnectDatabase().Find(&categories).Error

	if err != nil {
		panic("failed to query blog posts")
	}

	c.IndentedJSON(http.StatusOK, categories)
}
