package rpc

import (
	"context"
	"todo-list/app/user/controllers/rpc/pack"
	"todo-list/app/user/domain"
	"todo-list/app/user/usercase"
	user "todo-list/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{
	usecase *usercase.UserUseCase
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = &user.RegisterResponse{}
	userID, err := s.usecase.Register(ctx, &domain.User{
		Username: req.Name,
		Password: req.Password,
	})

	if err != nil {
		resp.Base = pack.NewBadResp(err)
		return 
	}
	resp.Base = pack.NewGoodResp()
	resp.UserId = userID
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = &user.LoginResponse{}
	user, err := s.usecase.Login(ctx, req.Name, req.Password)
	if err != nil {
		resp.Base = pack.NewBadResp(err)
		return
	}
	resp.Base = pack.NewGoodResp()
	resp.Info = pack.NewUserInfo(user)
	return
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	resp = &user.GetUserResponse{}
	user, err := s.usecase.GetById(ctx, req.UserId)
	if err != nil {
		resp.Base = pack.NewBadResp(err)
		return
	}
	resp.Base = pack.NewGoodResp()
	resp.Info = pack.NewUserInfo(user)
	return
}

// Updateusername implements the UserServiceImpl interface.
func (s *UserServiceImpl) Updateusername(ctx context.Context, req *user.UpdateusernameRequest) (resp *user.UpdateusernameResponse, err error) {
	resp = &user.UpdateusernameResponse{}
	err = s.usecase.UpdateUsername(ctx, req.UserId, req.Username)
	if err != nil {
		resp.Base = pack.NewBadResp(err)
		return
	}
	resp.Base = pack.NewGoodResp()
	return
}
