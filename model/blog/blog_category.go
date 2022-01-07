package blog

import "dogrod-web-service/global"

type BlogCategory struct {
	global.BaseModel
	Title string `json:"title" gorm:"comment:category title"`
	Slug  string `json:"slug" gorm:"comment:unique identifying part of blog category"`
}
