package blog

import (
	blog_post "dogrod-web-service/model/blog"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type BlogPostService struct{}

func (blogPostService *BlogPostService) ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

// initialize db
func (blogPostService *BlogPostService) InitDatabase() {
	blogPostService.ConnectDatabase().AutoMigrate(&blog_post.BlogPost{})
}

// create blog post
func (blogPostService *BlogPostService) AddBlogPost(c *gin.Context) {
	var newPost blog_post.BlogPost

	if err := c.BindJSON(&newPost); err != nil {
		panic("An error occurred when parse request data")
	}

	blogPostService.ConnectDatabase().Create(&newPost)

	c.IndentedJSON(http.StatusOK, newPost)
}

// retrieve all blog posts
func (blogPostService *BlogPostService) GetBlogPosts(c *gin.Context) {
	var posts []blog_post.BlogPost

	result := blogPostService.ConnectDatabase().Find(&posts)

	if result.Error != nil {
		panic("failed to query blog posts")
	}

	c.IndentedJSON(http.StatusOK, posts)
}

// retrieve blog post via post slug
func (blogPostService *BlogPostService) GetBlogPostBySlug(c *gin.Context) {
	var post blog_post.BlogPost

	slug := c.Param("slug")

	db := blogPostService.ConnectDatabase()

	err := db.Where("slug = ?", slug).First(&post).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, post)
}
