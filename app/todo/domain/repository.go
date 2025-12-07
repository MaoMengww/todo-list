package domain

type TodoRepository interface {
	Create(todo *Todo) (int64, error)
	GetById(todoId int64) (*Todo, error)
	Update(todo *Todo) error
	Delete(todoId int64) error
	ListByUserId(userId int64) ([]*Todo, error)
}