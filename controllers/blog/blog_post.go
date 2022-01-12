package blog

import (
	"dogrod-web-service/model/blog"
	"dogrod-web-service/service"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlogPostController struct{}

var blogPostService = service.ServiceGroupApp.BlogServiceGroup.BlogPostService

// create blog post
func (blogPostController *BlogPostController) CreatePost(c *gin.Context) {
	var post blog.BlogPost
	_ = c.ShouldBindJSON(&post)

	if err := blogPostService.CreateBlogPost(post); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error"})
	} else {
		c.IndentedJSON(http.StatusOK, post)
	}
}

// retrieve all blog posts
func (blogPostController *BlogPostController) GetBlogPosts(c *gin.Context) {
	err, posts := blogPostService.GetBlogPosts()

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No post found"})
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error"})
	} else {
		c.IndentedJSON(http.StatusOK, posts)
	}
}

// retrieve specific post via slug
func (blogPostController *BlogPostController) GetBlogPostBySlug(c *gin.Context) {
	slug := c.Param("slug")

	err, post := blogPostService.GetBlogPostBySlug(slug)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No post found"})
	} else if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error"})
	} else {
		c.IndentedJSON(http.StatusOK, post)
	}
}
