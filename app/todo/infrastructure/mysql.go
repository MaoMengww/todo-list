package infrastructure

import (
	"time"
	"todo-list/app/todo/domain"

	"gorm.io/gorm"
)

type TodoModel struct {
	TodoId    int64     `gorm:"column:todo_id;primaryKey;autoIncrement"`
	UserId    int64     `gorm:"column:user_id;not null"`
	Title     string    `gorm:"column:title;type:varchar(255);not null"`
	Content   string    `gorm:"column:content;type:text"`
	Completed bool      `gorm:"column:completed;not null;default:false"`
	Priority  int64     `gorm:"column:priority;not null;default:0"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	DiedAt    time.Time `gorm:"column:died_at;autoUpdateTime"`
}

type MysqlTodoRepository struct {
	DB *gorm.DB
}

func NewMysqlTodoRepository(db *gorm.DB) *MysqlTodoRepository {
	return &MysqlTodoRepository{
		DB: db,
	}
}

func (u TodoModel) TableName() string {
	return "todos"
}

func (r *MysqlTodoRepository) Create(todo *domain.Todo) (int64, error) {
	rtodo := &TodoModel{
		UserId:    todo.UserId,
		Title:     todo.Title,
		Content:   todo.Content,
		Completed: todo.Completed,
		Priority:  todo.Priority,
	}
	result := r.DB.Create(rtodo)
	return rtodo.TodoId, result.Error
}

func (r *MysqlTodoRepository) GetById(todoId int64) (*domain.Todo, error) {
	var todoModel TodoModel
	err := r.DB.Where("todo_id = ?", todoId).First(&todoModel).Error
	if err != nil {
		return nil, err
	}
	return &domain.Todo{
		TodoId:    todoModel.TodoId,
		UserId:    todoModel.UserId,
		Title:     todoModel.Title,
		Content:   todoModel.Content,
		Completed: todoModel.Completed,
		Priority:  todoModel.Priority,
		CreatedAt: todoModel.CreatedAt,
		DiedAt:    todoModel.DiedAt,
	}, nil
}

func (r *MysqlTodoRepository) Update(todo *domain.Todo) error {
	rtodo := &TodoModel{
		TodoId:    todo.TodoId,
		UserId:    todo.UserId,
		Title:     todo.Title,
		Content:   todo.Content,
		Completed: todo.Completed,
		DiedAt:    todo.DiedAt,
		Priority:  todo.Priority,
	}
	return r.DB.Model(rtodo).Updates(rtodo).Error
}

func (r *MysqlTodoRepository) Delete(todoId int64) error {
	return r.DB.Delete(&TodoModel{}, todoId).Error
}

func (r *MysqlTodoRepository) ListByUserId(userId int64) ([]*domain.Todo, error) {
	var todos []*TodoModel
	err := r.DB.Where("user_id = ?", userId).Find(&todos).Error
	if err != nil {
		return nil, err
	}
	var result []*domain.Todo
	for _, todoModel := range todos {
		result = append(result, &domain.Todo{
			TodoId:    todoModel.TodoId,
			UserId:    todoModel.UserId,
			Title:     todoModel.Title,
			Content:   todoModel.Content,
			Completed: todoModel.Completed,
			Priority:  todoModel.Priority,
			CreatedAt: todoModel.CreatedAt,
			DiedAt:    todoModel.DiedAt,
		})
	}
	return result, nil
}
