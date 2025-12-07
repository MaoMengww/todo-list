package usecase

import "todo-list/app/todo/domain"

type Usecase struct {
	TodoRepo domain.TodoRepository
}

func NewUsecase(todoRepo domain.TodoRepository) *Usecase { {
	return &Usecase{
		TodoRepo: todoRepo,
	}
}
}

func (uc *Usecase) Create(todo *domain.Todo) (int64, error) {
	return uc.TodoRepo.Create(todo)
}

func (uc *Usecase) GetById(todoId int64) (*domain.Todo, error) {
	return uc.TodoRepo.GetById(todoId)
}

func (uc *Usecase) Update(todo *domain.Todo) error {
	return uc.TodoRepo.Update(todo)
}

func (uc *Usecase) Delete(todoId int64) error {
	return uc.TodoRepo.Delete(todoId)
}

func (uc *Usecase) ListByUserId(userId int64) ([]*domain.Todo, error) {
	return uc.TodoRepo.ListByUserId(userId)
}