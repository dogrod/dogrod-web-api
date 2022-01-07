package blog

import "dogrod-web-service/global"

type BlogTag struct {
	global.BaseModel
	Name string `json:"name" gorm:"comment:slug name"`
	Slug string `json:"slug" gorm:"comment:unique identifying part of tag"`
}
