package main

import (
	"context"
	todo "todo-list/kitex_gen/todo"
)

// TodoServiceImpl implements the last service interface defined in the IDL.
type TodoServiceImpl struct{}

// AddTodo implements the TodoServiceImpl interface.
func (s *TodoServiceImpl) AddTodo(ctx context.Context, request *todo.AddTodoRequest) (resp *todo.AddTodoResponse, err error) {
	// TODO: Your code here...
	return
}

// DeleteTodo implements the TodoServiceImpl interface.
func (s *TodoServiceImpl) DeleteTodo(ctx context.Context, request *todo.DeleteTodoRequest) (resp *todo.DeleteTodoResponse, err error) {
	// TODO: Your code here...
	return
}

// UpdateTodo implements the TodoServiceImpl interface.
func (s *TodoServiceImpl) UpdateTodo(ctx context.Context, request *todo.UpdateTodoRequest) (resp *todo.UpdateTodoResponse, err error) {
	// TODO: Your code here...
	return
}

// GetTodo implements the TodoServiceImpl interface.
func (s *TodoServiceImpl) GetTodo(ctx context.Context, request *todo.GetTodoRequest) (resp *todo.GetTodoResponse, err error) {
	// TODO: Your code here...
	return
}

// ListTodo implements the TodoServiceImpl interface.
func (s *TodoServiceImpl) ListTodo(ctx context.Context, request *todo.ListTodoRequest) (resp *todo.ListTodoResponse, err error) {
	// TODO: Your code here...
	return
}
