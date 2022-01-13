package blog

import (
	"dogrod-web-service/global"
	"time"
)

type BlogPost struct {
	global.BaseModel
	Title       string     `json:"title" gorm:"comment:blog title"`
	Slug        string     `json:"slug" gorm:"comment:unique identifying part of blog post;index"`
	CoverImg    string     `json:"coverImg" gorm:"comment:url of cover image"`
	Summary     string     `json:"summary" gorm:"comment:summary of post content"`
	Content     string     `json:"content" gorm:"type:text;comment:blog content"`
	Status      string     `json:"status" gorm:"comment:draft/publish/private/trash;default:draft"`
	AuthorId    uint       `json:"authorId" gorm:"comment:author id;default:0"`
	CategoryId  uint       `json:"categoryId" gorm:"comment:post category id;default:0"`
	PublishedAt *time.Time `json:"publishedAt" gorm:"comment:time that post published"`
}
