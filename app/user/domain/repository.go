package domain

import "context"

type UserRepository interface {
    CreateUser(context.Context, *User) (int64, error)
    GetById(context.Context, string) (*User, error)
}