package todo

import "time"

type Items struct {
	Id          int
	Title       string
	IsCompleted bool
	CreatedAt   time.Time
	CompletedAt time.Time
}
