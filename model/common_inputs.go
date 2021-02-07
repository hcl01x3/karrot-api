package model

type PagingInput struct {
	Cursor *int64 `validate:"omitempty,gte=1,lte=9223372036854775807" json:"cursor"`
	After  int    `validate:"gte=1,lte=100" json:"after"`
}
