package blog

import (
	"dogrod-web-service/global"
	"dogrod-web-service/model/blog"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlogPostService struct{}

// initialize db
func (blogPostService *BlogPostService) InitDatabase() {
	global.ConnectDatabase().AutoMigrate(&blog.BlogPost{})
}

// create blog post
func (blogPostService *BlogPostService) AddBlogPost(c *gin.Context) {
	var newPost blog.BlogPost

	if err := c.BindJSON(&newPost); err != nil {
		panic("An error occurred when parse request data")
	}

	// TODO prevent duplicate slug
	db := global.ConnectDatabase()

	var duplicatePost blog.BlogPost

	db.Where("slug = ?", newPost.Slug).First(&duplicatePost)

	if duplicatePost.ID != 0 {
		// panic("slug already exist")
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "slug already exist"})
		return
	}

	db.Create(&newPost)

	c.IndentedJSON(http.StatusOK, newPost)
}

// retrieve all blog posts
func (blogPostService *BlogPostService) GetBlogPosts(c *gin.Context) {
	var posts []blog.BlogPost

	result := global.ConnectDatabase().Find(&posts)

	if result.Error != nil {
		panic("failed to query blog posts")
	}

	c.IndentedJSON(http.StatusOK, posts)
}

// retrieve blog post via post slug
func (blogPostService *BlogPostService) GetBlogPostBySlug(c *gin.Context) {
	var post blog.BlogPost

	slug := c.Param("slug")

	db := global.ConnectDatabase()

	err := db.Where("slug = ?", slug).First(&post).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, post)
}
