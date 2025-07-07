package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jkeeya/toado/models"
	. "github.com/jkeeya/toado/models"

	sqlite "gorm.io/driver/sqlite"
	gorm "gorm.io/gorm"
)

// Инициализация базы
func InitDB(source string) (DB *gorm.DB) {
	if _, err := os.Stat(source); os.IsNotExist(err) {
		fmt.Println("База не найдена. Создаём новую...")
	}

	var err error
	DB, err = gorm.Open(sqlite.Open(source), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	fmt.Println("Соединение с базой данных установлено!")
	err = DB.AutoMigrate(&(Task{}))

	return
}

type SQLiteTaskRepository struct {
	DB *gorm.DB
}

// Реализация TaskRepository
func (r SQLiteTaskRepository) AddTask(task *Task) error {
	return r.DB.Create(task).Error
}

func (r SQLiteTaskRepository) MarkDone(id uint) error {
	return r.DB.Model(Task{}).Where("id = ?", id).Updates(models.Task{
		Done: true,
	}).Error
}

func (r *SQLiteTaskRepository) DeleteTask(id uint) error {
	return r.DB.Delete(&Task{}, id).Error
}

func (r *SQLiteTaskRepository) GetTasks() []Task {
	var tasks []Task
	err := r.DB.Find(&tasks).Error
	if err != nil {
		return nil
	} else {
		return tasks
	}
}
