package usercase

import (
	"context"
	"todo-list/app/user/domain"
)

type UserRepository interface {
	CreateUser(context.Context, *domain.User) (int32, error)
	GetById(context.Context, int32) (*domain.User, error)
	GetByUsername(context.Context, string) (*domain.User, error)
	UpdateUsername(context.Context, int32, string) error
}

type UserUseCase struct {
	repo UserRepository
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) Register(ctx context.Context, user *domain.User) (int32, error) {
	 return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUseCase) GetById(ctx context.Context, userId int32) (*domain.User, error) {
	return uc.repo.GetById(ctx, userId)
}
func (uc *UserUseCase) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	return uc.repo.GetByUsername(ctx, username)
}

func (uc *UserUseCase) UpdateUsername(ctx context.Context, userId int32, username string) error {
	return uc.repo.UpdateUsername(ctx, userId, username)
}