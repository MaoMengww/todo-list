package pack

import (
	"todo-list/app/todo/domain"
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

func NewTodoInfo(todo *domain.Todo) *model.TodoInfo {
	return &model.TodoInfo{
		Id:    todo.TodoId,
		UserId:    todo.UserId,
		Title:     todo.Title,
		Content:   todo.Content,
		Completed: todo.Completed,
		Priority:  todo.Priority,
		CreatedAt: todo.CreatedAt.Format("2006-01-02 15:04:05"),
		DiedAt:    todo.DiedAt.Format("2006-01-02 15:04:05"),
	}
}


