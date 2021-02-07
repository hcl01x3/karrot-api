package model

import "time"

type Todo struct {
	Id        int64     `xorm:"pk autoincr" json:"id"`
	Text      string    `xorm:"notnull text" json:"text"`
	AuthorId  int64     `xorm:"notnull index" json:"authorId"`
	CreatedAt time.Time `xorm:"notnull created" json:"createdAt"`
}

func (Todo) TableName() string {
	return "todos"
}

type NewTodoInput struct {
	Text string `json:"text"`
}

type UpdateTodoInput struct {
	Text string `json:"text"`
}
