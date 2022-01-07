package blog

import "dogrod-web-service/global"

type BlogTaggedPost struct {
	global.BaseModel
	ContentObjectId uint `json:"contentObjectId" gorm:"comment:db id of content object"`
	TagId           uint `json:"tagId" gorm:"comment:db id of tag"`
}
