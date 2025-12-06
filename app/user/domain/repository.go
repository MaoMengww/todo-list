package domain

import "context"

//提供接口
type UserRepository interface {
	Create(context.Context, *User) (int64, error)
	GetById(context.Context, int64) (*User, error)
	GetByUsername(context.Context, string) (*User, error)
	UpdateUsername(context.Context, int64, string) error
}