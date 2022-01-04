package main

import (
	blog_post "dogrod-web-service/model/blog"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// album slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	getDB().AutoMigrate(&blog_post.BlogPost{})

	router.GET("/blog/posts", getBlogPosts)
	router.GET("/blog/posts/:slug", getBlogPostBySlug)
	router.POST("/blog/posts", addBlogPost)

	router.Run("0.0.0.0:9009")
}

func getDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

// create blog post
func addBlogPost(c *gin.Context) {
	var newPost blog_post.BlogPost

	if err := c.BindJSON(&newPost); err != nil {
		panic("An error ocurred when parse request data")
	}

	getDB().Create(&newPost)

	c.IndentedJSON(http.StatusOK, newPost)
}

// retrive all blog posts
func getBlogPosts(c *gin.Context) {
	var posts []blog_post.BlogPost

	result := getDB().Find(&posts)

	if result.Error != nil {
		panic("failed to query blog posts")
	}

	c.IndentedJSON(http.StatusOK, posts)
}

func getBlogPostBySlug(c *gin.Context) {
	var post blog_post.BlogPost

	slug := c.Param("slug")

	db := getDB()

	err := db.Where("slug = ?", slug).First(&post).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "post not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, post)
}
