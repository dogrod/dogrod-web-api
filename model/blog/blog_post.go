package blog

import (
	"dogrod-web-service/global"
)

type BlogPost struct {
	global.BaseModel
	Title string `json:"title" gorm:"comment:blog title"`
	Slug  string `json:"slug" gorm:"comment:unique identifying part of blog post"`
	// Author
	Content string `json:"content" gorm:"comment:blog content"`
	// PublishAt
	// UpdateAt
	// Status
	// Tags
	// Category
	CoverImg string `json:"coverImg" gorm:"comment:url of cover image"`
}
