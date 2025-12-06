package main

import (
	"context"
	user "todo-list/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	// TODO: Your code here...
	return
}

// Updateusername implements the UserServiceImpl interface.
func (s *UserServiceImpl) Updateusername(ctx context.Context, req *user.UpdateusernameRequest) (resp *user.UpdateusernameResponse, err error) {
	// TODO: Your code here...
	return
}
