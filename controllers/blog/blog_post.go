package blog

import (
	"dogrod-web-service/model/blog"
	"dogrod-web-service/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogPostController struct{}

var blogPostService = service.ServiceGroupApp.BlogServiceGroup.BlogPostService

func (blogPostController *BlogPostController) CreatePost(c *gin.Context) {
	var post blog.BlogPost
	_ = c.ShouldBindJSON(&post)

	if err := blogPostService.CreateBlogPost(post); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal error"})
	} else {
		c.IndentedJSON(http.StatusOK, post)
	}
}
