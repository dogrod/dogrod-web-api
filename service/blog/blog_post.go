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
func (blogPostService *BlogPostService) CreateBlogPost(newPost blog.BlogPost) (err error) {
	// var newPost blog.BlogPost

	// if err := c.BindJSON(&newPost); err != nil {
	// 	panic("An error occurred when parse request data")
	// }

	// TODO prevent duplicate slug
	db := global.ConnectDatabase()

	var duplicatePost blog.BlogPost

	dupErr := db.Where("slug = ?", newPost.Slug).First(&duplicatePost).Error

	if !errors.Is(dupErr, gorm.ErrRecordNotFound) {
		err = errors.New("slug already exist")
	} else {
		err = db.Create(&newPost).Error
	}

	return err
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
