package global

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey;comment:unique id"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:time that entity created"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:time that entity updated"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
