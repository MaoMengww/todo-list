package domain

import "context"


type UserRepository interface {
	CreateUser(context.Context, *User) (int64, error)
	GetById(context.Context, int64) (*User, error)
	GetByUsername(context.Context, string) (*User, error)
	UpdateUsername(context.Context, int64, string) error
}