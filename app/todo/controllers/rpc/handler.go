package rpc

import (
	"context"
	"time"
	"todo-list/app/todo/controllers/rpc/pack"
	"todo-list/app/todo/domain"
	"todo-list/app/todo/usecase"
	todo "todo-list/kitex_gen/todo"
)

// TodoServiceImpl implements the last service interface defined in the IDL.
type TodoServiceImpl struct{
	usecase *usecase.Usecase
}
func NewTodoServiceImpl(uc *usecase.Usecase) *TodoServiceImpl {
	return &TodoServiceImpl{
		usecase: uc,
	}
}

// AddTodo implements the TodoServiceImpl interface.
func (s *TodoServiceImpl) AddTodo(ctx context.Context, request *todo.AddTodoRequest) (resp *todo.AddTodoResponse, err error) {
	resp = new(todo.AddTodoResponse)
	var diedAt time.Time
	if request.DiedAt != "" {
		diedAt, err = time.Parse("2006-01-02 15:04:05", request.DiedAt)
		if err != nil {
			return
		}
	}
	var todoItem = &domain.Todo{
		Title:     request.Title,
		Content:   request.Content,
		UserId:    request.UserId,
		DiedAt:    diedAt,
		Priority:  request.Priority,
	}
	todoId, err := s.usecase.Create(todoItem)
	if err != nil {
		return &todo.AddTodoResponse{
			Base: pack.NewBadResp(err),
		}, err
	}
	resp = &todo.AddTodoResponse{
		Id: todoId,
		Base: pack.NewGoodResp(),
	}
	return
}

// DeleteTodo implements the TodoServiceImpl interface.
func (s *TodoServiceImpl) DeleteTodo(ctx context.Context, request *todo.DeleteTodoRequest) (resp *todo.DeleteTodoResponse, err error) {
	err = s.usecase.Delete(request.Id)
	if err != nil {
		return
	}
	resp = &todo.DeleteTodoResponse{
		Success: true,
		Base: pack.NewGoodResp(),
	}
	return
}

// UpdateTodo implements the TodoServiceImpl interface.
func (s *TodoServiceImpl) UpdateTodo(ctx context.Context, request *todo.UpdateTodoRequest) (resp *todo.UpdateTodoResponse, err error) {
	resp = new(todo.UpdateTodoResponse)
	var diedAt time.Time
	if *request.DiedAt != "" {
		diedAt, err = time.Parse("2006-01-02 15:04:05", *request.DiedAt)
		if err != nil {
			return &todo.UpdateTodoResponse{
				Base: pack.NewBadResp(err),
			}, err
		}
	}
	var todoItem = &domain.Todo{
		TodoId:    request.Id,
		Title:     *request.Title,
		Content:   *request.Content,
		Completed: *request.Completed,
		DiedAt:    diedAt,
		Priority:  *request.Priority,
	}
	err = s.usecase.Update(todoItem)
	if err != nil {
		return &todo.UpdateTodoResponse{
			Base: pack.NewBadResp(err),
		}, err
	}
	resp = &todo.UpdateTodoResponse{
		Base: pack.NewGoodResp(),
	}
	return
}

// GetTodo implements the TodoServiceImpl interface.
func (s *TodoServiceImpl) GetTodo(ctx context.Context, request *todo.GetTodoRequest) (resp *todo.GetTodoResponse, err error) {
	todoItem, err := s.usecase.GetById(request.Id)
	if err != nil {
		return &todo.GetTodoResponse{
			Base: pack.NewBadResp(err),
		}, err
	}
	resp = &todo.GetTodoResponse{
		Base: pack.NewGoodResp(),
		Todo: &todo.Todo{
			Id:        todoItem.TodoId,
			Title:     todoItem.Title,
			Content:   todoItem.Content,
			UserId:    todoItem.UserId,
			Completed: todoItem.Completed,
			Priority:  todoItem.Priority,
			CreatedAt: todoItem.CreatedAt.Format("2006-01-02 15:04:05"),
			DiedAt:    todoItem.DiedAt.Format("2006-01-02 15:04:05"),
		},
	}
	return
}

// ListTodo implements the TodoServiceImpl interface.
func (s *TodoServiceImpl) ListTodo(ctx context.Context, request *todo.ListTodoRequest) (resp *todo.ListTodoResponse, err error) {
	todoItems, err := s.usecase.ListByUserId(request.UserId)
	if err != nil {
		return &todo.ListTodoResponse{
			Base: pack.NewBadResp(err),
		}, err
	}
	var todoInfos []*todo.Todo
	for _, item := range todoItems {
		todoInfos = append(todoInfos, &todo.Todo{
			Id:        item.TodoId,
			Title:     item.Title,
			Content:   item.Content,
			UserId:    item.UserId,
			Completed: item.Completed,
			Priority:  item.Priority,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
			DiedAt:    item.DiedAt.Format("2006-01-02 15:04:05"),
		})
	}
	resp = &todo.ListTodoResponse{
		Base: pack.NewGoodResp(),
		Todos: todoInfos,
	}
	return
}
