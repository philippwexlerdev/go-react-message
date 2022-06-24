package model

import "github.com/google/uuid"

type PostResponse struct {
	BaseModel
	Message string    `json:"message" gorm:"column:message; type:text"`
	UserID  uuid.UUID `json:"user_id" gorm:"column:user_id; type:uuid; NOT NULL"`
	PostID  uuid.UUID `json:"post_id" gorm:"column:post_id; type:uuid; NOT NULL"`
	User    User      `json:"-" gorm:"foreignkey:user_id;references:id"`
}

type PostResponseCreate struct {
	Message *string    `json:"message"`
	UserID  *uuid.UUID `json:"user_id"`
	PostID  *uuid.UUID `json:"post_id"`
}

type PostResponseUpdate struct {
	Message *string `json:"message"`
}

func (r *PostResponse) TableName() string {
	return "post_response"
}
