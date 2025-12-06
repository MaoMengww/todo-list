package infrastructure

import (
	"context"
	"todo-list/app/user/domain"

	"gorm.io/gorm"
)

type UserModel struct {
	Id       int64 `gorm:"primaryKey;autoIncrement"`
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

func (u UserModel) TableName() string {
	return "users"
}

func (r *MysqlUserRepository) Create(ctx context.Context, duser *domain.User) (int64, error) {
	user := &UserModel{
		Username: duser.Username,
		Password: duser.Password,
		Id:   duser.UserId,
	}
	result := r.DB.Create(user)
	return user.Id, result.Error
}


func (r *MysqlUserRepository) GetById(ctx context.Context, userId int64) (*domain.User, error) {
    var userModel UserModel
    err := r.DB.WithContext(ctx).Where("id = ?", userId).First(&userModel).Error
    if err != nil {
        return nil, err
    }
    return &domain.User{
        UserId:   int64(userModel.Id),
        Username: userModel.Username,
        Password: userModel.Password,
    }, nil
}

func (r *MysqlUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user UserModel
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &domain.User{
		UserId:   int64(user.Id),
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func (r *MysqlUserRepository) UpdateUsername(ctx context.Context, uid int64, username string) error {
    return r.DB.WithContext(ctx).
        Model(&UserModel{}).
        Where("id = ?", uid).
        Update("username", username).Error
}