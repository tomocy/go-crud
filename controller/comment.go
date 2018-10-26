package controller

import "time"

type Comment struct {
	ID        int
	UserID    int
	PostID    int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
