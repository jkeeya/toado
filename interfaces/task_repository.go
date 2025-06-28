package interfaces

import (
	. "toado/models"
)

type TaskRepository interface {
	AddTask(task *Task) error
	MarkDone(uint) error
	DeleteTask(uint) error
	GetTasks() ([]Task, error)
}
