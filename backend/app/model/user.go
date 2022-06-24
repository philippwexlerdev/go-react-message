package model

type User struct {
	BaseModel
	Username string `json:"username" gorm:"column:username; type:varchar(255); UNIQUE, NOT NULL"`
	Password string `json:"password" gorm:"column:password; type:text; NOT NULL"`
}

type UserCreateOrLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UseUpdate struct {
	Username *string `json:"username"`
	Token    *string `json:"token"`
	Password *string `json:"password"`
}

func (r *User) TableName() string {
	return "user"
}
