package infrastructure

import (
	"context"
	"todo-list/app/user/domain"

	"gorm.io/gorm"
)

type UserModel struct {
	Id       int32 `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"uniqueIndex;size:255"`
	Password string  `gorm:"size:255"`
}

type MysqlUserRepository struct {
	DB *gorm.DB
}

func NewMysqlUserRepository(db *gorm.DB) *MysqlUserRepository {
	return &MysqlUserRepository{
		DB: db,
	}
}

func (r *MysqlUserRepository) Create(ctx context.Context, user *domain.User) error {
	return r.DB.Create(user).Error
}


func (r *MysqlUserRepository) GetById(ctx context.Context, userId string) (*domain.User, error) {
	var user domain.User
	err := r.DB.Where("user_id = ?", userId).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MysqlUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MysqlUserRepository) UpdateUsername(ctx context.Context, uid int64, username string) error {
    return r.DB.WithContext(ctx).
        Model(&UserModel{}).
        Where("id = ?", uid).
        Update("username", username).Error
}