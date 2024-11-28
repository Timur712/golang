package taskService


type TaskService interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTask(id uint, task Task) (Task, error)
	DeleteTask(id uint) error
}

type taskService struct {
	repository TaskRepository
}

func NewTaskService(repository TaskRepository) TaskService {
	return &taskService{repository}
}

func (s *taskService) CreateTask(task Task) (Task, error) {
	return s.repository.Create(task)
}

func (s *taskService) GetAllTasks() ([]Task, error) {
	return s.repository.GetAll()
}

func (s *taskService) UpdateTask(id uint, task Task) (Task, error) {
	return s.repository.Update(id, task)
}

func (s *taskService) DeleteTask(id uint) error {
	return s.repository.Delete(id)
}
