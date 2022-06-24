package model

import "github.com/google/uuid"

type Post struct {
	BaseModel
	Message string    `json:"message" gorm:"column:message; type:text"`
	UserID  uuid.UUID `json:"user_id" gorm:"column:user_id; type:uuid; NOT NULL"`
	User    User      `json:"-" gorm:"foreignkey:user_id;references:id"`
}

type PostCreate struct {
	Message *string    `json:"message"`
	UserID  *uuid.UUID `json:"user_id"`
}

type PostUpdate struct {
	Message *string `json:"message" gorm:"column:message; type:text"`
}

func (r *Post) TableName() string {
	return "post"
}
