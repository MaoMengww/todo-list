package domain

import "time"

type Todo struct {
	TodoId      int64  
	UserId      int64  
	Title       string 
	Content     string 
	Completed   bool   
	CreatedAt   time.Time
	DiedAt   time.Time
	Priority    int64
}