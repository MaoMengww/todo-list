package usercase

import (
	"context"
	"errors"
	"todo-list/app/user/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Register(ctx context.Context, user *domain.User) (int64, error) {
	existUser, err := uc.repo.GetByUsername(ctx, user.Username)
	if err == nil && existUser != nil {
		return 0, errors.New("用户已存在")
	}

	//密码加密储存
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(hashedPassword)
	return uc.repo.Create(ctx, user)
}

func (uc *UserUseCase) Login(ctx context.Context, username string, password string) (*domain.User, error) {
	user, err := uc.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("密码错误")
	}
	return uc.repo.GetByUsername(ctx, username)
}

func (uc *UserUseCase) GetById(ctx context.Context, userId int64) (*domain.User, error) {
	return uc.repo.GetById(ctx, userId)
}
func (uc *UserUseCase) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	return uc.repo.GetByUsername(ctx, username)
}

func (uc *UserUseCase) UpdateUsername(ctx context.Context, userId int64, username string) error {
	return uc.repo.UpdateUsername(ctx, userId, username)
}
