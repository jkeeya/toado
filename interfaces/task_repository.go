package interfaces

import (
	. "toado/models"
)

type TaskRepository interface {
	AddTask(task *Task) error
	DeleteTask(uint) error
	GetTasks() ([]Task, error)
}
