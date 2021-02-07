package model

import "time"

type User struct {
	Id        int64     `xorm:"pk autoincr" json:"id"`
	Name      string    `xorm:"varchar(10) notnull unique" json:"name"`
	Mobile    string    `xorm:"varchar(11) notnull unique" json:"mobile"`
	Password  string    `xorm:"text notnull" json:"-"`
	Role      Role      `xorm:"notnull" json:"role"`
	CreatedAt time.Time `xorm:"notnull created" json:"createdAt"`
}

func (User) TableName() string {
	return "users"
}

type NewUserInput struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	Name   *string `json:"name"`
	Mobile *string `json:"mobile"`
}

type UpdatePasswordInput struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type LogInInput struct {
	Mobile   string `validate:"numeric,min=4" json:"mobile"`
	Password string `validate:"min=3,max=20" json:"password"`
}
