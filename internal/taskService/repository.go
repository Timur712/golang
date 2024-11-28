package taskService

import (
	"myproject/internal/database" 
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task Task) (Task, error)
	GetAll() ([]Task, error)
	Update(id uint, task Task) (Task, error)
	Delete(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository() TaskRepository {
	return &taskRepository{db: database.DB}
}

func (r *taskRepository) Create(task Task) (Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *taskRepository) GetAll() ([]Task, error) {
	var tasks []Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) Update(id uint, updatedTask Task) (Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return Task{}, err
	}
	task.Text = updatedTask.Text
	task.IsDone = updatedTask.IsDone
	task.Task = updatedTask.Task
	if err := r.db.Save(&task).Error; err != nil {
		return Task{}, err
	}
	return task, nil
}

func (r *taskRepository) Delete(id uint) error {
	if err := r.db.Delete(&Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
