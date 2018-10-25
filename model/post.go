package model

import "time"

type Post struct {
	ID        int
	UserID    int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
