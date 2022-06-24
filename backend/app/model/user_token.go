package model

import (
	"github.com/google/uuid"
	"time"
)

type UserToken struct {
	BaseModel
	UserID   uuid.UUID `json:"user_id" gorm:"column:user_id; type:uuid; NOT NULL"`
	User     User      `json:"-" gorm:"foreignkey:user_id;references:id"`
	Token    string    `json:"token" gorm:"column:token; type:varchar(255); NOT NULL"`
	ExpireAt time.Time `json:"expire_at" gorm:"column:expire_at; type:timestamp; NOT NULL"`
}

func (r *UserToken) TableName() string {
	return "user_token"
}
