package blog

import (
	"dogrod-web-service/global"
	"dogrod-web-service/model/blog"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	if newCategory.Slug == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Slug is required"})
		return
	}

	db := global.ConnectDatabase()

	// check if slug is duplicated
	var duplicateCategory blog.BlogCategory
	err := db.Where("slug = ?", newCategory.Slug).First(&duplicateCategory).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Slug is already in use"})
		return
	}

	db.Create(&newCategory)

	// generate slug if slug not provide
	// if newCategory.Slug == "" {
	// 	newCategory.Slug = strconv.FormatUint(uint64(newCategory.ID), 10)

	// 	db.Save(&newCategory)
	// }

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
