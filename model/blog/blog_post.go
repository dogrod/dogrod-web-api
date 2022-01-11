package blog

import (
	"database/sql"
	"dogrod-web-service/global"
)

type BlogPost struct {
	global.BaseModel
	Title       string       `json:"title" gorm:"comment:blog title"`
	Slug        string       `json:"slug" gorm:"comment:unique identifying part of blog post;index"`
	CoverImg    string       `json:"coverImg" gorm:"comment:url of cover image"`
	Content     string       `json:"content" gorm:"type:text;comment:blog content"`
	AuthorId    uint         `json:"authorId" gorm:"comment:author id;default:0"`
	CategoryId  uint         `json:"categoryId" gorm:"comment:post category id;default:0"`
	PublishedAt sql.NullTime `json:"publishedAt" gorm:"comment:time that post published"`
}
