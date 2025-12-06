package pack

import (
	"todo-list/app/user/domain"
	"todo-list/kitex_gen/model"
)

func NewBadResp(err error) *model.RespBase {
    return &model.RespBase{
        Code:    400,
        Message: err.Error(),
    }
}

func NewGoodResp() *model.RespBase {
    return &model.RespBase{
        Code:    200,
        Message: "success",
    }
}

func NewUserInfo(user *domain.User) *model.UserInfo {
    return &model.UserInfo{
        UserId:   user.UserId,
        Name: user.Username,
    }
}
