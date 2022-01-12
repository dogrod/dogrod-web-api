package blog

import (
	"dogrod-web-service/global"
	"dogrod-web-service/model/blog"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type BlogPostService struct{}

// initialize db
func (blogPostService *BlogPostService) InitDatabase() {
	global.ConnectDatabase().AutoMigrate(&blog.BlogPost{})
}

// create blog post
func (blogPostService *BlogPostService) CreateBlogPost(newPost blog.BlogPost) (err error) {
	// TODO prevent duplicate slug
	db := global.ConnectDatabase()

	var duplicatePost blog.BlogPost

	dupErr := db.Where("slug = ?", newPost.Slug).First(&duplicatePost).Error

	if !errors.Is(dupErr, gorm.ErrRecordNotFound) {
		err = fmt.Errorf("slug already exist")
	} else {
		err = db.Create(&newPost).Error
	}

	return err
}

// retrieve all blog posts
func (blogPostService *BlogPostService) GetBlogPosts() (err error, list []blog.BlogPost) {
	var posts []blog.BlogPost

	db := global.ConnectDatabase()

	err = db.Find(&posts).Error

	return err, posts
}

// retrieve blog post via post slug
func (blogPostService *BlogPostService) GetBlogPostBySlug(slug string) (err error, post blog.BlogPost) {
	db := global.ConnectDatabase()

	err = db.Where("slug = ?", slug).First(&post).Error

	return err, post
}
