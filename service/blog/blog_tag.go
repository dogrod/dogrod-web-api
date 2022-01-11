package blog

import (
	"dogrod-web-service/global"
	"dogrod-web-service/model/blog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogTagService struct{}

func (blogTagService *BlogTagService) InitDatabase() {
	global.ConnectDatabase().AutoMigrate(&blog.BlogTag{})
}

// create blog tag
func (blogTagService *BlogTagService) AddBlogTag(c *gin.Context) {
	var newTag blog.BlogTag

	if err := c.BindJSON(&newTag); err != nil {
		panic("An error occurred when parse request data")
	}

	db := global.ConnectDatabase()
	db.Create(&newTag)

	c.IndentedJSON(http.StatusOK, newTag)
}

// retrieve all blog tags
func (blogTagService *BlogTagService) GetBlogTags(c *gin.Context) {
	var tags []blog.BlogTag

	err := global.ConnectDatabase().Find(&tags).Error

	if err != nil {
		panic("failed to query blog tags")
	}

	c.IndentedJSON(http.StatusOK, tags)
}
